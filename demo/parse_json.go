package main

import (
	"encoding/json"
	. "github.com/bitly/go-simplejson"
)

func main() {
	var s ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	// 通过地址传递，这样就可以对原始赋值了
	json.Unmarshal([]byte(str), &s)
	//fmt.Println(s)

	js, _ := NewJson([]byte(`{
	"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
	}
}`))
	i, _ := js.Get("test").Get("int").Int()
	println(i)
	//println(err)
}

type Server struct {
	ServerName string

	ServerIP string
}

type ServerSlice struct {
	Servers []Server
}
