package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ServerRequest struct {
	A int
	B int
}
type ServerResponse struct {
	C int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatalln("连接失败")
	}
	req := ServerRequest{1, 5}
	var res ServerResponse
	err = conn.Call("Calculator.Add", req, &res)
	if err != nil {
		log.Fatalln("调用失败")
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.C)
}
