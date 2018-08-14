package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 创建连接
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8888,
	})
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}
	defer socket.Close()
	// 发送数据
	senddata := []byte("hello server!")
	_, err = socket.Write(senddata)

	fmt.Println("3----> ", time.Now())
	if err != nil {
		fmt.Println("发送数据失败!", err)
		return
	}
	// 接收数据
	data := make([]byte, 4096)
	read, remoteAddr, err := socket.ReadFromUDP(data)

	fmt.Println("4----> ", time.Now())
	if err != nil {
		fmt.Println("读取数据失败!", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data)
}
