package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ping(s string, c chan string, done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(time.Second)
			c <- s
		}
	}

}

func pong(s string, c chan string, done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(time.Second)
			c <- s
		}
	}

}

func main() {
	c := make(chan string)
	done := make(chan bool)

	go ping("ping", c, done)
	go pong("pong", c, done)

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		<-sig

		done <- true
		done <- true
	}()

	go func() {
		fmt.Println("Pressione qualquer tecla para sair...")
		var input string
		fmt.Scanln(&input)

		done <- true
		done <- true
	}()

	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-done:
			fmt.Println("Saindo...")
			return
		}
	}
}
