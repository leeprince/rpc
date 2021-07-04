package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A int
	B int
}
type Quotient struct {
	Quo, Rem int
}
type Arith int

const listenAddress = "127.0.0.1:12345"

func main() {
	l, e := net.Listen("tcp", listenAddress)
	log.Println("listen at ", listenAddress)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()

	/** go 服务需要注册是因为 go 是强语言，没办法像 php 一样字符串加双括号就可以作为方法「字符串()」 */
	rpc.RegisterName("arith", new(Arith))

	// conn, err := l.Accept()
	// if err != nil {
	// 	log.Fatal("Accept error:", err)
	// }
	// jsonrpc.ServeConn(conn)
	/*go jsonrpc.ServeConn(conn) // 堵塞后，主协程退出
	time.Sleep(100000 * time.Second)// 防止主协程退出，仅用于调试 */
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
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
