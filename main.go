// Copyright © 2019 Sonoko Mizuki, Ltd. All rights reserved.

package main

import (
	"github.com/MizukiSonoko/goapp-template/controller"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf("start go app template")

	errC := make(chan error)
	endpoint := ":8080"
	go func() {
		if err := controller.Run(endpoint); err != nil {
			errC <- err
		}
	}()

	quitC := make(chan os.Signal, 1)
	signal.Notify(quitC, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errC:
		panic(err)
	case <-quitC:
		log.Printf("System shutdown without problem")
	}
}

