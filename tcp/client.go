package main

import (
	"net"
	"fmt"
)

func main() {
	//net.dial 拨号 获取tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("获取127.0.0.1：65535的tcp连接成功...")

	//客户端这里不用使用协程。使用协程的话main函数退出，所有go 协程全部死掉。
	defer conn.Close()

	//服务端会返回这段数据
	conn.Write([]byte("echo data to server ,then to client!!!"))
	fmt.Println("test server")
	//读取到buffer
	buffer := make([]byte, 1024)
	//如果服务端没有把数据传递过来，那么客户端阻塞，直到服务端向其中写入了数据。
	conn.Read(buffer)
	fmt.Println(string(buffer))
}
