package postgres

import (
	"database/sql"
	pb "podcast_service/genproto/episodes"
	"testing"

	"github.com/google/uuid"
)

func TestValidateEpisodeId(t *testing.T) {
	id := pb.ID{Id: "fe28dcac-94a3-4a48-afcf-9e2dfb7109e4"}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	_, err = NewEpisodeRepo(db).ValidateEpisodeId(&id)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestCreatePodcastEpisode(t *testing.T) {
	newEpisode := pb.EpisodeCreate{
		PodcastId:   "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		UserId:      "65592165-c1e2-4cac-9c02-7546b34a8d27",
		Title:       "How to drink water",
		FileAudio:   []byte("Saidakbar"),
		Description: "It teaches you how to drink water",
		Duration:    15 * 60,
		Genre:       "science fiction",
		Tags:        []string{"water", "drink"},
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	got, err := NewEpisodeRepo(db).CreatePodcastEpisode(&newEpisode)
	if err != nil {
		t.Error(err)
	}
	if _, err := uuid.Parse(got); err != nil {
		t.Errorf("got %d want uuid", err)
	}
}

func TestGetEpisodesByPodcastId(t *testing.T) {
	podcastId := pb.Filter{
		Id:     "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		Limit:  10,
		Offset: 0,
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	_, err = NewEpisodeRepo(db).GetEpisodesByPodcastId(&podcastId)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestUpdateEpisode(t *testing.T) {
	ep := pb.EpisodeCreate{
		PodcastId:   "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		UserId:      "65592165-c1e2-4cac-9c02-7546b34a8d27",
		Title:       "How to drink water",
		FileAudio:   []byte("Saidakbar"),
		Description: "It teaches you how to drink water",
		Duration:    900,
		Genre:       "science fiction",
		Tags:        []string{"water", "drink"},
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	req := pb.IDs{
		PodcastId: "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		EpisodeId: "fe28dcac-94a3-4a48-afcf-9e2dfb7109e4",
		Episode:   &ep,
	}
	_, err = NewEpisodeRepo(db).UpdateEpisode(&req)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestDeletePodcastEpisode(t *testing.T) {
	ids := pb.IDsForDelete{
		PodcastId: "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		EpisodeId: "fe28dcac-94a3-4a48-afcf-9e2dfb7109e4",
	}
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	err = NewEpisodeRepo(db).DeletePodcastEpisode(&ids)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}
