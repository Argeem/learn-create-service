package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {

	var s *grpc.Server

	// รับค่าจาก flag ใน command line โดยจะ return มาเป็น pointer
	tls := flag.Bool("tls", false, "use a secure TLS connection") //name, default value, usage string
	flag.Parse()                                                  // เพื่อให้ใช้งาน flag ได้

	if *tls {
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		// เพื่อใช้เป็นแบบ tls เลยสร้าง credential ดังนี้
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile) // ต้องการ cert ไฟล์และ key ไฟล์
		if err != nil {
			log.Fatal(err)
		}
		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		s = grpc.NewServer()
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())
	reflection.Register(s)

	fmt.Print("gRPC server listening on port 50001")
	if *tls {
		fmt.Println(" with TLS")
	}
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
