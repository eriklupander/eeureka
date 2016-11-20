package eeureka

import (
	"encoding/json"
	"fmt"
	"github.com/twinj/uuid"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var instanceId string
var discoveryServerUrl = "http://192.168.99.100:8761"

var regTpl = `{
  "instance": {
    "hostName":"${ipAddress}",
    "app":"${appName}",
    "ipAddr":"${ipAddress}",
    "vipAddress":"${appName}",
    "status":"UP",
    "port":"${port}",
    "securePort" : "${securePort}",
    "homePageUrl" : "http://${ipAddress}:${port}/",
    "statusPageUrl": "http://${ipAddress}:${port}/info",
    "healthCheckUrl": "http://${ipAddress}:${port}/health",
    "dataCenterInfo" : {
      "name": "MyOwn"
    },
    "metadata": {
      "instanceId" : "${appName}:${instanceId}"
    }
  }
}`

func RegisterAt(eurekaUrl string, appName string, port string, securePort string) {
	discoveryServerUrl = eurekaUrl
	Register(appName, port, securePort)
}

func Register(appName string, port string, securePort string) {
	instanceId = getUUID()

	tpl := string(regTpl)
	tpl = strings.Replace(tpl, "${ipAddress}", getLocalIP(), -1)
	tpl = strings.Replace(tpl, "${port}", port, -1)
	tpl = strings.Replace(tpl, "${securePort}", securePort, -1)
	tpl = strings.Replace(tpl, "${instanceId}", instanceId, -1)
	tpl = strings.Replace(tpl, "${appName}", appName, -1)

	// Register.
	registerAction := HttpAction{
		Url:         discoveryServerUrl + "/eureka/apps/" + appName,
		Method:      "POST",
		ContentType: "application/json;charset=UTF-8",
		Body:        tpl,
	}
	var result bool
	for {
		result = doHttpRequest(registerAction)
		if result {
			fmt.Println("Registration OK")
			handleSigterm(appName)
			go startHeartbeat(appName)
			break
		} else {
			fmt.Println("Registration attempt of " + appName + " failed...")
			time.Sleep(time.Second * 5)
		}
	}

}

func startHeartbeat(appName string) {
	for {
		time.Sleep(time.Second * 30)
		heartbeat(appName)
	}
}

func GetServiceInstances(appName string) (EurekaApplication, error) {
	var m EurekaServiceResponse

	queryAction := HttpAction{
		Url:    discoveryServerUrl + "/eureka/apps/" + appName,
		Method: "GET",
		Accept: "application/json;charset=UTF-8",
	}
	log.Println("Doing queryAction using URL: " + queryAction.Url)
	bytes, err := executeQuery(queryAction)
	if err != nil {
		return EurekaApplication{}, err
	} else {
		err := json.Unmarshal(bytes, &m)
		if err != nil {
			fmt.Println("Problem parsing JSON response from Eureka: " + err.Error())
		}
		return m.Application, nil
	}
}

func heartbeat(appName string) {
	heartbeatAction := HttpAction{
		Url:    discoveryServerUrl + "/eureka/apps/" + appName + "/" + getLocalIP(),
		Method: "PUT",
	}
	doHttpRequest(heartbeatAction)
}

func Deregister(appName string) {
	fmt.Println("Trying to deregister application...")
	// Deregister
	deregisterAction := HttpAction{
		Url:    discoveryServerUrl + "/eureka/apps/" + appName + "/" + getLocalIP(),
		Method: "DELETE",
	}
	doHttpRequest(deregisterAction)
	fmt.Println("Deregistered application, exiting. Check Eureka...")
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getUUID() string {
	return uuid.NewV4().String()
}

func handleSigterm(appName string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		Deregister(appName)
		os.Exit(1)
	}()
}
