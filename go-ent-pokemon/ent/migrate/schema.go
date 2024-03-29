// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BattlesColumns holds the columns for the "battles" table.
	BattlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "result", Type: field.TypeString, Size: 2147483647},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "pokemon_fights", Type: field.TypeInt, Nullable: true},
		{Name: "pokemon_opponents", Type: field.TypeInt, Nullable: true},
	}
	// BattlesTable holds the schema information for the "battles" table.
	BattlesTable = &schema.Table{
		Name:       "battles",
		Columns:    BattlesColumns,
		PrimaryKey: []*schema.Column{BattlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "battles_pokemons_fights",
				Columns:    []*schema.Column{BattlesColumns[4]},
				RefColumns: []*schema.Column{PokemonsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "battles_pokemons_opponents",
				Columns:    []*schema.Column{BattlesColumns[5]},
				RefColumns: []*schema.Column{PokemonsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PokemonsColumns holds the columns for the "pokemons" table.
	PokemonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "weight", Type: field.TypeFloat64},
		{Name: "height", Type: field.TypeFloat64},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// PokemonsTable holds the schema information for the "pokemons" table.
	PokemonsTable = &schema.Table{
		Name:       "pokemons",
		Columns:    PokemonsColumns,
		PrimaryKey: []*schema.Column{PokemonsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BattlesTable,
		PokemonsTable,
	}
)

func init() {
	BattlesTable.ForeignKeys[0].RefTable = PokemonsTable
	BattlesTable.ForeignKeys[1].RefTable = PokemonsTable
}
