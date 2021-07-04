package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}
type Arith int

const listenAddress = "0.0.0.0:1234"

func main() {
	/** go 服务需要注册是因为 go 是强语言，没办法像 php 一样字符串加双括号就可以作为方法「字符串()」 */
	arith := new(Arith)
	rpc.Register(arith)

	rpc.HandleHTTP()

	l, e := net.Listen("tcp", listenAddress)
	log.Println("listen at ", listenAddress)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()

	http.Serve(l, nil)
	/*go http.Serve(l, nil) // 堵塞后，主协程退出
	time.Sleep(100000 * time.Second) // 防止主协程退出，仅用于调试*/
}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
