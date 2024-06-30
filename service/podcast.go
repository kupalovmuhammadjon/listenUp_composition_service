package service

import (
	"context"
	"database/sql"
	pb "podcast_service/genproto/podcasts"
	"podcast_service/storage/postgres"
)

type PodcastService struct {
	pb.UnimplementedPodcastsServer
	PodcastRepo *postgres.PodcastRepo
}

func NewPodcastService(db *sql.DB) *PodcastService {
	podcastRepo := postgres.NewPodcastRepo(db)
	return &PodcastService{PodcastRepo: podcastRepo}
}

func (p *PodcastService) UpdatePodcast(ctx context.Context, podcast *pb.PodcastUpdate) (*pb.Void, error) {
	err := p.PodcastRepo.UpdatePodcast(podcast)

	return &pb.Void{}, err
}
