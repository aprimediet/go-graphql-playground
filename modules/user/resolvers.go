package user

import (
	"log"

	"github.com/aprimediet/gographql/db"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// List is graphql query to list all users
var List = &graphql.Field{
	Type:        graphql.NewList(UserType),
	Description: "User List",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		users := []User{}
		context := params.Context

		log.Print(context)

		err := db.DB.C("users").Find(bson.M{}).All(&users)

		return users, err
	},
}

// Retrieve is graphql query to get single user
var Retrieve = &graphql.Field{
	Type:        UserType,
	Description: "Single User",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(string)
		user := User{}

		err := db.DB.C("users").FindId(bson.ObjectIdHex(id)).One(&user)

		return user, err
	},
}

// Create is graphql mutation to create user
var Create = &graphql.Field{
	Type:        UserType,
	Description: "Create New User",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name, _ := params.Args["name"].(string)
		email, _ := params.Args["email"].(string)

		newUser := User{
			ID:    bson.NewObjectId(),
			Name:  name,
			Email: email,
		}

		err := db.DB.C("users").Insert(newUser)

		return newUser, err
	},
}

// Update is graphql mutation to update user
var Update = &graphql.Field{
	Type:        UserType,
	Description: "Update User",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"groups": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(string)
		name, _ := params.Args["name"].(string)
		email, _ := params.Args["email"].(string)
		groups, _ := params.Args["groups"].([]interface{})
		affectedUser := User{}
		newGroup := []bson.ObjectId{}

		err := db.DB.C("users").FindId(bson.ObjectIdHex(id)).One(&affectedUser)

		if name != "" {
			affectedUser.Name = name
		}

		if email != "" {
			affectedUser.Email = email
		}

		if len(groups) != 0 {
			for _, v := range groups {
				// sGroup := group.Group{}
				// sGroup.ID = bson.ObjectIdHex(v.(string))
				newGroup = append(newGroup, bson.ObjectIdHex(v.(string)))
			}
			affectedUser.Groups = newGroup
		}

		err = db.DB.C("users").UpdateId(bson.ObjectIdHex(id), affectedUser)

		return affectedUser, err
	},
}

// Delete is graphql mutation to remove user
var Delete = &graphql.Field{
	Type:        UserType,
	Description: "Remove user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(string)
		affectedUser := User{}

		err := db.DB.C("users").FindId(bson.ObjectIdHex(id)).One(&affectedUser)

		err = db.DB.C("users").RemoveId(bson.ObjectIdHex(id))

		return affectedUser, err
	},
}
