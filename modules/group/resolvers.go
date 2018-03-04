package group

import (
	"github.com/aprimediet/gographql/db"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// List is graphql query to list all groups
var List = &graphql.Field{
	Type:        graphql.NewList(GroupType),
	Description: "Group List",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		groups := []Group{}

		err := db.DB.C("groups").Find(bson.M{}).All(&groups)

		return groups, err
	},
}

// Retrieve is graphql query to get single user
var Retrieve = &graphql.Field{
	Type:        GroupType,
	Description: "Single Group",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(string)
		group := Group{}

		err := db.DB.C("groups").FindId(bson.ObjectIdHex(id)).One(&group)

		return group, err
	},
}

// Create is graphql mutation to create user
var Create = &graphql.Field{
	Type:        GroupType,
	Description: "Create New Group",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)

		newGroup := Group{
			ID:          bson.NewObjectId(),
			Name:        name,
			Description: description,
		}

		err := db.DB.C("groups").Insert(newGroup)

		return newGroup, err
	},
}

// Update is graphql mutation to update user
var Update = &graphql.Field{
	Type:        GroupType,
	Description: "Update Group",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(string)
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)
		affectedGroup := Group{}

		err := db.DB.C("groups").FindId(bson.ObjectIdHex(id)).One(&affectedGroup)

		if name != "" {
			affectedGroup.Name = name
		}

		if description != "" {
			affectedGroup.Description = description
		}

		err = db.DB.C("groups").UpdateId(bson.ObjectIdHex(id), affectedGroup)

		return affectedGroup, err
	},
}

// Delete is graphql mutation to remove user
var Delete = &graphql.Field{
	Type:        GroupType,
	Description: "Remove Group",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(string)
		affectedGroup := Group{}

		err := db.DB.C("groups").FindId(bson.ObjectIdHex(id)).One(&affectedGroup)

		err = db.DB.C("groups").RemoveId(bson.ObjectIdHex(id))

		return affectedGroup, err
	},
}
