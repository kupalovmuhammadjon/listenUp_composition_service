package service

import (
	"context"
	"database/sql"
	pb "podcast_service/genproto/episodes"
	"podcast_service/storage/postgres"
)

type EpisodeService struct {
	pb.UnimplementedEpisodesServiceServer
	EpisodeRepo *postgres.EpisodeRepo
	PodcastRepo *postgres.PodcastRepo
}

func NewEpisodeService(db *sql.DB) *EpisodeService {
	episodeRepo := postgres.NewEpisodeRepo(db)
	podcastRepo := postgres.NewPodcastRepo(db)
	return &EpisodeService{EpisodeRepo: episodeRepo, PodcastRepo: podcastRepo}
}

func (e *EpisodeService) CreatePodcastEpisode(ctx context.Context, podcast *pb.EpisodeCreate) (*pb.ID, error) {
	id, err := e.EpisodeRepo.CreatePodcastEpisode(podcast)

	return &pb.ID{Id: id}, err
}

func (e *EpisodeService) DeleteEpisode(ctx context.Context, ids *pb.IDsForDelete) (*pb.Void, error) {
	err := e.EpisodeRepo.DeletePodcastEpisode(ids)

	return &pb.Void{}, err
}

func (e *EpisodeService) PublishPodcast(ctx context.Context, id *pb.ID) (*pb.Success, error) {
	success, err := e.PodcastRepo.PublishPodcast(id)

	return success, err
}
