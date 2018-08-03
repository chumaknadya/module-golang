package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net"
	"time"
)

var tcpIp = "127.0.0.1:8888"

func fibonacci(n *big.Int) *big.Int {
	f2 := big.NewInt(0)
	f1 := big.NewInt(1)
	if n.Cmp(big.NewInt(1)) == 0 {
		return f2
	}
	if n.Cmp(big.NewInt(2)) == 0 {
		return f1
	}
	for i := 2; n.Cmp(big.NewInt(int64(i))) >= 0; i++ {
		next := big.NewInt(0)
		next.Add(f2, f1)
		f2 = f1
		f1 = next
	}
	return f1
}

func Produce(conn net.Conn) {
	cash := make(map[int64]*big.Int)
	for {
		var msg int64
		var ans string
		d := json.NewDecoder(conn)
		err := d.Decode(&msg)
		if err == io.EOF {
			conn.Close()
			return
		}
		fmt.Println("Msg Recieved:", msg)
		if val, ok := cash[msg]; !ok {
			start := time.Now()
			answer := fibonacci(big.NewInt(msg))
			cash[msg] = answer
			time := time.Since(start)
			fmt.Println(time)
			ans = time.String() + " " + answer.String()
		} else {
			ans = time.Duration(0).String() + " " + val.String()
		}
		Send(conn, ans)
	}
}

func Send(conn net.Conn, answer string) {
	encoder := json.NewEncoder(conn)
	fmt.Println("Send number int the Fibonacci sequence to client:", answer)
	err := encoder.Encode(answer)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", tcpIp)
	if err != nil {
		fmt.Println("Can't launch server")
		fmt.Println("Reason: ", err)
	}
	fmt.Println("Launch server...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Can't recieve data")
			fmt.Println("Reason: ", err)
		}
		go Produce(conn)
	}
}
