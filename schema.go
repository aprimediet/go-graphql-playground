package main

import "github.com/graphql-go/graphql"

// Schema is Schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})
