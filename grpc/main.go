package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)



func main(){
	fmt.Print("hello world")
	elliot := &Person{
		Name: "Elliot",
		Age: 24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshalling error",err)
	}
	fmt.Println(data)
	newElliot := &Person{}
	err =proto.Unmarshal(data,newElliot)
	if err!= nil{
		log.Fatal("un marshalling error")
	}
	fmt.Print(newElliot.GetAge())
	fmt.Print(newElliot.GetName())
}
