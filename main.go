package main

import (
	"DataServer/m/config"
	c "DataServer/m/controller"
	"fmt"
	"net"
	"strconv"
)

func main() {
	fmt.Println("正在开启TCP服务...")
	src := "192.168.43.243:" + strconv.Itoa(config.HTTPPort)
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("TCP开启成功： %s.\n", src)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("连接失败： %s\n", err)
		}
		go c.HandleConnection(conn)
	}
}
