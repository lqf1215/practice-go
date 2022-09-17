package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("连接服务器失败， err:", err)
		return
	}

	defer socket.Close()
	sendData := []byte("hello sereve")
	_, err = socket.Write(sendData) //发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, addr, err := socket.ReadFromUDP(data) // 接数数据
	if err != nil {
		fmt.Println("接收数据失败， err:", err)
		return
	}

	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), addr, n)

}
