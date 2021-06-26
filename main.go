package main

import (
	"context"
	"fmt"
	"grpcintro/github.com/eduFDiaz/grpc_intro"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// protoc -I=./ --go_out=./ --go-grpc_out=./ ./addressbook.proto

var (
	p = grpc_intro.Person{
		Name:        "Eduardo Fernandez",
		Id:          1,
		Email:       "fernandez9000@gmail.com",
		LastUpdated: timestamppb.Now(),
		Phones: []*grpc_intro.Person_PhoneNumber{
			{
				Type:   grpc_intro.Person_HOME,
				Number: "(469) 831-4596",
			},
			{
				Type:   grpc_intro.Person_MOBILE,
				Number: "(469) 831-4111",
			},
		},
	}

	a = &grpc_intro.AddressBook{
		People: []*grpc_intro.Person{
			&p,
			{
				Name:        "Andy Garcia Fernandez",
				Id:          1,
				Email:       "andygarcia@gmail.com",
				LastUpdated: timestamppb.Now(),
				Phones: []*grpc_intro.Person_PhoneNumber{
					{
						Type:   grpc_intro.Person_HOME,
						Number: "(469) 831-4591",
					},
					{
						Type:   grpc_intro.Person_MOBILE,
						Number: "(469) 831-4113",
					},
				},
			},
		},
	}
)

type server struct {
	grpc_intro.UnimplementedAddressBookServiceServer
}

func (*server) GetAllAddresses(ctx context.Context, reques *grpc_intro.Empty) (*grpc_intro.AddressBook, error) {
	return a, nil
}

func (*server) GetPersonByName(ctx context.Context, reques *grpc_intro.ReqGetPerson) (*grpc_intro.Person, error) {
	for _, v := range a.People {
		if v.Name == reques.Name {
			return v, nil
		}
	}
	return nil, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	grpc_intro.RegisterAddressBookServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
