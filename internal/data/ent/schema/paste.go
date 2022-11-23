package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Paste holds the schema definition for the User entity.
type Paste struct {
	ent.Schema
}

// Fields of the User.
func (Paste) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content").Comment("剪切板内容"),
		field.String("creator").Comment("创建人(所属者)"),
	}
}

// Edges of the User.
func (Paste) Edges() []ent.Edge {
	return nil
}
