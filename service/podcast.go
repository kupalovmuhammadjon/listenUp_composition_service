package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pbCol "podcast_service/genproto/collaborations"
	pb "podcast_service/genproto/podcasts"
	pbUser "podcast_service/genproto/user"

	"podcast_service/pkg"
	"podcast_service/storage/postgres"
)

type PodcastService struct {
	pb.UnimplementedPodcastsServer
	Podcast              *postgres.PodcastRepo
	UserClient           pbUser.UserManagementClient
	CollaborationsClient pbCol.CollaborationsClient
}

func NewPodcastService(db *sql.DB) *PodcastService {
	podcast := postgres.NewPodcastRepo(db)
	user, collaborations := CreateClientsForPodcast()
	return &PodcastService{
		Podcast:              podcast,
		UserClient:           *user,
		CollaborationsClient: *collaborations,
	}
}

func CreateClientsForPodcast() (*pbUser.UserManagementClient,
	*pbCol.CollaborationsClient) {
	usermanagement, err := pkg.CreateUserManagementClient()
	collaborations := pkg.NewCollaborationClient()
	if err != nil {
		log.Println(err)
	}
	return &usermanagement, &collaborations
}

func (p *PodcastService) CreatePodcast(ctx context.Context, req *pb.PodcastCreate) (*pb.ID, error) {
	// check user id valid
	valid1, err := p.UserClient.ValidateUserId(
		ctx, &pbUser.ID{Id: req.UserId})
	if err != nil || !valid1.Success {
		return nil, fmt.Errorf("user Id is not valid %s", err)
	}

	resp, err := p.Podcast.CreatePodcast(req)
	if err != nil {
		return nil, err
	}

	_, err = p.CollaborationsClient.CreateOwner(ctx, &pbCol.CreateAsOwner{
		PodcastId: *resp,
		UserId:    req.UserId,
		Role:      "owner",
	})
	if err != nil {
		return nil, err
	}
	return &pb.ID{Id: *resp}, nil
}

func (p *PodcastService) GetPodcastById(ctx context.Context, req *pb.ID) (*pb.Podcast, error) {
	resp, err := p.Podcast.GetPodcastById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) UpdatePodcast(ctx context.Context, podcast *pb.PodcastUpdate) (*pb.Void, error) {
	// check user id valid
	valid1, err := p.UserClient.ValidateUserId(
		ctx, &pbUser.ID{Id: podcast.UserId})
	if err != nil || !valid1.Success {
		return nil, fmt.Errorf("user Id is not valid %s", err)
	}

	err = p.Podcast.UpdatePodcast(podcast)
	return &pb.Void{}, err
}

func (p *PodcastService) DeletePodcast(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	resp, err := p.Podcast.DeletePodcast(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) GetUserPodcasts(ctx context.Context, req *pb.Filter) (*pb.UserPodcasts, error) {
	// check user id valid
	valid1, err := p.UserClient.ValidateUserId(
		ctx, &pbUser.ID{Id: req.Id})
	if err != nil || !valid1.Success {
		return nil, fmt.Errorf("user Id is not valid %s", err)
	}

	resp, err := p.Podcast.GetUserPodcasts(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) PublishPodcast(ctx context.Context, id *pb.ID) (*pb.Success, error) {
	success, err := p.Podcast.PublishPodcast(id)

	return success, err
}

func (e *PodcastService) ValidatePodcastId(ctx context.Context, req *pb.ID) (*pb.Success, error) {
	resp, err := e.Podcast.ValidatePodcastId(req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
