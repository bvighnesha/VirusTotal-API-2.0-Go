package main

import "me/vighnesh/api/virustotal"
import "fmt"

func main() {
	virustotalapi, _ := virustotal.Configure("APIKEY")
	response, e := virustotalapi.FileReport("Resource")
	if e != nil {
		fmt.Println(e)
	}

	for engine, report := range response.Scans {
		fmt.Println("Scan Engine", engine)
		if report.Detected {
			fmt.Println("Version", report.Version)
			fmt.Println("Updated", report.Update)
			fmt.Println("Malware", report.Malware)
		} else {
			fmt.Println("No Malware Detected")
		}

	}
}
