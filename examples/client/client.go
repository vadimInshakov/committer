package main

import (
	"github.com/vadiminshakov/committer/client"
	pb "github.com/vadiminshakov/committer/proto"
)
func main(){
	cli, err := client.New("localhost:3000")
	if err != nil {
		panic(err)
	}
	resp, err := cli.Put("1", []byte("2"))
	if err != nil {
		panic(err)
	}
	if resp.Type != pb.Type_ACK {
		panic("msg is not acknowledged")
	}
}