package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	pb "podcast_service/genproto/episodes"
	"time"
)

type EpisodeRepo struct {
	Db *sql.DB
}

func NewEpisodeRepo(db *sql.DB) *EpisodeRepo {
	return &EpisodeRepo{Db: db}
}

func (e *EpisodeRepo) CreatePodcastEpisode(episode *pb.EpisodeCreate) (string, error) {
	query := `
	insert into
		episodes(
		episode_id,
		podcast_id,
   	 	user_id,
      	title,
      	file_audio,
      	description,
      	duration,
		updated_at
      	)
	values ($1, $2, $3, $4, $5, $6. $7, $8)	
`
	id := uuid.NewString()
	tx, err := e.Db.Begin()
	if err != nil {
		return "", err
	}
	_, err = tx.Exec(query, id, episode.PodcastId, episode.UserId, episode.Title, episode.FileAudio,
		episode.Description, episode.Duration, time.Now())
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (e *EpisodeRepo) DeletePodcastEpisode(ids pb.IDsForDelete) error {
	query := `
	update 
	    episodes 
	set 
	    deleted_at = now() 
	where 
	    episode_id = $1,
		podcast_id = $2
`
	tx, err := e.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, ids.EpisodeId, ids.PodcastId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
