package main

import (
	"github.com/AxonC/identitatem/pkg/authentication"
)

func main() {
	router := authentication.NewIdentifyProvider()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
