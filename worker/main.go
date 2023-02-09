package main

import (
	"fmt"
	"sync"

	"github.com/MantisSTS/BountyProcess/server/helpers"
)

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 100)
	// Fetch the RabbitMQ messages
	rmq := helpers.RabbitMQHelper{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				rmq.Fetch("chan.recon.subdomains.amass", "domain", results)
			}
		}()
	}

	// Process the results
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case res := <-results:
					fmt.Println(res)
				}
			}
		}()
	}

	wg.Wait()
}
