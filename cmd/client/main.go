package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevin-zx/fxtservice/pkg/rpc/fxtservice_rpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:11844", grpc.WithInsecure())
	//conn, err := grpc.Dial("45.40.251.69:11844", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()
	c := fxtservice_rpc.NewFxtServiceClient(conn)

	getGuanbaoProject(c)
	getGuanwangProject(c)
}

func getGuanwangProject(c fxtservice_rpc.FxtServiceClient) {
	d, err := c.GetGuanwangProject(context.Background(), &fxtservice_rpc.ProjectRequest{
		SiteUrl:    "www.1618.tech",
		ClientName: "苏州易尔邦自动化设备",
	})
	if err != nil {
		log.Fatalf("getguanbaoproject err :%v", err)
	}
	for _, project := range d.GuanwangProjects {
		for _, keyword := range project.SiteKeywords {
			for _, rank := range keyword.SiteKeywordRanks {

				fmt.Println(time.Unix(rank.Date,0).Format("2006-01-02"))
			}
		}
	}
	data, err := json.Marshal(d)
	if err != nil {
		log.Fatalf("json encode err:%v", d)
	}
	fmt.Println(string(data))
}

func getGuanbaoProject(c fxtservice_rpc.FxtServiceClient) {
	d, err := c.GetGuanbaoProject(context.Background(), &fxtservice_rpc.ProjectRequest{
		SiteUrl:    "www.1618.tech",
		ClientName: "苏州易尔邦自动化设备",
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
