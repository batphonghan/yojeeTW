package main

import (
	"fmt"
	"log"
	"yojeeTW/pb"

	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

var client pb.TweetServiceClient

func lookupServiceWithConsul() string {
	log.Println("Lookup Service With Consul")

	config := consulapi.DefaultConfig()
	consul, error := consulapi.NewClient(config)
	if error != nil {
		fmt.Println(error)
	}

	services, error := consul.Agent().Services()
	if error != nil {
		fmt.Println(error)
	}

	service := services["yojee-grpc-server"]
	address := service.Address
	port := service.Port

	url := fmt.Sprintf("%s:%v", address, port)

	discoveryURL = url
	log.Printf("Lookup Service [%s] \n", url)
	return url
}

var discoveryURL string

func initClient() {
	url := lookupServiceWithConsul()
	log.Println("====================================================")
	log.Printf("Found Service With Consul at [%s] \n", url)
	log.Println("====================================================")
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to GRPC server: [%v]", url)
	}

	client = pb.NewTweetServiceClient(conn)
}
