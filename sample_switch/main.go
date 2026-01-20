package main

import (
	"context"
	"fmt"

	pb "github.com/didirdt/proto_test/protobuff"
	"github.com/randyardiansyah25/gostashlg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func LogWrite(logstring string) {
	event := "LOG"
	detail := "\nDetail: " + logstring + "\n\n"
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

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("127.0.0.1:9090", opts...)
	if err != nil {
		LogWrite(fmt.Sprintf("failed to connect: %s", err))
	}
	defer conn.Close()

	client := pb.NewTransactionServiceClient(conn)
	LogWrite("gRPC Client connected to :9090")
	_ = client

	msg := &pb.Message{
		Responsecode: "1111",
		Msgresponse:  "Send Message",
	}
	resInquiry, err := Inquiry(client, msg)
	if err != nil {
		LogWrite(fmt.Sprintf("Error calling Inquiry: %s", err))
	}
	LogWrite(fmt.Sprintf("Inquiry Response: %+v", resInquiry))

	resTransaction, err := Transaction(client, msg)
	if err != nil {
		LogWrite(fmt.Sprintf("Error calling Transaction: %s", err))
	}
	LogWrite(fmt.Sprintf("Transaction Response: %+v", resTransaction))

	resReversal, err := Reversal(client, msg)
	if err != nil {
		LogWrite(fmt.Sprintf("Error calling Reversal: %s", err))
	}
	LogWrite(fmt.Sprintf("Reversal Response: %+v", resReversal))

	resAdvice, err := Advice(client, msg)
	if err != nil {
		LogWrite(fmt.Sprintf("Error calling Advice: %s", err))
	}
	LogWrite(fmt.Sprintf("Advice Response: %+v", resAdvice))
}

func Inquiry(client pb.TransactionServiceClient, msg *pb.Message) (*pb.Message, error) {
	return client.Inquiry(context.Background(), msg)
}
func Transaction(client pb.TransactionServiceClient, msg *pb.Message) (*pb.Message, error) {
	return client.Transaction(context.Background(), msg)
}
func Reversal(client pb.TransactionServiceClient, msg *pb.Message) (*pb.Message, error) {
	return client.Reversal(context.Background(), msg)
}
func Advice(client pb.TransactionServiceClient, msg *pb.Message) (*pb.Message, error) {
	return client.Advice(context.Background(), msg)
}
