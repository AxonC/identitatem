package main

import (
	"os"

	"github.com/AxonC/identitatem/pkg/authentication"
)

func main() {
	port := os.Getenv("LISTEN_PORT")
	router := authentication.NewIdentityProvider()
	router.Run(port)
}
