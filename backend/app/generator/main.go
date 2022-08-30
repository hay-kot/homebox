package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/ent"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/automapper"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

// generateMappers serialized the config file into a list of automapper struct
func generateMappers() []automapper.AutoMapper {
	return []automapper.AutoMapper{
		{
			Package: "mapper",
			Prefix:  "users",
			Name:    "User Out",
			Schema: automapper.Schema{
				Type:   types.UserOut{},
				Prefix: "types",
			},
			Model: automapper.Model{
				Type:   ent.User{},
				Prefix: "ent",
			},
			Imports: []string{},
		},
	}
}

func generateTypeScript() {
	// Configuration
	converter := typescriptify.New()
	converter.CreateInterface = true
	converter.ManageType(uuid.UUID{}, typescriptify.TypeOptions{TSType: "string"})
	converter.ManageType(time.Time{}, typescriptify.TypeOptions{TSType: "Date", TSTransform: "new Date(__VALUE__)"})

	// General
	public := []any{
		// Base Types
		types.ApiSummary{},

		// User Types
		types.UserOut{},
		types.UserCreate{},
		types.UserIn{},
		types.UserUpdate{},

		// Auth Types
		types.LoginForm{},
		types.TokenResponse{},
	}

	for i := 0; i < len(public); i++ {
		converter.Add(public[i])
	}

	// Creation
	converter.ConvertToFile("./generated-types.ts")

}

func main() {
	automappers := generateMappers()
	conf := automapper.DefaultConf()

	automapper.Generate(automappers, conf)

	generateTypeScript()
}
