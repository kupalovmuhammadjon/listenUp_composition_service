package postgres

import (
	"context"
	"database/sql"
	pb "podcast_service/genproto/podcasts"
)

type PodcastRepo struct {
	Db *sql.DB
}

func (p *PodcastRepo) GetPodcastById(ctx context.Context, in *pb.ID) (*pb.Podcast, error) {
	podcast := &pb.Podcast{Id: in.Id}

	query := `select user_id, title, description, created_at, updated_at
	from podcasts where id = $1 and deleted_at = null`

	err := p.Db.QueryRow(query, in.Id).Scan(&podcast.UserId, &podcast.Title, &podcast.Description, &podcast.CreatedAt, &podcast.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return podcast, nil
}

func (p *PodcastRepo) GetUserPodcasts(ctx context.Context, in *pb.ID) (*pb.UserPodcasts, error) {
	query := `select id, user_id, title, description, created_at, updated_at
	from podcasts where user_id = $1 and deleted_at = null`

	rows, err := p.Db.Query(query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []*pb.Podcast
	for rows.Next() {
		var podcast pb.Podcast
		err := rows.Scan(&podcast.Id, &podcast.UserId, &podcast.Title, &podcast.Description, &podcast.CreatedAt, &podcast.UpdatedAt)
		if err != nil {
			return nil, err
		}
		podcasts = append(podcasts, &podcast)
	}

	return &pb.UserPodcasts{Podcasts: podcasts}, nil
}
