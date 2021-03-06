// Overly complicated tcp hello world
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatalf("Could not dial tcp: %v", err)
	}

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, os.Kill, syscall.SIGQUIT)
	defer close(sigquit)

	conchan := make(chan net.Conn, 10)
	defer close(conchan)

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %v\n", err)
				continue
			}
			conchan <- c
		}
	}()

loop:
	for {
		select {
		case c := <-conchan:
			fmt.Fprint(c, "Hello World!")
			c.Close()
		case <-sigquit:
			break loop
		}
	}

	log.Println("Exited gracefully")
}
