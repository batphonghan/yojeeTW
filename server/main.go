package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"yojeeTW/pb"

	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

func main() {
	registerServiceWithConsul()
	lis, err := net.Listen("tcp", port())
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	log.Println("Starting serve at ", port())
	pb.RegisterTweetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "50051"
	}
	return ":" + port
}

func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	var registration = new(consulapi.AgentServiceRegistration)

	registration.ID = "yojee-grpc-server"
	registration.Name = "yojee-grpc-server"

	address := hostname()
	registration.Address = address
	port, _ := strconv.Atoi(port()[1:len(port())])
	registration.Port = port

	consul.Agent().ServiceRegister(registration)
}
