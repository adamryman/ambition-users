package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pkg/errors"

	// This Service
	pb "github.com/adamryman/ambition-users/users-service"
	clientHandler "github.com/adamryman/ambition-users/users-service/generated/cli/handlers"
	grpcclient "github.com/adamryman/ambition-users/users-service/generated/client/grpc"
	httpclient "github.com/adamryman/ambition-users/users-service/generated/client/http"
)

var (
	_ = strconv.ParseInt
	_ = strings.Split
	_ = json.Compact
	_ = errors.Wrapf
	_ = pb.RegisterUsersServer
)

func main() {
	// The addcli presumes no service discovery system, and expects users to
	// provide the direct address of an service. This presumption is reflected in
	// the cli binary and the the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings.

	var (
		httpAddr = flag.String("http.addr", "", "HTTP address of addsvc")
		grpcAddr = flag.String("grpc.addr", ":5040", "gRPC (HTTP) address of addsvc")
		method   = flag.String("method", "createuser", "createuser,readuser,updateuser,deleteuser")
	)

	var (
		flagIDUpdateUser     = flag.Int64("updateuser.id", 0, "")
		flagInfoUpdateUser   = flag.String("updateuser.info", "", "")
		flagTrelloUpdateUser = flag.String("updateuser.trello", "", "")
		flagIDDeleteUser     = flag.Int64("deleteuser.id", 0, "")
		flagInfoDeleteUser   = flag.String("deleteuser.info", "", "")
		flagTrelloDeleteUser = flag.String("deleteuser.trello", "", "")
		flagIDCreateUser     = flag.Int64("createuser.id", 0, "")
		flagInfoCreateUser   = flag.String("createuser.info", "", "")
		flagTrelloCreateUser = flag.String("createuser.trello", "", "")
		flagIDReadUser       = flag.Int64("readuser.id", 0, "")
		flagInfoReadUser     = flag.String("readuser.info", "", "")
		flagTrelloReadUser   = flag.String("readuser.trello", "", "")
	)
	flag.Parse()

	var (
		service pb.UsersServer
		err     error
	)
	if *httpAddr != "" {
		service, err = httpclient.New(*httpAddr)
	} else if *grpcAddr != "" {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while dialing grpc connection: %v", err)
			os.Exit(1)
		}
		defer conn.Close()
		service, err = grpcclient.New(conn)
	} else {
		fmt.Fprintf(os.Stderr, "error: no remote address specified\n")
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch *method {

	case "createuser":
		var err error
		IDCreateUser := *flagIDCreateUser

		var InfoCreateUser pb.UserInfo
		if flagInfoCreateUser != nil && len(*flagInfoCreateUser) > 0 {
			err = json.Unmarshal([]byte(*flagInfoCreateUser), &InfoCreateUser)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling InfoCreateUser from %v:", flagInfoCreateUser))
			}
		}

		var TrelloCreateUser pb.TrelloInfo
		if flagTrelloCreateUser != nil && len(*flagTrelloCreateUser) > 0 {
			err = json.Unmarshal([]byte(*flagTrelloCreateUser), &TrelloCreateUser)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling TrelloCreateUser from %v:", flagTrelloCreateUser))
			}
		}

		request, err := clientHandler.CreateUser(IDCreateUser, InfoCreateUser, TrelloCreateUser)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.CreateUser: %v\n", err)
			os.Exit(1)
		}

		v, err := service.CreateUser(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.CreateUser: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IDCreateUser, InfoCreateUser, TrelloCreateUser)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "readuser":
		var err error
		IDReadUser := *flagIDReadUser

		var InfoReadUser pb.UserInfo
		if flagInfoReadUser != nil && len(*flagInfoReadUser) > 0 {
			err = json.Unmarshal([]byte(*flagInfoReadUser), &InfoReadUser)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling InfoReadUser from %v:", flagInfoReadUser))
			}
		}

		var TrelloReadUser pb.TrelloInfo
		if flagTrelloReadUser != nil && len(*flagTrelloReadUser) > 0 {
			err = json.Unmarshal([]byte(*flagTrelloReadUser), &TrelloReadUser)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling TrelloReadUser from %v:", flagTrelloReadUser))
			}
		}

		request, err := clientHandler.ReadUser(IDReadUser, InfoReadUser, TrelloReadUser)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.ReadUser: %v\n", err)
			os.Exit(1)
		}

		v, err := service.ReadUser(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.ReadUser: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IDReadUser, InfoReadUser, TrelloReadUser)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "updateuser":
		var err error
		IDUpdateUser := *flagIDUpdateUser

		var InfoUpdateUser pb.UserInfo
		if flagInfoUpdateUser != nil && len(*flagInfoUpdateUser) > 0 {
			err = json.Unmarshal([]byte(*flagInfoUpdateUser), &InfoUpdateUser)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling InfoUpdateUser from %v:", flagInfoUpdateUser))
			}
		}

		var TrelloUpdateUser pb.TrelloInfo
		if flagTrelloUpdateUser != nil && len(*flagTrelloUpdateUser) > 0 {
			err = json.Unmarshal([]byte(*flagTrelloUpdateUser), &TrelloUpdateUser)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling TrelloUpdateUser from %v:", flagTrelloUpdateUser))
			}
		}

		request, err := clientHandler.UpdateUser(IDUpdateUser, InfoUpdateUser, TrelloUpdateUser)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.UpdateUser: %v\n", err)
			os.Exit(1)
		}

		v, err := service.UpdateUser(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.UpdateUser: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IDUpdateUser, InfoUpdateUser, TrelloUpdateUser)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "deleteuser":
		var err error
		IDDeleteUser := *flagIDDeleteUser

		var InfoDeleteUser pb.UserInfo
		if flagInfoDeleteUser != nil && len(*flagInfoDeleteUser) > 0 {
			err = json.Unmarshal([]byte(*flagInfoDeleteUser), &InfoDeleteUser)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling InfoDeleteUser from %v:", flagInfoDeleteUser))
			}
		}

		var TrelloDeleteUser pb.TrelloInfo
		if flagTrelloDeleteUser != nil && len(*flagTrelloDeleteUser) > 0 {
			err = json.Unmarshal([]byte(*flagTrelloDeleteUser), &TrelloDeleteUser)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling TrelloDeleteUser from %v:", flagTrelloDeleteUser))
			}
		}

		request, err := clientHandler.DeleteUser(IDDeleteUser, InfoDeleteUser, TrelloDeleteUser)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.DeleteUser: %v\n", err)
			os.Exit(1)
		}

		v, err := service.DeleteUser(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.DeleteUser: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IDDeleteUser, InfoDeleteUser, TrelloDeleteUser)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", method)
		os.Exit(1)
	}
}
