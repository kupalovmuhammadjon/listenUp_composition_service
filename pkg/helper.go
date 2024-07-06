package pkg

import (
	"errors"
	"log"
	"podcast_service/config"
	pbCollaboration "podcast_service/genproto/collaborations"
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

	u := user.NewUserManagementClient(conn)
	return u, nil
}

func NewCollaborationClient() pbCollaboration.CollaborationsClient {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.COLLABORATIONS_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting collaborations service ", err)
	}
	a := pbCollaboration.NewCollaborationsClient(conn)

	return a
}
