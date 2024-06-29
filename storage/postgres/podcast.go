package postgres

import (
	"database/sql"
	"fmt"
	pb "podcast_service/genproto/podcasts"

	"github.com/google/uuid"
)

type PodcastRepo struct {
	Db *sql.DB
}

func NewPodcastRepo(db *sql.DB) *PodcastRepo {
	return &PodcastRepo{Db: db}
}

func (p *PodcastRepo) CreatePodcast(podcast *pb.PodcastCreate) (*string, error) {
	tx, err := p.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	newId := uuid.NewString()
	query := `insert into podcasts (id, user_id, title, description, 
	status) values ($1,$2,$3,$4,$5)`

	res, err := tx.Exec(query, newId, podcast.UserId, podcast.Title,
		podcast.Description, podcast.Status)
	if err != nil {
		return nil, err
	}
	if n, _ := res.RowsAffected(); n <= 0 {
		return nil, fmt.Errorf("cannot created")
	}
	return &newId, nil
}

func (p *PodcastRepo) DeletePodcast(podcastId *pb.ID) (*pb.Void, error) {
	tx, err := p.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	query := `update podcasts set deleted_at = $1 where deleted_at 
	is null`
	res, err := tx.Exec(query, podcastId)
	if err != nil {
		return nil, err
	}
	if n, _ := res.RowsAffected(); n <= 0 {
		return nil, fmt.Errorf("nothing is deleted")
	}
	return &pb.Void{}, nil
}
