package main

import "log"

const URL string = "https://www.boredapi.com/api/activity"

func main() {
	service := NewActivityService(URL)
	service = NewLoggingService(service)
	api := NewAPIService(service)

	log.Fatal(api.Start(":3000"))
}
