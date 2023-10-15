package uam

import (
	"github.com/prongbang/uam-service/internal/shared/auth"
	"github.com/prongbang/uam-service/internal/shared/role"
	"github.com/prongbang/uam-service/internal/shared/user"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50052"

type GRPCListener interface {
	Serve()
}

type gRPCListener struct {
	AuthServer auth.AuthServer
	RoleServer role.RoleServer
	UserServer user.UserServer
}

// Serve implements GRPCListener.
func (l *gRPCListener) Serve() {
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		auth.RegisterAuthServer(s, l.AuthServer)
		role.RegisterRoleServer(s, l.RoleServer)
		user.RegisterUserServer(s, l.UserServer)

		// Register reflection service on gRPC server.
		reflection.Register(s)
		log.Printf("Server listening at %v", lis.Addr())
		if err = s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func NewListener(
	authServer auth.AuthServer,
	roleServer role.RoleServer,
	userServer user.UserServer,
) GRPCListener {
	return &gRPCListener{
		AuthServer: authServer,
		RoleServer: roleServer,
		UserServer: userServer,
	}
}
