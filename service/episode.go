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
}

func NewEpisodeService(db *sql.DB) *EpisodeService {
	episodeRepo := postgres.NewEpisodeRepo(db)
	return &EpisodeService{EpisodeRepo: episodeRepo}
}

func (e *EpisodeService) CreatePodcastEpisode(ctx context.Context, podcast *pb.EpisodeCreate) (*pb.ID, error) {
	id, err := e.EpisodeRepo.CreatePodcastEpisode(podcast)

	return &pb.ID{Id: id}, err
}

func (e *EpisodeService) DeleteEpisode(ctx context.Context, ids *pb.IDsForDelete) (*pb.Void, error) {
	err := e.EpisodeRepo.DeletePodcastEpisode(ids)

	return &pb.Void{}, err
}
