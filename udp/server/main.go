package main

import (
	"fmt"
	"net"
)

// UDP 协议 是无连接的传输层协议，不需要建立连接就可以进行数据发送和接收，属于不可靠，没有时序的通信，但是 实时性比较好，

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}

	defer conn.Close()

	for {
		var data [1024]byte
		n, addr, err := conn.ReadFromUDP(data[:]) //接收数据
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			continue
		}

		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = conn.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed,err:", err)
			continue
		}
	}
}
