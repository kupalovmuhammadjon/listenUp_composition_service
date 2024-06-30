package service

import (
	"context"
	pb "podcast_service/genproto/episodes"
	"podcast_service/storage/postgres"
)

type EpisodeService struct {
	pb.UnimplementedEpisodesServiceServer
	Episode *postgres.EpisodeRepo
}

func NewEpisodeService(Episode *postgres.EpisodeRepo) *EpisodeService {
	return &EpisodeService{Episode: Episode}
}

func (e *EpisodeService) UpdateEpisode(ctx context.Context, req *pb.IDs) (*pb.Void, error) {
	resp, err := e.Episode.UpdateEpisode(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
