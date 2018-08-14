package main

import (
	"net"
	"fmt"
)

func main() {
	// tcp 监听并接受端口
	l, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	//最后关闭
	defer l.Close()
	fmt.Println("tcp服务端开始监听65535端口...")
	// 使用循环一直接受连接
	for {
		fmt.Println("loop test")
		//Listener.Accept() 接受连接
		//conn 是双方的。和长度为1的channel有些类似。
		c, err := l.Accept()
		if err != nil {
			return
		}
		//处理tcp请求
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	//一些代码逻辑...
	fmt.Println("tcp服务端开始处理请求...")
	//读取
	buffer := make([]byte, 1024)
	//如果客户端无数据则会阻塞，服务端阻塞，直到等待客户端传递数据。
	c.Read(buffer)

	//服务端成功从阻塞状态走出，读取客户端的数据，并根据自身的接口输出buffer
	c.Write(buffer)
	fmt.Println("tcp服务端开始处理请求完毕...")
}
