package service

import (
	"context"
	pb "podcast_service/genproto/episodes"
	"podcast_service/storage/postgres"
)

type EpisodeService struct {
	Repo *postgres.EpisodeRepo
}

func (e *EpisodeService) UpdateEpisode(ctx context.Context, req *pb.IDs) (*pb.Void, error) {
	resp, err := e.Repo.UpdateEpisode(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
