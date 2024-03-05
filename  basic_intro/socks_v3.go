package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05
const cmdBind = 0x01
const atypeIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func process_v3(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	err := auth_v3(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed %v", conn.RemoteAddr(), err)
		return
	}
	err = connect(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
	}
}
func auth_v3(reader *bufio.Reader, conn net.Conn) (err error) {
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%w", ver)
	}
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}

	return nil
}

func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("reader failed %w", err)
	}
	ver, cmd, atype := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("reader header failed: %w", err)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd: %w", err)
	}
	addr := ""
	switch atype {
	case atypeIPV4:
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read artpe failed: %w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case atypeHOST:
		hostsize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostsize failed:%w", err)
		}
		host := make([]byte, hostsize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed: %w", err)
		}
		addr = string(host)
	case atypeIPV6:
		return errors.New("Not supported yet")
	default:
		return errors.New("invalid atyp")
	}

	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed: %w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])
	log.Println("dial", addr, port)
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	return nil
}

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process_v3(client)
	}
}
