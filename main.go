package main

import (
	"fmt"
	"transfer"
)

func main() {
	/*
	   	v := `{
	   	"computer_id": "BFEBFBFF000906E90000000000000000221B2ED0",
	   	"ip_list": "169.254.45.61,192.168.2.1,192.168.230.1,192.168.0.203",
	   	"computer_type": 2,
	   	"computer_system": "Windows 10 Enterprise()",
	   	"compputer_system_bit": 64,
	   	"screen_pix": "1920×1080",
	   	"memory_info": {
	   		"computer_total": "15.87G",
	   		"computer_available": "10.65G",
	   		"Collection_process_memmory_status": ""
	   	},
	   	"cpu_info": {
	   		"Collection_process_cpu_status": "",
	   		"cpu_model_name": "Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz"
	   	},
	   	"version": "1.0.2",
	   	"mdc_process_status": 0,
	   	"firewall_status": "Firewall Domain is enabled\nFirewall Private is enabled\nFirewall Public is enabled\n",
	   	"mac_address": "以太网:10:7B:44:80:ED:7C,VMware Network Adapter VMnet1:00:50:56:C0:00:01,VMware Network Adapter VMnet8:00:50:56:C0:00:08,Npcap Loopback Adapter:02:00:4C:4F:4F:50",
	   	"remote_status": "Remote port:3389\r\nRemote Connect: disabled",
	   	"tv_sunflower_status": "Temviewers: 开启,SunloginClient: 开启",
	   	"robot_version": "",
	   	"sendTime": 1564042362
	   }`
	   	var t interface{}
	   	err := json.Unmarshal([]byte(v), &t)
	   	if err != nil {
	   		fmt.Println(err)
	   		return
	   	}
	   	data := t.(map[string]interface{})
	   	value, ok := data["computer_type"].(float64)
	   	if ok == true {
	   		fmt.Println(value)
	   	} else {
	   		fmt.Println(value)
	   	}
	*/
	transfer.Start()
	fmt.Println("----------")
	/*
		data := xmlparse.Reseult{Coll_proce: "none", Version: "none"}
		ok := xmlparse.GetConfig(&data)
		if ok == false {
			fmt.Println("parse filad")
			return
		}
		fmt.Println(data.Type)
		fmt.Println(data.While_time)
		fmt.Println(data.Coll_proce)
		fmt.Println(data.Remote_monitor)
		for key, value := range data.Remote_monitor {
			fmt.Println(key, "------", value.Ischeck, value.Timeout, value.Remote_item)
		}
		fmt.Println(data.Request_address)
		fmt.Println(data.Delegation_serivce)
		for key, value := range data.Delegation_serivce {
			fmt.Println(key, "------", value)
		}
		fmt.Println(data.Version)
		fmt.Println(data.Client_timeout)
		fmt.Println(data.Client_num)
	*/

}
