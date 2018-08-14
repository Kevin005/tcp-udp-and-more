package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 创建监听
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8888,
	})
	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	fmt.Println("监听成功")
	defer socket.Close()
	for {
		// 读取数据
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)

		fmt.Println("1----> ", time.Now())
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}
		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n\n", data)
		// 发送数据
		senddata := []byte("hello client!")
		_, err = socket.WriteToUDP(senddata, remoteAddr)

		fmt.Println("2----> ",time.Now())
		if err != nil {
			return
			fmt.Println("发送数据失败!", err)
		}
	}
}
