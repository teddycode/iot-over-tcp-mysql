package controller

import (
	"DataServer/m/models"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"
)

func HandleConnection(conn net.Conn) {
	device := Device{ // 创建连接对象
		Did:  0,
		Role: 0,
		Con:  conn,
	}
	buffer := make([]byte, 10240) // 建立缓冲区
	fmt.Println("设备已连接：" + device.Con.RemoteAddr().String())
	for {
		num, err := conn.Read(buffer)
		if err != nil { // 断开连接
			break
		}
		fmt.Println("number:", num)
		HandleMessage(buffer[:num], &device)
	}
	HandleDisconnect(&device)
}

// 消息类型
// APP: 认证：{"did":1234,"role":1}
// APP: 命令{ "to_did":0,"cmd":"getNodes"}
func HandleMessage(msg []byte, device *Device) {
	var err error
	var id identity
	var cmd command
	var data itemData
	if device.Role == ROLE_USER {
		fmt.Println("APP消息:", device.Con.RemoteAddr().String(), "> "+string(msg))
	}else if device.Role == ROLE_NODE {
		fmt.Println("节点消息:", device.Con.RemoteAddr().String(), "> "+string(msg))
	}else {
		fmt.Println("未知设备:", device.Con.RemoteAddr().String(), "> "+string(msg))
	}

	switch device.Role {
	case ROLE_USER:
		err = json.Unmarshal(msg, &cmd)
		if err != nil {
			device.Con.Write([]byte("Error:" + err.Error()))
			fmt.Println(err.Error())
		}
		switch cmd.CMD {
		case "getData": // 查询100条记录
			items, err := models.FindItemsByDId(cmd.ToDid)
			if err != nil {
				device.Con.Write([]byte("Error:" + err.Error()))
				fmt.Println(err.Error())
				return
			}
			jsonData, err := json.Marshal(items)
			if err != nil {
				device.Con.Write([]byte("Error:" + err.Error()))
				fmt.Println(err.Error())
				return
			}
			device.Con.Write(jsonData)
		case "getNodes": // 获取节点列表
			var n []int
			for k,_ := range NodeList {
				n = append(n, k)
			}
			json1, err := json.Marshal(n)
			if err != nil {
				device.Con.Write([]byte("Error:" + err.Error()))
				fmt.Println(err.Error())
				return
			}
			device.Con.Write(json1)
		case "subNode": // 订阅节点
			SubsList[cmd.ToDid] = append(SubsList[cmd.ToDid], device.Did)
		case "deSubNode": // 取消订阅节点
			index :=-1
			for i := range SubsList[cmd.ToDid]{
				if SubsList[cmd.ToDid][i] == device.Did{
					index=i
				}
			}
			if index!=-1{
				SubsList[cmd.ToDid] = append(SubsList[cmd.ToDid][:index],SubsList[cmd.ToDid][index+1:]...)
			}

		default:
			str := strings.Split(cmd.CMD,":")
			if str[0] == "send" {
				//下发指令到下位机
				node, ok := NodeList[cmd.ToDid]
				if !ok {
					device.Con.Write([]byte("Node Not Found!"))
					return
				}
				node.Con.Write([]byte(str[1]))
			}
		}
	case ROLE_NODE:
		// 存储至mysql
		err = json.Unmarshal(msg, &data)
		if err != nil {
			//device.Con.Write([]byte("Error:" + err.Error()))
			fmt.Println(err.Error())
		}
		item := models.Item{
			CreatedOn: int(time.Now().Unix()),
			Did:       device.Did,
			Light:     data.Light,
			Mq2:       data.Mq2,
			Mq135:     data.Mq135,
			Temp:      data.Temp,
			Wet:       data.Wet,
		}
		_, err = models.NewItem(&item)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			//device.Con.Write([]byte("Error:" + err.Error()))
			return
		}
		// 发送至已订阅节点的APP
		for _,app := range SubsList[device.Did]{
			if AppList[app].Con != nil{
				AppList[app].Con.Write(msg)
			}
		}
	default: // 获取身份
		err = json.Unmarshal(msg, &id)
		if err != nil {
			device.Con.Write([]byte("Error:" + err.Error()))
			fmt.Println(err.Error())
			return
		}
		device.Did = id.Did
		device.Role = id.Role
		if device.Role == ROLE_USER {
			AppList[device.Did] = *device //添加至App在线列表
			fmt.Println("当前设备身份为：APP")
		} else if device.Role == ROLE_NODE {
			NodeList[device.Did] = *device //添加至Node在线列表
			fmt.Println("当前设备身份为：节点")
		}
	}
}

func HandleDisconnect(device *Device) {
	key := -1
	for k, v := range AppList {
		if v.Con == device.Con {
			key = k
			fmt.Println("APP_", k, " at ", device.Con.RemoteAddr().String()+" 连接已断开.")
			break
		}
	}
	if key != -1 {
		delete(AppList, key)
		return
	}
	key = -1
	for k, v := range NodeList {
		if v.Con == device.Con {
			key = k
			fmt.Println("Node_", k, " at ", device.Con.RemoteAddr().String()+" 连接已断开.")
			break
		}
	}
	if key != -1 {
		delete(NodeList, key)
		return
	}
	fmt.Println("未知设备", " at ", device.Con.RemoteAddr().String()+" 连接已断开.")
	return
}
