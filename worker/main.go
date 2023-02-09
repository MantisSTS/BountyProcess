package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	helpers "github.com/MantisSTS/BountyProcess/server"
)

var (
	// Create a RabbitMQHelper
	rmq helpers.RabbitMQHelper
)

func runSubdomainRecon(domain string, wg *sync.WaitGroup, results chan string) {
	log.Println("Running recon on:", domain)

	outDir := "./output/recon/subdomains/" + domain + "/"

	// Execute amass
	go func() {

		// create directory
		err := os.MkdirAll(outDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("amass", "enum", "-d", domain, "-json", outDir+"amass.json")
		cmd.Run()

		// Read file using bufio.NewScanner
		file, err := os.Open(outDir + "amass.json")
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)
		if err != nil {
			log.Fatal(err)
		}

		for scanner.Scan() {
			rmq.Publish("chan.recon.subdomains", "subdomain")
			results <- strings.TrimSpace(scanner.Text())
		}
	}()

	// Execute subfinde

	// Execute assetfinder
	// Execute findomain

}

func main() {

	var wg sync.WaitGroup

	domainResults := make(chan string, 100)
	subdomainReconResults := make(chan string, 100)
	// Fetch the RabbitMQ messages

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				rmq.Fetch("chan.recon.domains", "domain", domainResults)
			}
		}()
	}

	// Process the domainResults
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case res := <-domainResults:
					runSubdomainRecon(res, &wg, subdomainReconResults)
				}
			}
		}()
	}

	wg.Wait()
}
