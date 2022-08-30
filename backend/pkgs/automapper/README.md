# Automapper


Automapper is an opinionated Go library that provides a dead simple interface to mapping 1-1 models To/From a database Model to a DTO or Schema using value semantics. It does not rely on code comments, but instead uses standard Go code to define your mapping and configuration to make it easy to use an refactor. 

Current Limitation
- flat/single level models
- single schema to model per config entry
- limited configuration (support lowercase, camelcase, snakecase, etc)


Future Considerations
- [ ] Recursive mapping of embed structs
- [ ] Optional generate time type checker.
- [ ] Ensure values are copied to the destination and not just a reference
- [ ] ?!?!?


## Example Configuration

```go
package main

import (
	"github.com/mealie-recipes/mealie-analytics/ent"
	"github.com/mealie-recipes/mealie-analytics/internal/types"
	"github.com/mealie-recipes/mealie-analytics/pkgs/automapper"
)

// getMappers serialized the config file into a list of automapper struct
func getMappers() []automapper.AutoMapper {
	return []automapper.AutoMapper{
		{
			Package: "mapper", // generated package name   
			Prefix:  "analytics", // generating file prefix -> analytics_automapper.go
			Name:    "Mealie Analytics", // For console output
			Schema: automapper.Schema{
				Type:   types.Analytics{}, 
				Prefix: "types", // Package namespace
			},
			Model: automapper.Model{
				Type:   ent.Analytics{},
				Prefix: "ent", // Package namespace
			},
			Imports: []string{}, // Specify additional imports here
		},
	}
}

func main() {
	automappers := getMappers()
	conf := automapper.DefaultConf()

	automapper.Generate(automappers, conf)
}
```