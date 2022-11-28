package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds)) // client connection : ใช้ติดต่อไปยัง gRPC server
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	//err = calculatorService.Hello("Bond")
	err = calculatorService.Fibonacci(3)
	if err != nil {
		log.Fatal(err)
	}
}
