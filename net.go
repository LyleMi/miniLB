package main

import (
	"io"
	"log"
	"net"
)

func forwardData(src, dst net.Conn) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Printf("Failed to forward data: %v", err)
	}
}

func handleConnection(entry *Entry, conn net.Conn) {
	defer conn.Close()

	remote, err := net.Dial("tcp", entry.TargetService)
	if err != nil {
		log.Printf("Failed to connect to target service: %v", err)
		return
	}
	defer remote.Close()

	go forwardData(conn, remote)
	forwardData(remote, conn)
}

func HandleEntry(entry *Entry) {
	listener, err := net.Listen("tcp", entry.ListenAddress)
	log.Printf("forward %s to %s", entry.ListenAddress, entry.TargetService)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", entry.ListenAddress, err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		log.Printf("Accept connection: %s -> %s", conn.RemoteAddr().String(), entry.TargetService)
		go handleConnection(entry, conn)
	}
}
