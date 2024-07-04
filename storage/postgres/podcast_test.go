package postgres

import (
	"database/sql"
	pb "podcast_service/genproto/podcasts"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestCreatePodcast(t *testing.T) {
	newPodcast := pb.PodcastCreate{
		UserId:      "65592165-c1e2-4cac-9c02-7546b34a8d27",
		Title:       "BBS learning english",
		Description: "learning english",
		Status:      "completed",
	}
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("error: %s", err)
	}
	podcastsRepo := NewPodcastRepo(db)
	got, err := podcastsRepo.CreatePodcast(&newPodcast)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
		return
	}
	if _, err := uuid.Parse(*got); err != nil {
		t.Errorf("got %d want correct uuid", got)
	}
}

func TestGetPodcastById(t *testing.T) {
	id := pb.ID{Id: "8e89c32d-1425-4f6f-b86a-ab85c4af870c"}
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("error: %s", err)
	}
	want := pb.Podcast{
		Id:          "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		UserId:      "65592165-c1e2-4cac-9c02-7546b34a8d27",
		Title:       "BBS learning english",
		Description: "learning english",
		CreatedAt:   "2024-07-03T19:27:51.479733+05:00",
	}
	got, err := NewPodcastRepo(db).GetPodcastById(&id)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %s\nwant %s", got, &want)
	}
}

func TestGetUserPodcasts(t *testing.T) {
	id := pb.ID{Id: "65592165-c1e2-4cac-9c02-7546b34a8d27"}
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("error: %s", err)
	}
	pod := pb.Podcast{
		Id:          "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		UserId:      "65592165-c1e2-4cac-9c02-7546b34a8d27",
		Title:       "BBS learning english",
		Description: "learning english",
		CreatedAt:   "2024-07-03T19:27:51.479733+05:00",
	}
	want := pb.UserPodcasts{
		Podcasts: []*pb.Podcast{&pod},
	}
	got, err := NewPodcastRepo(db).GetUserPodcasts(&id)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %s\nwant %s", got, &want)
	}
}

func TestUpdatePodcast(t *testing.T) {
	podcast := pb.PodcastUpdate{
		Id:          "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		UserId:      "65592165-c1e2-4cac-9c02-7546b34a8d27",
		Title:       "BBS learning english",
		Description: "learning english",
		Status:      "in_progress",
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	got := NewPodcastRepo(db).UpdatePodcast(&podcast)
	if got != nil || got == sql.ErrNoRows {
		t.Errorf("got %s want nil", got)
	}
}

func TestDeletePodcast(t *testing.T) {
	podcastId := pb.ID{Id: "8e89c32d-1425-4f6f-b86a-ab85c4af870c"}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	_, got := NewPodcastRepo(db).DeletePodcast(&podcastId)
	if got != nil || got == sql.ErrNoRows {
		t.Errorf("got %s want nil", got)
	}
}
