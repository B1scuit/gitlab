package gitlab_test

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/B1scuit/gitlab"
	"github.com/B1scuit/gitlab/types"
)

var client *gitlab.Client

func TestMain(m *testing.M) {

	log := log.New(os.Stdout, "", 0)

	client = gitlab.NewClient(&types.ClientOptions{
		HttpClient: &http.Client{
			Timeout: time.Second * 3,
		},
		Log:    log,
		ErrLog: log,
	})

	m.Run()
}

func TestProjectsList(t *testing.T) {
	projects, err := client.Project.List(
		client.Project.List.SetContext(context.TODO()),
		client.Project.List.SetOwned(true),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(projects)
}
