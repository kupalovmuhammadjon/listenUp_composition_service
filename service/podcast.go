package service

import (
	"context"
	pb "podcast_service/genproto/podcasts"
	"podcast_service/storage/postgres"
)

type PodcastService struct {
	Repo *postgres.PodcastRepo
}

func (p *PodcastService) CreatePodcast(req *pb.PodcastCreate) (*string, error) {
	resp, err := p.Repo.CreatePodcast(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) DeletePodcast(req *pb.ID) (*pb.Void, error) {
	resp, err := p.Repo.DeletePodcast(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) GetPodcastById(ctx context.Context, req *pb.ID) (*pb.Podcast, error) {
	resp, err := p.Repo.GetPodcastById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) GetUserPodcasts(ctx context.Context, req *pb.ID) (*pb.UserPodcasts, error) {
	resp, err := p.Repo.GetUserPodcasts(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
