package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csv := flag.String("file", "", "csv file to read")
	flag.Parse()
	if *csv == "" {
		fmt.Println("csv file is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !strings.Contains(*csv, ".csv") {
		fmt.Println("csv file is required")
		os.Exit(1)
	}
	ips, err := filter(*csv)
	if err != nil {
		fmt.Printf("Error filtering %s: %v", *csv, err)
		os.Exit(1)
	}
	for _, ip := range ips {
		ping(ip, *csv)
	}
}
