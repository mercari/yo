// Code generated by yo. DO NOT EDIT.

// Package customtypes contains the types.
package customtypes

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// Item represents a row from 'Items'.
type Item struct {
	ID    int64 `spanner:"ID" json:"ID"`       // ID
	Price int64 `spanner:"Price" json:"Price"` // Price
}

func ItemPrimaryKeys() []string {
	return []string{
		"ID",
	}
}

func ItemColumns() []string {
	return []string{
		"ID",
		"Price",
	}
}

func (i *Item) columnsToPtrs(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "ID":
			ret = append(ret, yoDecode(&i.ID))
		case "Price":
			ret = append(ret, yoDecode(&i.Price))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (i *Item) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "ID":
			ret = append(ret, yoEncode(i.ID))
		case "Price":
			ret = append(ret, yoEncode(i.Price))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newItem_Decoder returns a decoder which reads a row from *spanner.Row
// into Item. The decoder is not goroutine-safe. Don't use it concurrently.
func newItem_Decoder(cols []string) func(*spanner.Row) (*Item, error) {
	return func(row *spanner.Row) (*Item, error) {
		var i Item
		ptrs, err := i.columnsToPtrs(cols)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &i, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (i *Item) Insert(ctx context.Context) *spanner.Mutation {
	values, _ := i.columnsToValues(ItemColumns())
	return spanner.Insert("Items", ItemColumns(), values)
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (i *Item) Update(ctx context.Context) *spanner.Mutation {
	values, _ := i.columnsToValues(ItemColumns())
	return spanner.Update("Items", ItemColumns(), values)
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (i *Item) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	values, _ := i.columnsToValues(ItemColumns())
	return spanner.InsertOrUpdate("Items", ItemColumns(), values)
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (i *Item) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, ItemPrimaryKeys()...)

	values, err := i.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "Item.UpdateColumns", "Items", err)
	}

	return spanner.Update("Items", colsWithPKeys, values), nil
}

// FindItem gets a Item by primary key
func FindItem(ctx context.Context, db YORODB, id int64) (*Item, error) {
	key := spanner.Key{id}
	row, err := db.ReadRow(ctx, "Items", key, ItemColumns())
	if err != nil {
		return nil, newError("FindItem", "Items", err)
	}

	decoder := newItem_Decoder(ItemColumns())
	i, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindItem", "Items", err)
	}

	return i, nil
}

// ReadItem retrieves multiples rows from Item by KeySet as a slice.
func ReadItem(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*Item, error) {
	var res []*Item

	decoder := newItem_Decoder(ItemColumns())

	rows := db.Read(ctx, "Items", keys, ItemColumns())
	err := rows.Do(func(row *spanner.Row) error {
		i, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, i)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadItem", "Items", err)
	}

	return res, nil
}

// Delete deletes the Item from the database.
func (i *Item) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := i.columnsToValues(ItemPrimaryKeys())
	return spanner.Delete("Items", spanner.Key(values))
}
