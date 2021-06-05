package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"

	"github.com/tora0091/get-account-data/config"
)

func main() {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, config.GetDatabasePath())
	if err != nil {
		panic(err)
	}
	defer client.Close()

	columns := []string{"property_id", "plan_id", "category_id", "price", "name", "start_date", "end_date", "section"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate(config.GetTableName(), columns, []interface{}{10001, 9001, 1, 670000, "New York Hot!", "2021-07-01", "2021-07-10", 7}),
		spanner.InsertOrUpdate(config.GetTableName(), columns, []interface{}{34001, 67, 5, 1200, "Pepper music", "2021-07-20", "2021-07-30", 1}),
		spanner.InsertOrUpdate(config.GetTableName(), columns, []interface{}{55601, 44, 6, 9000, "Santa Cluz", "2021-06-01", "2021-07-31", 4}),
	}
	_, err = client.Apply(ctx, m)
	if err != nil {
		panic(err)
	}
	fmt.Println("> Insert OK.")
}
