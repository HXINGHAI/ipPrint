package main

import (
	"net"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"strings"
)

/**
  @Auth: H_XING海
  @Date: 2019.6.26
  @Description: IP域名核查检测
 */
 
func checked(data map[string]string)(errorResult map[string]string,okResult map[string]string){
	errorResult = make(map[string]string,0)
	okResult = make(map[string]string,0)
	for k,v:= range data {
		conn,err := net.Dial("ip:icmp",k)
		if err != nil {
			errorResult[k] = "No such host !" +"@" + v
		}else {
			ipAddress := conn.RemoteAddr().String()
			if ipAddress != v {
				errorResult[k] = ipAddress + "@" + v
			}else{
				okResult[k] = ipAddress + "@" + v
			}
		}
	}
	return errorResult,okResult
}

func main(){

	file,err := ioutil.ReadFile("host-ip.txt")
	if err != nil {
		fmt.Println("文件名不正确，请改为 host-ip,格式为txt")
	}
	ipMap := map[string]string{}
	err = json.Unmarshal([]byte(string(file)),&ipMap)
	if err != nil {
		fmt.Println("文件内容格式不正确！（改为json格式）")
	}
	errorResult := make(map[string]string)
	okResult := make(map[string]string)
	errorResult,okResult = checked(ipMap)
	fmt.Println("以下IP域名信息对应不正确（存在问题）：")
	for k,v := range errorResult{
		var ipstring []string
		ipstring = strings.Split(v,"@")
		fmt.Printf("ERROR------域名为：%s  正确ip为：%s  您传入的IP为：%s\n",k,ipstring[0],ipstring[1])
	}
	fmt.Println("以下IP域名信息相互对应：")
	for k,v := range okResult {
		var ipstring []string
		ipstring = strings.Split(v,"@")
		fmt.Printf("OK------域名为：%s  正确ip为：%s  您传入的IP为：%s\n",k,ipstring[0],ipstring[1])
	}
	fmt.Println("press anykey stop!")
	fmt.Scanf("%s")
}
