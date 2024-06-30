package service

import (
	"context"
	pb "podcast_service/genproto/podcasts"
	"podcast_service/storage/postgres"
)

type PodcastService struct {
	pb.UnimplementedPodcastsServer
	Podcast *postgres.PodcastRepo
}

func NewPodcastService(Podcast *postgres.PodcastRepo) *PodcastService {
	return &PodcastService{Podcast: Podcast}
}

func (p *PodcastService) GetPodcastById(ctx context.Context, req *pb.ID) (*pb.Podcast, error) {
	resp, err := p.Podcast.GetPodcastById(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) GetUserPodcasts(ctx context.Context, req *pb.ID) (*pb.UserPodcasts, error) {
	resp, err := p.Podcast.GetUserPodcasts(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
