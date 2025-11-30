package main

import "github.com/faizan1191/url-shortner/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080") // Start server on port 8080
}
