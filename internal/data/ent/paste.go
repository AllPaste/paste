// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/AllPaste/paste/internal/data/ent/paste"
)

// Paste is the model entity for the Paste schema.
type Paste struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 剪切板内容
	Content string `json:"content,omitempty"`
	// 创建人(所属者)
	Creator string `json:"creator,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Paste) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case paste.FieldID, paste.FieldCreatedAt, paste.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case paste.FieldContent, paste.FieldCreator:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Paste", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Paste fields.
func (pa *Paste) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case paste.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int64(value.Int64)
		case paste.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Int64
			}
		case paste.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Int64
			}
		case paste.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pa.Content = value.String
			}
		case paste.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				pa.Creator = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Paste.
// Note that you need to call Paste.Unwrap() before calling this method if this Paste
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Paste) Update() *PasteUpdateOne {
	return (&PasteClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Paste entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Paste) Unwrap() *Paste {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Paste is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Paste) String() string {
	var builder strings.Builder
	builder.WriteString("Paste(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pa.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pa.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(pa.Content)
	builder.WriteString(", ")
	builder.WriteString("creator=")
	builder.WriteString(pa.Creator)
	builder.WriteByte(')')
	return builder.String()
}

// Pastes is a parsable slice of Paste.
type Pastes []*Paste

func (pa Pastes) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
