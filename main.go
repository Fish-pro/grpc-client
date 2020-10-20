package main

import (
	"context"
	"fmt"
	"github.com/Fish-pro/grpc-client/helper"
	"github.com/Fish-pro/grpc-client/services"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
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

	//orderClient := services.NewOrderServiceClient(conn) // 新建订单服务客户端
	//
	//// 新建订单
	//t := timestamp.Timestamp{Seconds: time.Now().Unix()}
	//order := services.OrderMain{
	//	OrderId:    1,
	//	OrderMoney: 20.5,
	//	OrderNo:    "23423423",
	//	OrderTime:  &t,
	//}
	//response, err := orderClient.NewOrder(context.Background(), &order)
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	//fmt.Println(response)

	userClient := services.NewUserServiceClient(conn)

	// 普通方法获取
	//var i int32
	//req := services.UserScoreRequest{}
	//req.Users = make([]*services.UserInfo, 0)
	//
	//for i = 1; i < 6; i++ {
	//	req.Users = append(req.Users, &services.UserInfo{UserId: i})
	//}
	//
	//response, err := userClient.GetUserScore(context.Background(), &req)
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	//fmt.Println(response.Users)

	// 服务端流式
	//var i int32
	//req := services.UserScoreRequest{}
	//req.Users = make([]*services.UserInfo, 0)
	//
	//for i = 1; i < 6; i++ {
	//	req.Users = append(req.Users, &services.UserInfo{UserId: i})
	//}
	//
	//stream, err := userClient.GetUserScoreByServerStream(context.Background(), &req)
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	//
	//for {
	//	res, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Println(err.Error())
	//		os.Exit(1)
	//	}
	//	fmt.Println(res.Users)
	//}

	// 客户端流式
	//var i int32
	//
	//stream, err := userClient.GetUserScoreByClientStream(context.Background())
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	//
	//for j := 1; j <= 3; j++ {
	//	req := services.UserScoreRequest{}
	//	req.Users = make([]*services.UserInfo, 0)
	//	for i = 1; i < 6; i++ {
	//		req.Users = append(req.Users, &services.UserInfo{UserId: i})
	//	}
	//	err = stream.Send(&req)
	//	if err != nil {
	//		log.Println(err.Error())
	//		os.Exit(1)
	//	}
	//}
	//response, err := stream.CloseAndRecv()
	//if err != nil {
	//	log.Println(err.Error())
	//	os.Exit(1)
	//}
	//fmt.Println(response.Users)

	// 双向流式
	stream, err := userClient.GetUserScoreByTWS(context.Background())
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	var uid int32 = 1
	for j := 1; j <= 3; j++ {
		req := services.UserScoreRequest{}
		req.Users = make([]*services.UserInfo, 0)
		for i := 1; i < 6; i++ {
			req.Users = append(req.Users, &services.UserInfo{UserId: uid})
			uid++
		}
		err = stream.Send(&req)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}

		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(res.Users)
	}

}
