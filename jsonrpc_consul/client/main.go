package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"rpc/jsonrpc_consul/client/consul"
)

var serverAddress string

type Args struct {
	A int
	B int
}

/*type Args struct {
	A int "json:`A`"
	B int "json:`B`"
}*/
type Quotient struct {
	Quo, Rem int
}

// TODO: [客户端发现服务] - prince_todo 2021/6/26 下午7:19
var ServerByConsul interface{}

func main() {
	// 从 consul 中发现服务
	serverAddress, err := consul.AgentHealthServiceByName()
	if err != nil {
		log.Fatal("获取服务地址失败:", err)
	}
	fmt.Println("获取服务地址：", serverAddress, err)
	
	client, err := jsonrpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
		return
	}
	defer client.Close()

	// Synchronous call: 同步调用
	args := &Args{7, 8}
	var reply int
	err = client.Call("arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d \n", reply)
	// fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply)

	// Asynchronous call: 异步调用
	quotient := new(Quotient)
	divCall := client.Go("arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done	// will be equal to divCall
	// check errors, print, etc.
	fmt.Printf("replyCall: %v %v %v", *replyCall, replyCall.Args, replyCall.Reply)
}
