# eeureka
Simplistic Eureka client for Go Microservices

## Usage

Import

    import "github.com/eriklupander/eeureka"
    
In your code, call the Register method:

    eeureka.Register("myMicroservice", "8080", "8443")
    
    or
    
    eeureka.RegisterAt("http://192.168.123.123:8761","myMicroservice", "8080", "8443")
    
    
    
The register method will try to contact the Eureka server indefinitely. When registration succeeds (HTTP 204), heartbeats (PUTs) will be issued every 30 seconds. When the microservice exits by a Sigterm or OS interrupt signal, the microservice will deregister itself with Eureka before shutting down.

### Configuration

By default, the Eureka server is assumed to be at http://192.168.99.100:8761

TODO fix configuration...

### Spring Cloud vs Netflix Eureka

This lib has only been tested with the Spring Cloud flavour of Eureka. Internally, this lib uses the REST endpoints of Eureka which exists in two versions. With "/v2" or without. This lib is only tested with the non-v2 version.

