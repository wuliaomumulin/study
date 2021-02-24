package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack"
)

type s struct{
	data map[string]interface{}

}

type Person struct{
	Name string
	Age int
	Gender string
}

func main(){
	jsonDemo()
	godDemo()
	fmt.Println("----------msgpackDemo------------")
	msgpackDemo()
}
/**
* json序列化的问题
*/
func jsonDemo(){
	var s1 = s{
		data:make(map[string]interface{},8),
	}
	s1.data["count"] = 1
	//序列化
	ret,err := json.Marshal(s1.data)
	if err!=nil {
		fmt.Println("marshal failed",err)
	}
	fmt.Printf("%#v\n",string(ret))

	var s2 = s{
		data:make(map[string]interface{},8),
	}
	err =json.Unmarshal(ret,&s2.data)
	if err!=nil {
		fmt.Println("unmarshal failed",err)
	}
	fmt.Println(s2)
	for _,v := range s2.data{
		fmt.Printf("values:%v,type:%T\n",v,v)
	}
}
/**
标准库gob是golang提供的“私有”的编解码方式，它的效率会比json，xml等更高，特别适合在Go语言程序间传递数据。
*/
func godDemo(){
	var s1 = s{
		data:make(map[string]interface{},8),
	}
	s1.data["count"] =1
	//encode
	buf:=new(bytes.Buffer)
	enc:=gob.NewEncoder(buf)
	err:=enc.Encode(s1.data)
	if err!=nil {
		fmt.Println("gob encode failed",err)
		return
	}
	b:=buf.Bytes()
	fmt.Println(b)
	var s2 = s{
		data:make(map[string]interface{},8),
	}
	//decode
	dec:=gob.NewDecoder(bytes.NewBuffer(b))
	err =dec.Decode(&s2.data)
	if err!= nil {
		fmt.Println("gob decode failed error:",err)
		return
	}
	fmt.Println(s2.data)
	for _,v:=range s2.data {
		fmt.Printf("值为%v,类型为%T",v,v)
	}
}
func msgpackDemo() {
	p1 := Person{
		Name:   "沙赫扎奥",
		Age:    119,
		Gender: "未知",
	}
	//序列化
	b, err := msgpack.Marshal(p1)
	if err != nil {
		fmt.Printf("msgpack failed error:%v", err)
		return
	}
	//反序列化
	var p2 Person
	err = msgpack.Unmarshal(b, &p2)
	if err!=nil {
		fmt.Printf("msgpack unmarsahal failed err:%v",err)
		return
	}
	fmt.Printf("p2:%#v\n",p2)
	//p2:main.Person{Name:"沙赫扎奥", Age:119, Gender:"未知"}
}