package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const serverAddress string = "127.0.0.1:1234"

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()
	
	// Synchronous call: 同步调用
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply)
	
	// Asynchronous call: 异步调用
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done	// will be equal to divCall
	// check errors, print, etc.
	fmt.Printf("replyCall: %v %v %v", replyCall, replyCall.Args, replyCall.Reply)
}
