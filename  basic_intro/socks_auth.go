package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed: %w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver: %v", ver)
	}
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed: %w", err)
	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed: %w", err)
	}
	log.Println("ver", ver, "method", method)
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("writed failed: %w", err)
	}
	return nil
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed: %v", conn.RemoteAddr(), err)
		return
	}
	log.Println("author acess")
}

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}
