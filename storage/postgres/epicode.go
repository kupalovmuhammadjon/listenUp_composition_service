package postgres

import (
	"database/sql"
	"fmt"
	pb "podcast_service/genproto/episodes"
)

type EpisodeRepo struct {
	Db *sql.DB
}

func NewEpisodeRepo(db *sql.DB) *EpisodeRepo {
	return &EpisodeRepo{Db: db}
}

func (p *PodcastRepo) GetEpisodesByPodcastId(podcastId *pb.ID) (*pb.Episodes, error) {
	query := `select id, podcast_id, user_id, title, file_audio,
	description, duration, created_at, updated_at from episodes 
	where podcast_id = $1`

	rows, err := p.Db.Query(query, podcastId)
	if err != nil {
		return nil, err
	}

	episodes := pb.Episodes{}
	for rows.Next() {
		episode := pb.Episode{}
		err := rows.Scan(&episode.Id, &episode.PodcastId, &episode.UserId,
			&episode.Title, &episode.FileAudio, &episode.Description,
			&episode.Duration, &episode.CreatedAt, &episode.UpdatedAt)
		if err != nil {
			return nil, err
		}
		episodes.Episodes = append(episodes.Episodes, &episode)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &episodes, nil
}

func (p *PodcastRepo) PublishPodcast(podcastId *pb.ID) (*pb.Success, error) {
	tx, err := p.Db.Begin()
	if err != nil {
		return &pb.Success{Success: false}, err
	}
	defer tx.Commit()

	query := `update podcasts set status = 'published' where status != 
	'published' and id = $1`

	res, err := tx.Exec(query, podcastId)
	if err != nil {
		return &pb.Success{Success: false}, err
	}

	if err, _ := res.RowsAffected(); err <= 0 {
		return &pb.Success{Success: false}, fmt.Errorf("no rows affected")
	}
	return &pb.Success{Success: true}, nil
}
