# eeureka
Simplistic Eureka client for Go Microservices

## What is is?
I use this in my personal Microservice projects where I use Go-based microservices in a Spring Cloud context deployed on docker containers.

Internally, it's basically wraps a little HTTP client for the following operations of the [Netflix Eureka REST API](https://github.com/Netflix/eureka/wiki/Eureka-REST-operations) :

- register appId
- heartbeat appId
- deregister appId
- instances of appId

The top three is enough to handle the lifecycle of a Microservice while the "instances of appId" can be used to get hostname/port for all running instances of a particular appId. Useful for client-side load-balancing.

## Usage

Import

    import "github.com/eriklupander/eeureka"
    
In your code, call the Register method:

    eeureka.Register("myMicroservice", "8080", "8443")
    
or
    
    eeureka.RegisterAt("http://192.168.123.123:8761","myMicroservice", "8080", "8443")
        
The register method will try to contact the Eureka server indefinitely. When registration succeeds (HTTP 204), heartbeats (PUTs) will be issued every 30 seconds. When the microservice exits by a Sigterm or OS interrupt signal, the microservice will deregister itself with Eureka before shutting down.

### Public functions
- RegisterAt - registers your application at a specific Eureka service URL.
- Register - registers your application at the default Eureka service URL http://192.168.99.100:8761 (e.g. typical local Docker installation)
- GetServiceInstances - Returns all running instances of a given appName

The register methods automatically handles retries, heartbeats and deregistration.

### Configuration

By default, the Eureka server is assumed to be at http://192.168.99.100:8761

Otherwise, use the _RegisterAt_ function.

### Spring Cloud vs Netflix Eureka

This lib has only been tested with the Spring Cloud flavour of Eureka. Internally, this lib uses the REST endpoints of Eureka which exists in two versions. With "/v2" or without. This lib is only tested with the non-v2 version.

