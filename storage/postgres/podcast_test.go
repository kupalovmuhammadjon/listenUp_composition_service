package postgres

import (
	pb "podcast_service/genproto/podcasts"
	"testing"

	"github.com/google/uuid"
)

func CreatePodcast(t *testing.T) {
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
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if _, err := uuid.Parse(*got); err != nil {
		t.Errorf("got %d want correct uuid", got)
	}
}
