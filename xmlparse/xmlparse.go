package xmlparse

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Remote_monitor struct {
	Ischeck     int     `xml:"ischeck,attr"`
	Timeout     float32 `xml:"timeout,attr"`
	Remote_item string  `xml:"remote_item"`
}
type Delegation_serivce struct {
	Isdelegation int     `xml:"isdelegation,attr"`
	Timeout      float32 `xml:"timeout,attr"`
	Dist_target  string  `xml:"dist_target"`
}
type Reseult struct {
	XMLname            xml.Name             `xml:"computer"`
	Type               int                  `xml:"type"`
	While_time         float32              `xml:"while_time"`
	Coll_proce         string               `xml:"coll_proce"`
	Remote_monitor     []Remote_monitor     `xml:"remote_monitor"`
	Request_address    string               `xml:"request_address"`
	Delegation_serivce []Delegation_serivce `xml:"delegation_serivce"`
	Version            string               `xml:"version"`
	Client_timeout     int                  `xml:"client_timeout"`
	Client_num         int                  `xml:"client_num"`
}

func readXml(path string) (result string) {
	fileinfo, err := os.Lstat(path)
	if err != nil {
		fmt.Println(err)
		result = ""
		return
	}
	// fmt.Println(fileinfo.Name(), fileinfo.IsDir())
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		result = ""
		return
	}
	buf := make([]byte, fileinfo.Size())
	for {
		n, err := file.Read(buf)
		result += string(buf)
		if err == io.EOF && n == 0 {
			break
		}
	}
	file.Close()
	// fmt.Println(result)
	return
}
func GetConfig(data *Reseult) (ok bool) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		ok = false
		return
	}
	path += "\\config.xml"
	xmldata := readXml(path)
	err = xml.Unmarshal([]byte(xmldata), data)
	if err != nil {
		fmt.Println(err)
		ok = false
		return
	}
	ok = true
	return
}
