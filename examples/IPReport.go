package main

import (
	"fmt"
	"vighnesh.org/virustotal"
)

func main() {
	virustotalapi, _ := virustotal.Configure("APIKEY")
	response, e := virustotalapi.IPReport("90.156.201.27")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Resolutions", response.Resolutions)
	fmt.Println("Detected", response.DetectedUrls)
}
