package queries

import (
	"github.com/misterek/steampipe-plugin-starwars/starwars/models"
	"github.com/shurcooL/graphql"
)

// GraphQL Queries

type PersonQuery struct {
	AllPeople struct {
		TotalCount graphql.Int
		People     []models.Person
	}
}
