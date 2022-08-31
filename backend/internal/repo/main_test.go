package repo

import (
	"context"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hay-kot/content/backend/ent"
	_ "github.com/mattn/go-sqlite3"
)

var (
	testEntClient *ent.Client
	testRepos     *AllRepos
	testUser      *ent.User
	testGroup     *ent.Group
)

func bootstrap() {
	ctx := context.Background()
	testGroup, _ = testRepos.Groups.Create(ctx, "test-group")
	testUser, _ = testRepos.Users.Create(ctx, UserFactory())

	if testGroup == nil || testUser == nil {
		log.Fatal("Failed to bootstrap test data")
	}

}

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

	bootstrap()

	os.Exit(m.Run())
}
