// Package grpc provides a gRPC client for the Users service.
package grpc

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/adamryman/ambition-users/users-service"
	svc "github.com/adamryman/ambition-users/users-service/generated"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.UsersServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var createuserEndpoint endpoint.Endpoint
	{
		createuserEndpoint = grpctransport.NewClient(
			conn,
			"users.Users",
			"CreateUser",
			svc.EncodeGRPCCreateUserRequest,
			svc.DecodeGRPCCreateUserResponse,
			pb.User{},
			clientOptions...,
		).Endpoint()
	}

	var readuserEndpoint endpoint.Endpoint
	{
		readuserEndpoint = grpctransport.NewClient(
			conn,
			"users.Users",
			"ReadUser",
			svc.EncodeGRPCReadUserRequest,
			svc.DecodeGRPCReadUserResponse,
			pb.User{},
			clientOptions...,
		).Endpoint()
	}

	var updateuserEndpoint endpoint.Endpoint
	{
		updateuserEndpoint = grpctransport.NewClient(
			conn,
			"users.Users",
			"UpdateUser",
			svc.EncodeGRPCUpdateUserRequest,
			svc.DecodeGRPCUpdateUserResponse,
			pb.User{},
			clientOptions...,
		).Endpoint()
	}

	var deleteuserEndpoint endpoint.Endpoint
	{
		deleteuserEndpoint = grpctransport.NewClient(
			conn,
			"users.Users",
			"DeleteUser",
			svc.EncodeGRPCDeleteUserRequest,
			svc.DecodeGRPCDeleteUserResponse,
			pb.User{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		CreateUserEndpoint: createuserEndpoint,
		ReadUserEndpoint:   readuserEndpoint,
		UpdateUserEndpoint: updateuserEndpoint,
		DeleteUserEndpoint: deleteuserEndpoint,
	}, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.RequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
