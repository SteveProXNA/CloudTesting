package main

import (
	"app/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	registry.SetupRegistryService()
	http.Handle("/services", &registry.Service{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = ":3000"

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			panic(err)
		}
		err = srv.Shutdown(ctx)
		if err != nil {
			panic(err)
		}
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
