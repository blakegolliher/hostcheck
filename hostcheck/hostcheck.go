package main

import (
	"os"
	"fmt"
	"log"
	"net"
	"flag"
	"time"
	"github.com/bogdanovich/dns_resolver"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: tcpcheck.go [hostname]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func dnscheck(host string) {
	resolver := dns_resolver.New([]string{"8.8.8.8", "8.8.4.4"})
	resolver.RetryTimes = 2
	//ip, err := resolver.LookupHost(host)
	_, err := resolver.LookupHost(host)
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Println(ip)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("hostname input missing.")
		os.Exit(1)
	}
	var status string
	host := os.Args[1]
	dnscheck(host)
	conn, err := net.DialTimeout("tcp", host+":22", 1*time.Second)
	if err != nil {
		log.Println("Connection error:", err)
		status = "Not Reachable on port 22"
	} else {
		status = "Reachable on port 22"
		defer conn.Close()
	}
	log.Println(host, status)
}
