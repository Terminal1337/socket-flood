package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/ogier/pflag"
)

var (
	host    string
	port    string
	threads string
	length  string
)

func openSocket() {
	for {
	restart:
		con, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		if err != nil {
			goto restart
		}

		msg := ""

		_, err = con.Write([]byte(msg))
		if err != nil {
			goto restart
		}

		reply := make([]byte, 1024)
		_, err = con.Read(reply)
		if err != nil {
			goto restart
		}

	}
}

func main() {
	var err error
	pflag.StringVar(&host, "host", "localhost", "Hostname or IP address")
	pflag.StringVar(&port, "port", "80", "Port number")
	pflag.StringVar(&threads, "threads", "100", "Number of threads")
	pflag.StringVar(&length, "length", "60", "Length of the client execution in seconds")
	pflag.Parse()

	if pflag.NFlag() == 0 {
		pflag.Usage()
		os.Exit(1)
	}
	threads_int, err := strconv.Atoi(threads)
	if err != nil {
		fmt.Println("Threads must be an integer")
		os.Exit(1)
	}
	length_int, err := strconv.Atoi(length)
	if err != nil {
		fmt.Println("Length must be an integer")
		os.Exit(1)
	}
	fmt.Printf("HOST: %s\n", host)
	fmt.Printf("PORT: %s\n", port)
	fmt.Printf("THREADS: %s\n", threads)
	for i := 0; i < threads_int; i++ {
		fmt.Printf("\rINFO: Starting thread [%d]", i)
		go openSocket()
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	time.Sleep(time.Duration(length_int) * time.Second)
}
