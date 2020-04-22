package controller

import "net"

const ROLE_OTHER = 0
const ROLE_USER = 1
const ROLE_NODE = 2

type Device struct {
	Did  int      `json:"did"`  // 节点ID
	Role int      `json:"role"` // 角色：APP/节点
	Con  net.Conn `json:"-"`    // 连接对象
}

type identity struct {
	Did  int `json:"did"`
	Role int `json:"role"`
}

type command struct {
	ToDid int    `json:"to_did"`
	CMD   string `json:"cmd"`
}

type itemData struct {
	Light float32 `json:"light"` // 光强
	Mq2   float32 `json:"mq2"`   // Mq2传感器值
	Mq135 float32 `json:"mq135"` // Mq135 传感器
	Temp  float32 `json:"temp"`  //温度
	Wet   float32 `json:"wet"`   //湿度
}

var NodeList = make(map[int]Device, 20) //在线节点列表
var AppList = make(map[int]Device, 10)  //在线App列表
