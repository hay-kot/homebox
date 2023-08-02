package repo

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/hay-kot/homebox/backend/internal/core/services/reporting/eventbus"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/pkgs/faker"
	_ "github.com/mattn/go-sqlite3"
)

var (
	fk = faker.NewFaker()
  tbus = eventbus.New()

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
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

  go tbus.Run()

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	tClient = client
	tRepos = New(tClient, tbus, os.TempDir())
	defer client.Close()

	bootstrap()

	os.Exit(m.Run())
}
