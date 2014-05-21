package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const KB = 1024

func main() {
	nbCpus := flag.Uint("nb-cpus", 1, "Number of CPUs to burn")
	flag.Parse()

	// Setup the number of OS threads used by Go runtime
	runtime.GOMAXPROCS(int(*nbCpus))

	beginningTime := time.Now()

	var i uint
	for i = 0; i < *nbCpus; i++ {
		go cpuBurn(i)
	}

	// Trap SIGINT SIGTERM and SIGQUIT
	signalNotification := make(chan os.Signal, 1)
	signal.Notify(signalNotification, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGSTOP)

	<-signalNotification

	endTime := time.Now()
	fmt.Printf("Elapsed time: %v\n", endTime.Sub(beginningTime))
	os.Exit(0)
}

func cpuBurn(index uint) {
	fmt.Println("Starting burning routing nb", index)

	hashBuffer := make([]byte, 1*KB)
	for {
		hash := sha512.New()
		hash.Write(hashBuffer)
		hash.Sum(nil)
	}
}
