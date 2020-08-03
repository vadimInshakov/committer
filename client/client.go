package client

import (
	"github.com/pkg/errors"
	pb "github.com/vadiminshakov/committer/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type CommitClient struct {
	Connection pb.CommitClient
}

func New(addr string) (*CommitClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	return &CommitClient{pb.NewCommitClient(conn)}, nil
}

func (client *CommitClient) Propose(*pb.ProposeRequest) (*pb.Response, error) {
	return client.Connection.Propose(context.Background(), &pb.ProposeRequest{Key: "hi", Value: []byte("hello"), CommitType: pb.CommitType_TWO_PHASE_COMMIT, Index: 1})
}

func (client *CommitClient) Precommit(*pb.PrecommitRequest) (*pb.Response, error) {
	return nil, nil
}

func (client *CommitClient) Commit(*pb.CommitRequest) (*pb.Response, error) {
	return nil, nil
}

func (client *CommitClient) Put(key string, value []byte) (*pb.Response, error) {
	return client.Connection.Put(context.Background(), &pb.Entry{Key: key, Value: value})
}