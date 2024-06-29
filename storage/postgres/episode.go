package postgres

import (
	"context"
	"database/sql"
	"fmt"
	pb "podcast_service/genproto/episodes"
	"time"
)

type EpisodeRepo struct {
	Db *sql.DB
}

func NewEpisodeRepo(db *sql.DB) *EpisodeRepo {
	return &EpisodeRepo{Db: db}
}

func (e *EpisodeRepo) UpdateEpisode(ctx context.Context, in *pb.IDs) (*pb.Void, error) {
	query := `update episodes set `
	params := []interface{}{}

	if in.Episode.UserId != "" {
		query += fmt.Sprintf("user_id = $%d, ", len(params)+1)
		params = append(params, in.Episode.UserId)
	}
	if in.Episode.Title != "" {
		query += fmt.Sprintf("title = $%d, ", len(params)+1)
		params = append(params, in.Episode.Title)
	}
	if in.Episode.FileAudio != nil {
		query += fmt.Sprintf("file_audio = $%d, ", len(params)+1)
		params = append(params, in.Episode.FileAudio)
	}
	if in.Episode.Description != "" {
		query += fmt.Sprintf("description = $%d, ", len(params)+1)
		params = append(params, in.Episode.Description)
	}
	if in.Episode.Duration > 0 {
		query += fmt.Sprintf("duration = $%d, ", len(params)+1)
		params = append(params, in.Episode.Duration)
	}
	query += fmt.Sprintf("updated_at = $%d ", len(params)+1)
	params = append(params, time.Now())
	query += fmt.Sprintf("where podcast_id = $%d ", len(params)+1)
	params = append(params, in.PodcastId)
	query += fmt.Sprintf(" and id = $%d and deleted_at = null", len(params)+1)
	params = append(params, in.EpisodeId)

	tr, err := e.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = e.Db.Exec(query, params...)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
