package repo

import (
	"context"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hay-kot/git-web-template/backend/ent"
	_ "github.com/mattn/go-sqlite3"
)

var testEntClient *ent.Client
var testRepos *AllRepos

func TestMain(m *testing.M) {
	rand.Seed(int64(time.Now().Unix()))

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	testEntClient = client
	testRepos = EntAllRepos(testEntClient)

	defer client.Close()

	m.Run()

	os.Exit(m.Run())
}
