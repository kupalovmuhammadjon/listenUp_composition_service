package postgres

import (
	"database/sql"
	pb "podcast_service/genproto/episodes"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type EpisodeRepo struct {
	Db *sql.DB
}

func NewEpisodeRepo(db *sql.DB) *EpisodeRepo {
	return &EpisodeRepo{Db: db}
}

func (e *EpisodeRepo) CreatePodcastEpisode(episode *pb.EpisodeCreate) (string, error) {
	query := `
	insert into episodes(
		id, podcast_id, user_id, title, file_audio, description,
      	duration, genre, tags,updated_at
    ) values (
	 	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
	)`

	id := uuid.NewString()
	tx, err := e.Db.Begin()
	if err != nil {
		return "", err
	}
	_, err = tx.Exec(query, id, episode.PodcastId, episode.UserId, episode.Title,
		episode.FileAudio, episode.Description, episode.Duration, episode.Genre,
		pq.Array(episode.Tags), time.Now())
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

func (p *EpisodeRepo) GetEpisodesByPodcastId(filter *pb.Filter) (*pb.Episodes, error) {
	query := `
	select 
		id, podcast_id, user_id, title, file_audio,
		description, duration, genre, tags, created_at, updated_at 
	from 
		episodes 
	where
		deleted_at is null
		and podcast_id = $1
	limit $2
	offset $3`

	rows, err := p.Db.Query(query, filter.Id, filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}

	episodes := pb.Episodes{}
	for rows.Next() {
		episode := pb.Episode{}
		err := rows.Scan(&episode.Id, &episode.PodcastId, &episode.UserId,
			&episode.Title, &episode.FileAudio, &episode.Description,
			&episode.Duration, &episode.Genre, pq.Array(&episode.Tags),
			&episode.CreatedAt, &episode.UpdatedAt)
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

func (e *EpisodeRepo) UpdateEpisode(podcastIds *pb.IDs) (*pb.Void, error) {
	query := `
		update
			episodes 
		set 
			user_id = $1, title = $2, file_audio = $3,
			description = $4, duration = $5, genre = $6,
			tags = $7, updated_at = $8
		where
			podcast_id = $9 and 
			id = $10 and 
			deleted_at IS NULL`

	tr, err := e.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	_, err = tr.Exec(query, podcastIds.Episode.UserId,
		podcastIds.Episode.Title, podcastIds.Episode.FileAudio,
		podcastIds.Episode.Description, podcastIds.Episode.Duration,
		podcastIds.Episode.Genre, pq.Array(podcastIds.Episode.Tags),
		time.Now(), podcastIds.PodcastId, podcastIds.EpisodeId)

	return &pb.Void{}, err
}

func (e *EpisodeRepo) DeletePodcastEpisode(ids *pb.IDsForDelete) error {
	query := `
	update
	    episodes 
	set 
	    deleted_at = now() 
	where
		deleted_at is null
	    and id = $1 and
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

func (e *EpisodeRepo) SearchEpisodeByTitle(title string) (*pb.Episode, error) {
	query := `
	select
		id, podcast_id, user_id, file_audio, description,
		duration, genre, tags, created_at, updated_at
	from
		episodes
	where
		deleted_at = 0
		and title = $1`

	ep := pb.Episode{Title: title}

	row := e.Db.QueryRow(query, title)
	err := row.Scan(&ep.Id, &ep.PodcastId, &ep.UserId, &ep.FileAudio, &ep.Description,
		&ep.Duration, &ep.Genre, &ep.Tags, &ep.CreatedAt, &ep.UpdatedAt)

	return &ep, err
}

func (e *EpisodeRepo) ValidateEpisodeId(id *pb.ID) (*pb.Success, error) {
	query := `
		select
			case 
				when id = $1 then true
			else
				false
			end
		from
			episodes
	`
	res := pb.Success{}
	err := e.Db.QueryRow(query, id.Id).Scan(&res.Success)

	return &res, err
}
