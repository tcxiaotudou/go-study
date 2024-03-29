// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-study/go-ent-pokemon/ent/battle"
	"go-study/go-ent-pokemon/ent/pokemon"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Battle is the model entity for the Battle schema.
type Battle struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"oid,omitempty"`
	// Result holds the value of the "result" field.
	Result string `json:"result,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BattleQuery when eager-loading is set.
	Edges             BattleEdges `json:"edges"`
	pokemon_fights    *int
	pokemon_opponents *int
}

// BattleEdges holds the relations/edges for other nodes in the graph.
type BattleEdges struct {
	// Contender holds the value of the contender edge.
	Contender *Pokemon `json:"contender,omitempty"`
	// Oponent holds the value of the oponent edge.
	Oponent *Pokemon `json:"oponent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ContenderOrErr returns the Contender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BattleEdges) ContenderOrErr() (*Pokemon, error) {
	if e.loadedTypes[0] {
		if e.Contender == nil {
			// The edge contender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pokemon.Label}
		}
		return e.Contender, nil
	}
	return nil, &NotLoadedError{edge: "contender"}
}

// OponentOrErr returns the Oponent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BattleEdges) OponentOrErr() (*Pokemon, error) {
	if e.loadedTypes[1] {
		if e.Oponent == nil {
			// The edge oponent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pokemon.Label}
		}
		return e.Oponent, nil
	}
	return nil, &NotLoadedError{edge: "oponent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Battle) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case battle.FieldID:
			values[i] = new(sql.NullInt64)
		case battle.FieldResult:
			values[i] = new(sql.NullString)
		case battle.FieldCreateAt, battle.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		case battle.ForeignKeys[0]: // pokemon_fights
			values[i] = new(sql.NullInt64)
		case battle.ForeignKeys[1]: // pokemon_opponents
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Battle", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Battle fields.
func (b *Battle) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case battle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case battle.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				b.Result = value.String
			}
		case battle.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				b.CreateAt = value.Time
			}
		case battle.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				b.UpdateAt = value.Time
			}
		case battle.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field pokemon_fights", value)
			} else if value.Valid {
				b.pokemon_fights = new(int)
				*b.pokemon_fights = int(value.Int64)
			}
		case battle.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field pokemon_opponents", value)
			} else if value.Valid {
				b.pokemon_opponents = new(int)
				*b.pokemon_opponents = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryContender queries the "contender" edge of the Battle entity.
func (b *Battle) QueryContender() *PokemonQuery {
	return (&BattleClient{config: b.config}).QueryContender(b)
}

// QueryOponent queries the "oponent" edge of the Battle entity.
func (b *Battle) QueryOponent() *PokemonQuery {
	return (&BattleClient{config: b.config}).QueryOponent(b)
}

// Update returns a builder for updating this Battle.
// Note that you need to call Battle.Unwrap() before calling this method if this Battle
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Battle) Update() *BattleUpdateOne {
	return (&BattleClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Battle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Battle) Unwrap() *Battle {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Battle is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Battle) String() string {
	var builder strings.Builder
	builder.WriteString("Battle(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("result=")
	builder.WriteString(b.Result)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(b.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(b.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Battles is a parsable slice of Battle.
type Battles []*Battle

func (b Battles) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
