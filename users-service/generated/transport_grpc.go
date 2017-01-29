package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/adamryman/ambition-users/users-service"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC UsersServer.
func MakeGRPCServer(ctx context.Context, endpoints Endpoints) pb.UsersServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	return &grpcServer{
		// users

		createuser: grpctransport.NewServer(
			ctx,
			endpoints.CreateUserEndpoint,
			DecodeGRPCCreateUserRequest,
			EncodeGRPCCreateUserResponse,
			serverOptions...,
		),
		readuser: grpctransport.NewServer(
			ctx,
			endpoints.ReadUserEndpoint,
			DecodeGRPCReadUserRequest,
			EncodeGRPCReadUserResponse,
			serverOptions...,
		),
		updateuser: grpctransport.NewServer(
			ctx,
			endpoints.UpdateUserEndpoint,
			DecodeGRPCUpdateUserRequest,
			EncodeGRPCUpdateUserResponse,
			serverOptions...,
		),
		deleteuser: grpctransport.NewServer(
			ctx,
			endpoints.DeleteUserEndpoint,
			DecodeGRPCDeleteUserRequest,
			EncodeGRPCDeleteUserResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the UsersServer interface
type grpcServer struct {
	createuser grpctransport.Handler
	readuser   grpctransport.Handler
	updateuser grpctransport.Handler
	deleteuser grpctransport.Handler
}

// Methods for grpcServer to implement UsersServer interface

func (s *grpcServer) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	_, rep, err := s.createuser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.User), nil
}

func (s *grpcServer) ReadUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	_, rep, err := s.readuser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.User), nil
}

func (s *grpcServer) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	_, rep, err := s.updateuser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.User), nil
}

func (s *grpcServer) DeleteUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	_, rep, err := s.deleteuser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.User), nil
}

// Server Decode

// DecodeGRPCCreateUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createuser request to a user-domain createuser request. Primarily useful in a server.
func DecodeGRPCCreateUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.User)
	return req, nil
}

// DecodeGRPCReadUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC readuser request to a user-domain readuser request. Primarily useful in a server.
func DecodeGRPCReadUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.User)
	return req, nil
}

// DecodeGRPCUpdateUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC updateuser request to a user-domain updateuser request. Primarily useful in a server.
func DecodeGRPCUpdateUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.User)
	return req, nil
}

// DecodeGRPCDeleteUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC deleteuser request to a user-domain deleteuser request. Primarily useful in a server.
func DecodeGRPCDeleteUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.User)
	return req, nil
}

// Client Decode

// DecodeGRPCCreateUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createuser reply to a user-domain createuser response. Primarily useful in a client.
func DecodeGRPCCreateUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.User)
	return reply, nil
}

// DecodeGRPCReadUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC readuser reply to a user-domain readuser response. Primarily useful in a client.
func DecodeGRPCReadUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.User)
	return reply, nil
}

// DecodeGRPCUpdateUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC updateuser reply to a user-domain updateuser response. Primarily useful in a client.
func DecodeGRPCUpdateUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.User)
	return reply, nil
}

// DecodeGRPCDeleteUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC deleteuser reply to a user-domain deleteuser response. Primarily useful in a client.
func DecodeGRPCDeleteUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.User)
	return reply, nil
}

// Server Encode

// EncodeGRPCCreateUserResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain createuser response to a gRPC createuser reply. Primarily useful in a server.
func EncodeGRPCCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.User)
	return resp, nil
}

// EncodeGRPCReadUserResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain readuser response to a gRPC readuser reply. Primarily useful in a server.
func EncodeGRPCReadUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.User)
	return resp, nil
}

// EncodeGRPCUpdateUserResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain updateuser response to a gRPC updateuser reply. Primarily useful in a server.
func EncodeGRPCUpdateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.User)
	return resp, nil
}

// EncodeGRPCDeleteUserResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain deleteuser response to a gRPC deleteuser reply. Primarily useful in a server.
func EncodeGRPCDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.User)
	return resp, nil
}

// Client Encode

// EncodeGRPCCreateUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createuser request to a gRPC createuser request. Primarily useful in a client.
func EncodeGRPCCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.User)
	return req, nil
}

// EncodeGRPCReadUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain readuser request to a gRPC readuser request. Primarily useful in a client.
func EncodeGRPCReadUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.User)
	return req, nil
}

// EncodeGRPCUpdateUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain updateuser request to a gRPC updateuser request. Primarily useful in a client.
func EncodeGRPCUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.User)
	return req, nil
}

// EncodeGRPCDeleteUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain deleteuser request to a gRPC deleteuser request. Primarily useful in a client.
func EncodeGRPCDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.User)
	return req, nil
}

// Helpers

func metadataToContext(ctx context.Context, md *metadata.MD) context.Context {
	for k, v := range *md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
