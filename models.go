/**
The MIT License (MIT)

Copyright (c) 2016 ErikL

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package eeureka

/**
  Defines a graph of structs that conforms to a part of the return type of the Eureka "get instances for appId", e.g:

  GET /eureka/v2/apps/appID

  The root is the EurekaServiceResponse which contains a single EurekaApplication, which in its turn contains an array
  of EurekaInstance instances.
*/

// Response for /eureka/apps/{appName}
type EurekaServiceResponse struct {
	Application EurekaApplication `json:"application"`
}

// Response for /eureka/apps
type EurekaApplicationsRootResponse struct {
	Resp EurekaApplicationsResponse `json:"applications"`
}

type EurekaApplicationsResponse struct {
	Version      string              `json:"versions__delta"`
	AppsHashcode string              `json:"versions__delta"`
	Applications []EurekaApplication `json:"application"`
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
