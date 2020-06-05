/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/kevin-zx/fxtservice/pkg/config"
	"github.com/kevin-zx/fxtservice/pkg/rpc/fxtservice_rpc"
	"github.com/kevin-zx/fxtservice/pkg/storage/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:11844")
	if err != nil {
		log.Fatal(err)
	}
	//s, err := sqlbase.NewStorage(config.XwDB["ENGINE"].(string), config.XwDB["PORT"].(int), config.XwDB["NAME"].(string), config.XwDB["USER"].(string), config.XwDB["HOST"].(string), config.XwDB["PASSWORD"].(string))
	//if err != nil {
	//	panic(err)
	//}
	//defer s.Close()
	fs, err := mysql.NewFxtStorage(config.XwDB["ENGINE"].(string), config.XwDB["PORT"].(int), config.XwDB["NAME"].(string), config.XwDB["USER"].(string), config.XwDB["HOST"].(string), config.XwDB["PASSWORD"].(string))
	if err != nil {
		panic(err)
	}
	defer fs.Close()
	grpcServer := grpc.NewServer()
	fxtservice_rpc.RegisterFxtServiceServer(grpcServer, fxtservice_rpc.NewFxtRPCService(fs))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server:%v", err)
	}
}
