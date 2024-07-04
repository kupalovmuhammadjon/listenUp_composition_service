package postgres

import (
	"database/sql"
	pb "podcast_service/genproto/episodes"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

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
	podcastId := pb.ID{Id: "8e89c32d-1425-4f6f-b86a-ab85c4af870c"}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	got, err := NewEpisodeRepo(db).GetEpisodesByPodcastId(&podcastId)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
	ep := pb.Episode{
		Id:          "fe28dcac-94a3-4a48-afcf-9e2dfb7109e4",
		PodcastId:   "8e89c32d-1425-4f6f-b86a-ab85c4af870c",
		UserId:      "65592165-c1e2-4cac-9c02-7546b34a8d27",
		Title:       "How to drink water",
		FileAudio:   []byte("Saidakbar"),
		Description: "It teaches you how to drink water",
		Duration:    900,
		Genre:       "science fiction",
		Tags:        []string{"water", "drink"},
		CreatedAt:   "2024-07-04T10:57:30.721807+05:00",
		UpdatedAt:   "2024-07-04T10:57:30.721895Z",
	}
	want := pb.Episodes{
		Episodes: []*pb.Episode{&ep},
	}
	if !reflect.DeepEqual(got, &want) {
		t.Errorf("\nGot %s\nwant %s", got, &want)
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
