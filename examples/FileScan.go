package main

import (
	"fmt"
	"vighnesh.org/virustotal"
)

func main() {
	virustotalapi, _ := virustotal.Configure("APIKEY")

	response, e := virustotalapi.ScanFile("FILE TO SCAN")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Verbose Message", response.Message)
	fmt.Println("SHA-1", response.Sha1)
	fmt.Println("SHA-256", response.Sha256)
	fmt.Println("MD5", response.Md5)
	fmt.Println("Permalink", response.Permalink)
	fmt.Println("Resource", response.Resource)
	fmt.Println("ScanId", response.ScanId)
	fmt.Println("Response Code", response.ResponseCode)
}
