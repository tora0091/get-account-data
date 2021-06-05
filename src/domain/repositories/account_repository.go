package repositories

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/tora0091/get-account-data/config"
	"github.com/tora0091/get-account-data/domain"
	"github.com/tora0091/get-account-data/domain/schema"
	"google.golang.org/api/iterator"
)

type AccountRepository struct{}

func NewAccountRepository(db *domain.SqlHandler) *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) FindAll(limit int64) (schema.Accounts, error) {
	sql := `select
		property_id,
		plan_id,
		category_id,
		price,
		name,
		format_date("%Y-%m-%d", start_date) start_date,
		format_date("%Y-%m-%d", end_date) end_date,
		section from ` + config.GetTableName() + ` limit @limit`

	params := map[string]interface{}{
		"limit": limit,
	}
	accounts, err := r.selectData(sql, params)
	if err != nil {
		return nil, err
	}
	return accounts, err
}

func (r *AccountRepository) FindById(propertyId int64) (schema.Accounts, error) {
	sql := `select
		property_id,
		plan_id,
		category_id,
		price,
		name,
		format_date("%Y-%m-%d", start_date) start_date,
		format_date("%Y-%m-%d", end_date) end_date,
		section from ` + config.GetTableName() + ` where property_id = @property_id `

	params := map[string]interface{}{
		"property_id": propertyId,
	}
	accounts, err := r.selectData(sql, params)
	if err != nil {
		return nil, err
	}
	return accounts, err
}

func (r *AccountRepository) selectData(sql string, params map[string]interface{}) (schema.Accounts, error) {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, config.GetDatabasePath())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	stmt := spanner.Statement{
		SQL:    sql,
		Params: params,
	}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()

	var accounts schema.Accounts
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var account schema.Account
		if err := row.Columns(
			&account.PROPERTY_ID,
			&account.PLAN_ID,
			&account.CATEGORY_ID,
			&account.PRICE,
			&account.NAME,
			&account.START_DATE,
			&account.END_DATE,
			&account.SECTION,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
