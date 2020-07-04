package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Calculator struct{}
type ServerRequest struct {
	A int
	B int
}
type ServerResponse struct {
	C int
}

func (c *Calculator) Add(req ServerRequest, res *ServerResponse) error {
	res.C = req.A * req.B
	return nil
}
func main() {
	rpc.Register(new(Calculator))
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal("监听失败")
	}
	http.Serve(listener, nil)
}
