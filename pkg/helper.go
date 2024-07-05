package pkg

import (
	"errors"
	"podcast_service/config"
	"podcast_service/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateUserManagementClient() (user.UserManagementClient, error) {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.USER_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	u := user.NewUserManagementClient(conn)
	return u, nil
}
