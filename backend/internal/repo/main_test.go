package repo

import (
	"context"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/pkgs/faker"
	_ "github.com/mattn/go-sqlite3"
)

var (
	fk = faker.NewFaker()

	tClient *ent.Client
	tRepos  *AllRepos
	tUser   UserOut
	tGroup  Group
)

func bootstrap() {
	var (
		err error
		ctx = context.Background()
	)

	tGroup, err = tRepos.Groups.GroupCreate(ctx, "test-group")
	if err != nil {
		log.Fatal(err)
	}

	tUser, err = tRepos.Users.Create(ctx, userFactory())
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	rand.Seed(int64(time.Now().Unix()))

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	tClient = client
	tRepos = EntAllRepos(tClient, os.TempDir())
	defer client.Close()

	bootstrap()

	os.Exit(m.Run())
}
