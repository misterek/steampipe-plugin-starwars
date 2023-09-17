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
func tablePeopleBuild() *plugin.Table {
	return &plugin.Table{
		Name:        "people",
		Description: "Star Wars People",
		Get: &plugin.GetConfig{
			Hydrate:    listPerson,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPerson,
		},

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: "ID"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name"},
			{Name: "height", Type: proto.ColumnType_INT, Description: "The Height in some unit"},
			{Name: "birth_year", Type: proto.ColumnType_STRING, Description: "The Birth Year"},
			{Name: "eye_color", Type: proto.ColumnType_STRING, Description: "The eye color"},
			{Name: "gender", Type: proto.ColumnType_STRING, Description: "The gender"},
			{Name: "hair_color", Type: proto.ColumnType_STRING, Description: "Hair Color"},
			{Name: "skin_color", Type: proto.ColumnType_STRING, Description: "Skin Color"},
		},
	}
}

// List Funciton
func listPerson(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Just hardcoded this
	client := graphql.NewClient("https://swapi-graphql.netlify.app/.netlify/functions/index", nil)

	query := queries.PersonQuery{}
	err := client.Query(context.Background(), &query, nil)

	if err != nil {
		fmt.Println(err)
	}

	for _, person := range query.AllPeople.People {
		d.StreamLeafListItem(ctx, person)
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
