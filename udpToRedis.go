package main

import (
	"fmt"
	"gopkg.in/redis.v5"
	"net"
	"os"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		//Network:  "unix",
		//Addr:     "/tmp/redis.sock",
		Addr:     "192.168.2.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":6000")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		client.Publish("hartvigHeading", string(buf[0:n]))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
