package main

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	avatar "rpc-hello/avatarRpc"
)

// avatarServer 用来实现 avatar.AvatarMovieServer

type avatarServer struct{}

// 实现 avatar.GetAvatarMovie 方法

// (context.Context, *AvatarMovieRequest) (*AvatarMovieReply, error)

func (s *avatarServer) GetAvatarMovie(ctx context.Context, in *avatar.AvatarMovieRequest) (*avatar.AvatarMovieResponse, error) {
	result := map[string]interface{}{
		"status:": 0,
		"ui_type": map[string]interface{}{
			"name": "block_list",
		},
		"meta": map[string]interface{}{
			"more": "more",
		},
	}

	ret := new(avatar.Output)
	ret.Data = &structpb.Value{Kind: &structpb.Value_StructValue{StructValue: avatar.ToStruct(result)}}
	movies:=avatar.Movies{Ids:in.Input[0].Movieid}
	moviemap:=make(map[string]*avatar.Movies,0)
	moviemap["ids"]=&movies
	ret.Foo=moviemap
	ret.Code=200
	return &avatar.AvatarMovieResponse{Output:ret}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	avatar.RegisterAvatarMovieServer(s, &avatarServer{})

	//在 server 中 注册 gRPC 的 reflection service

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

