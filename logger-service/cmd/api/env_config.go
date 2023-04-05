package main

import "os"

type IEnv struct {
	PORT     string
	mongoURI string
	rpcPORT  string
	grpcPORT string
}

var env = IEnv{
	PORT:     os.Getenv("PORT"),
	mongoURI: os.Getenv("mongoURI"),
	rpcPORT:  os.Getenv("RPC_PORT"),
	grpcPORT: os.Getenv("GRPC_PORT"),
}
