package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	// 微信SDK包
)

var port string

func main() {
	// in other programming languages, this might look like:
	//    s = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP)
	//    s.bind("fly-global-services", port)
	port = os.Getenv("PORT")

	udp, err := net.ListenPacket("udp", fmt.Sprintf("fly-global-services:%s", port))
	if err != nil {
		log.Fatalf("can't listen on %s/udp: %s", port, err)
	}

	// in other programming languages, this might look like:
	//    s = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP)
	//    s.bind("0.0.0.0", port)
	//    s.listen()

	tcp, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("can't listen on %s/tcp: %s", port, err)
	}

	go handleTCP(tcp)

	handleUDP(udp)
}

func handleTCP(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}

			log.Printf("error accepting on %s/tcp: %s", port, err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	lines := bufio.NewReader(c)

	for {
		line, err := lines.ReadString('\n')
		if err != nil {
			return
		}

		c.Write([]byte(line))
	}
}

func handleUDP(c net.PacketConn) {
	packet := make([]byte, 2000)

	for {
		n, addr, err := c.ReadFrom(packet)
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}

			log.Printf("error reading on %s/udp: %s", port, err)
			continue
		}

		c.WriteTo(packet[:n], addr)
	}
}
