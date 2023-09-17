package starwars

import (
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/misterek/steampipe-plugin-starwars/starwars/queries"
)

// Table Definition
func tableStarshipBuild() *plugin.Table {
	return &plugin.Table{
		Name:        "starwars_starship",
		Description: "Star Wars Starships",
		Get: &plugin.GetConfig{
			Hydrate:    listStarships,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listStarships,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: "ID"},
			{Name: "cargo_capacity", Type: proto.ColumnType_INT, Description: "The cargo capacity"},
			{Name: "cost_in_credits", Type: proto.ColumnType_INT, Description: "Cost, in credits is seems"},
			{Name: "crew", Type: proto.ColumnType_STRING, Description: "Names or something?"},
			{Name: "hyperdrive_rating", Type: proto.ColumnType_INT, Description: "The Hyperdrive rating"},
			{Name: "length", Type: proto.ColumnType_INT, Description: "The Length"},
			{Name: "max_atmosphering_speed", Type: proto.ColumnType_INT, Description: "I didn't even know atmosphering was a word"},
			{Name: "model", Type: proto.ColumnType_STRING, Description: "The model"},
			{Name: "passengers", Type: proto.ColumnType_STRING, Description: "The passengers"},
			{Name: "starship_class", Type: proto.ColumnType_STRING, Description: "The starship class"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name "},
		},
	}
}

// List Funciton
func listStarships(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Just hardcoded this
	client := graphql.NewClient("https://swapi-graphql.netlify.app/.netlify/functions/index", nil)

	query := queries.StarshipQuery{}
	err := client.Query(context.Background(), &query, nil)

	if err != nil {
		fmt.Println(err)

	}

	for _, starship := range query.AllStarships.Starships {
		d.StreamLeafListItem(ctx, starship)
	}

	return nil, nil
}

// Get function
// I probably need one of these, but 🤷‍♂️ so far.
// func getPerson(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	logger := plugin.Logger(ctx)
// 	logger.Debug("brad_thing.getBradThing")
// 	//thingOne := d.EqualsQualString("thing_one")

// 	// if thingOne == "" {
// 	// 	return nil, nil
// 	// }

// 	return Person{Id: "A1", Name: "First Name", Height: 100}, nil
// }
