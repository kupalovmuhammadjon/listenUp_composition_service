package postgres

import (
	"database/sql"
	"fmt"
	pb "podcast_service/genproto/podcasts"
	"time"
)

type PodcastRepo struct {
	Db *sql.DB
}

func NewPodcastRepo(db *sql.DB) *PodcastRepo {
	return &PodcastRepo{Db: db}
}

func (p *PodcastRepo) UpdatePodcast(poadcast *pb.PodcastUpdate) error {
	query := `
	update 
		podcasts
	set
	    user_id  = $1,
    	title = $2,
    	description = $3,
   	 	status = $4,
    	updated_at = $5
	where
	    id = $6
`
	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec(query, poadcast.UserId, poadcast.Title, poadcast.Description, poadcast.Status,
		time.Now(), poadcast.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	ra, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if ra == 0 {
		return fmt.Errorf("no rows affected")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
