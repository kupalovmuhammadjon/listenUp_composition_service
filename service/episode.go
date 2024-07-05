package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pb "podcast_service/genproto/episodes"
	pbUser "podcast_service/genproto/user"
	"podcast_service/pkg"
	"podcast_service/storage/postgres"
)

type EpisodeService struct {
	pb.UnimplementedEpisodesServiceServer
	Episode    *postgres.EpisodeRepo
	Podcast    *postgres.PodcastRepo
	UserClient pbUser.UserManagementClient
}

func NewEpisodeService(db *sql.DB) *EpisodeService {
	episodeRepo := postgres.NewEpisodeRepo(db)
	podcastRepo := postgres.NewPodcastRepo(db)

	user := CreateClientsForEpisode()
	return &EpisodeService{
		Episode:    episodeRepo,
		Podcast:    podcastRepo,
		UserClient: *user,
	}
}

func CreateClientsForEpisode() *pbUser.UserManagementClient {
	usermanagement, err := pkg.CreateUserManagementClient()
	if err != nil {
		log.Println(err)
	}
	return &usermanagement
}

func (e *EpisodeService) CreatePodcastEpisode(ctx context.Context, podcast *pb.EpisodeCreate) (*pb.ID, error) {
	// check podcastId id valid
	valid, err := e.Podcast.ValidatePodcastId(podcast.PodcastId)
	if err != nil || !valid.Success {
		return nil, fmt.Errorf("podcast Id is not valid %s", err)
	}

	// check user id valid
	valid1, err := e.UserClient.ValidateUserId(
		ctx, &pbUser.ID{Id: podcast.UserId})
	if err != nil || !valid1.Success {
		return nil, fmt.Errorf("user Id is not valid %s", err)
	}

	id, err := e.Episode.CreatePodcastEpisode(podcast)
	return &pb.ID{Id: id}, err
}

func (e *EpisodeService) GetEpisodesByPodcastId(ctx context.Context, req *pb.Filter) (*pb.Episodes, error) {
	resp, err := e.Episode.GetEpisodesByPodcastId(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EpisodeService) UpdateEpisode(ctx context.Context, req *pb.IDs) (*pb.Void, error) {
	// check podcastId id valid
	valid, err := e.Podcast.ValidatePodcastId(req.Episode.PodcastId)
	if err != nil || !valid.Success {
		return nil, fmt.Errorf("podcast Id is not valid %s", err)
	}

	// check user id valid
	valid1, err := e.UserClient.ValidateUserId(
		ctx, &pbUser.ID{Id: req.Episode.UserId})
	if err != nil || !valid1.Success {
		return nil, fmt.Errorf("user Id is not valid %s", err)
	}

	resp, err := e.Episode.UpdateEpisode(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EpisodeService) DeleteEpisode(ctx context.Context, ids *pb.IDsForDelete) (*pb.Void, error) {
	err := e.Episode.DeletePodcastEpisode(ids)

	return &pb.Void{}, err
}

func (e *EpisodeService) SearchEpisodeByTitle(ctx context.Context, title *pb.Title) (*pb.Episode, error) {
	resp, err := e.Episode.SearchEpisodeByTitle(title.Title)
	return resp, err
}

func (e *EpisodeService) ValidateEpisodeId(ctx context.Context, req *pb.ID) (*pb.Success, error) {
	resp, err := e.Episode.ValidateEpisodeId(req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
