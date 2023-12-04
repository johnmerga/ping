package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func ping(name string) {
	pinger, err := probing.NewPinger(name)
	pinger.Timeout = time.Second * 1
	if err != nil {
		return
	}

	// Listen for Ctrl-C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}

	pinger.OnFinish = func(stats *probing.Statistics) {
		// if it is reachable save the ip address to a file
		if stats.PacketsRecv > 0 {
			pingToFile(pinger.IPAddr().String())
		}
		// fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		// fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
		// 	stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		// fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
		// 	stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
}

// this save the ip address to a file if the file exit append to the file in new line
func pingToFile(ip string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile("ip.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(ip + "\n")
	if err != nil {
		return
	}

	// save changes
	err = file.Sync()
	if err != nil {
		return
	}

	//fmt.Println("==> done writing to file")
}
