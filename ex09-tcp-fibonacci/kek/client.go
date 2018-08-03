package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

var tcpIp = "127.0.0.1:8888"

func main() {
	conn, err := net.Dial("tcp", tcpIp)
	if err != nil {
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		var answer string
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		Send(conn, input)
		d := json.NewDecoder(conn)
		err = d.Decode(&answer)
		if err != nil {
			return
		}
		fmt.Println(answer)
	}
}

func Send(conn net.Conn, msg string) {
	encoder := json.NewEncoder(conn)
	tosend, _ := strconv.ParseInt(msg[:len(msg)-1], 10, 32)
	encoder.Encode(tosend)
}
