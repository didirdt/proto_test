package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	pb "github.com/didirdt/proto_test/protobuff"
	"github.com/randyardiansyah25/gostashlg"
	"google.golang.org/grpc"
)

func LogWrite(logstring string) {
	event := "LOG"
	detail := "\nDetail:" + logstring + "\n\n"
	lgs := "\n" + logstring
	log, _ := gostashlg.UseDefault()
	field := gostashlg.NewFields().
		SetLevel(gostashlg.INFO).
		SetEvent(event).
		SetMessage(lgs).
		SetData(detail).
		Get()

	log.Write(field)
}

type TransactionService struct {
	pb.UnimplementedTransactionServiceServer
	mu       sync.Mutex
	messages *pb.Message
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) Inquiry(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	msg.ResponseCode = "0000"
	msg.MsgResponse = "Inquiry"
	return msg, nil
}

func (s *TransactionService) Transaction(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	msg.ResponseCode = "0000"
	msg.MsgResponse = "Transaction"
	return msg, nil
}

func (s *TransactionService) Reversal(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	msg.ResponseCode = "0000"
	msg.MsgResponse = "Reversal"
	return msg, nil
}

func (s *TransactionService) Advice(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	msg.ResponseCode = "0000"
	msg.MsgResponse = "Advice"
	return msg, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprint(":", "9090"))
	if err != nil {
		LogWrite(fmt.Sprintf("Error listening TCP Server: %s", err))
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterTransactionServiceServer(grpcServer, NewTransactionService())

	LogWrite("gRPC Server is running on port : 9090")
	if err := grpcServer.Serve(listener); err != nil {
		LogWrite(fmt.Sprintf("Error starting gRPC Server: %s", err))
	}
}
