package main

import (
	"context"
	"fmt"
	desc "github.com/EviL345/email_sender/pkg/send_email_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"net/smtp"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedSendEmailServer
}

func (s *server) Send(ctx context.Context, req *desc.SendEmailRequest) (*emptypb.Empty, error) {
	auth := smtp.PlainAuth("", "extremegam1ng645@gmail.com", "nslr ufar ksrt jaet", "smtp.gmail.com")

	msg := fmt.Sprintf("Subject: %s\n%s", req.GetSubject(), req.GetBody())

	if err := smtp.SendMail("smtp.gmail.com:587", auth, "extremegam1ng645@gmail.com", []string{req.GetRecipient()}, []byte(msg)); err != nil {
		return nil, err
	}

	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterSendEmailServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
