package service

import (
	"app/registry"
	"context"
	"fmt"
	stlog "log"
	"net/http"
)

func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)

	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		stlog.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			stlog.Println(err)
		}
		err = registry.ShutdownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {
			stlog.Println(err)
		}
		err = srv.Shutdown(ctx)
		if err != nil {
			stlog.Println(err)
		}
		cancel()
	}()

	return ctx
}
