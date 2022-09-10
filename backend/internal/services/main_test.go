package services

import (
	"context"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/faker"
	_ "github.com/mattn/go-sqlite3"
)

var (
	fk = faker.NewFaker()

	tClient *ent.Client
	tRepos  *repo.AllRepos
	tUser   *ent.User
	tGroup  *ent.Group
)

func bootstrap() {
	var (
		err error
		ctx = context.Background()
	)

	tGroup, err = tRepos.Groups.Create(ctx, "test-group")
	if err != nil {
		log.Fatal(err)
	}

	tUser, err = tRepos.Users.Create(ctx, types.UserCreate{
		Name:        fk.Str(10),
		Email:       fk.Email(),
		Password:    fk.Str(10),
		IsSuperuser: fk.Bool(),
		GroupID:     tGroup.ID,
	})
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
	tRepos = repo.EntAllRepos(tClient)
	defer client.Close()

	bootstrap()

	os.Exit(m.Run())
}
