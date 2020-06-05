package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevin-zx/fxtservice/pkg/rpc/fxtservice_rpc"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:11844", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	c := fxtservice_rpc.NewFxtServiceClient(conn)
	//getGuanbaoProject(c)
	getGuanwangProject(c)
}

func getGuanwangProject(c fxtservice_rpc.FxtServiceClient) {
	d, err := c.GetGuanwangProject(context.Background(), &fxtservice_rpc.ProjectRequest{
		SiteUrl:    "www.surpon.com",
		ClientName: "苏州迅鹏仪器仪表有限公司",
	})
	if err != nil {
		log.Fatalf("getguanbaoproject err :%v", err)
	}
	data, err := json.Marshal(d)
	if err != nil {
		log.Fatalf("json encode err:%v", d)
	}
	fmt.Println(string(data))
}

func getGuanbaoProject(c fxtservice_rpc.FxtServiceClient) {
	d, err := c.GetGuanbaoProject(context.Background(), &fxtservice_rpc.ProjectRequest{
		SiteUrl:    "www.njhdbfc.com",
		ClientName: "南京市江宁区博文活动板房厂",
	})
	if err != nil {
		log.Fatalf("getguanbaoproject err :%v", err)
	}
	data, err := json.Marshal(d)
	if err != nil {
		log.Fatalf("json encode err:%v", d)
	}
	fmt.Println(string(data))
}
