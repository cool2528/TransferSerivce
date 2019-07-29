package transfer

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"xmlparse"
)

var (
	computer_num   map[int]string
	mDwClientExist int
	mRequest_url   string
)

/*
@ 接收客户端任务并且转发给服务端
*/

//http 请求
func httpCallBack() {
	http.HandleFunc("/sendData", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		buffer := make([]byte, 1024)
		var result string
		for {
			n, err := r.Body.Read(buffer)
			result += string(buffer[:n])
			if n == 0 && err == io.EOF {
				break
			}
		}
		fmt.Println(result)
		var t interface{}
		err := json.Unmarshal([]byte(result)[:len(result)], &t)
		if err != nil {
			w.Write([]byte("json error"))
			fmt.Println(err)
			return
		}
		data := t.(map[string]interface{})
		if value, ok := data["computer_type"].(float64); ok == true {
			computer_num[int(value)] = result
		}
		w.Write([]byte("ok!"))
	})
	http.ListenAndServe(":6778", nil)
}

/*
启动程序 这个函数做两件时间
@ 1 启动一个协程 处理来自http 的请求
@ 2 启动一个timer 根据间隔时间来发送来自与客户机的数据转处理数据来判断客户机是否存活来判断 所有机器是否存活
*/
func Start() bool {
	computer_num = make(map[int]string)
	data := xmlparse.Reseult{Coll_proce: "none", Version: "none"}
	ok := xmlparse.GetConfig(&data)
	if ok == false {
		fmt.Println("parse filad")
		return false
	}
	time_out := data.Client_timeout
	time_out *= 60
	mDwClientExist = data.Client_num
	mRequest_url = data.Request_address
	go httpCallBack()
	t := time.NewTicker(time.Duration(time_out) * time.Second)
	go func() {
		for {
			//定时器操作的
			<-t.C
			currentTime := time.Now().Unix()
			var dwInterval uint
			fmt.Println("timer 超时执行了")
			var dwExist int
			var sendData string
			dwExist = 0
			for _, value := range computer_num {
				var t interface{}
				err := json.Unmarshal([]byte(value), &t)
				if err != nil {
					continue
				}
				data := t.(map[string]interface{})
				if v, ok := data["sendTime"].(float64); ok == true {
					dwInterval = uint((currentTime - int64(v)) / 60)
					if v, ok := data["mdc_process_status"].(float64); ok == true {
						if dwInterval <= 10 && uint(v) == 1 {
							dwExist++
						}
					}
				}
			}
			if mDwClientExist > 0 {
				if dwExist >= mDwClientExist {
					sendData = "status=1"
				} else {
					sendData = "status=0"
				}
			} else {
				sendData = "status=0"
			}
			fmt.Println(sendData)
			if mRequest_url != "" {
				url := mRequest_url + "collectInfo?" + sendData
				reply, err := http.Get(url)
				defer reply.Body.Close()
				if err != nil {
					fmt.Println(err)
				}
				body, _ := ioutil.ReadAll(reply.Body)
				fmt.Println(string(body))
			}
		}
	}()
	for {
		//fmt.Println(computer_num)
	}
}
