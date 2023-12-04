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
		panic("csv file is required")
	}
	ips, err := filter(*csv)
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		ping(ip)
	}
}
