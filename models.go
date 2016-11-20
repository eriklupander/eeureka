package eeureka

type EurekaServiceResponse struct {
	Application EurekaApplication `json:"application"`
}

type EurekaApplication struct {
	Name     string           `json:"name"`
	Instance []EurekaInstance `json:"instance"`
}

type EurekaInstance struct {
	HostName string     `json:"hostName"`
	Port     EurekaPort `json:"port"`
}

type EurekaPort struct {
	Port int `json:"$"`
}
