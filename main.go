package main

import (
	"context"
	"fmt"
	"github.com/Fish-pro/grpc-client/helper"
	"github.com/Fish-pro/grpc-client/services"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	//cred, err := credentials.NewClientTLSFromFile("keys/server.crt", "grpcserver.com")
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	cred := helper.GetClientCred()

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	//prodClient := services.NewProdServiceClient(conn) // 新建商品服务客户端

	// 获取商品信息
	//response, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12, ProdArea: services.ProdAreas_A})
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	//fmt.Println(response.ProdStock)

	// 获取商品详情
	//response, err := prodClient.GetProdInfo(context.Background(), &services.ProdRequest{ProdId: 12, ProdArea: services.ProdAreas_A})
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	//fmt.Println(response)

	// 获取商品列表
	//response, err := prodClient.GetProdStocks(context.Background(), &services.QuerySize{Size: 10})
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	//fmt.Println(response.Prodres)

	orderClient := services.NewOrderServiceClient(conn) // 新建订单服务客户端

	// 新建订单
	t := timestamp.Timestamp{Seconds: time.Now().Unix()}
	order := services.OrderMain{
		OrderId:    1,
		OrderMoney: 20.5,
		OrderNo:    "23423423",
		OrderTime:  &t,
	}
	response, err := orderClient.NewOrder(context.Background(), &order)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(response)

}
