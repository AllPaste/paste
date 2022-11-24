package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type MetaObject struct {
	mixin.Schema
}

func (MetaObject) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique(),
		field.Int64("created_at").
			Immutable().
			DefaultFunc(time.Now().Unix).
			Comment("创建时间"),
		field.Int64("updated_at").
			DefaultFunc(time.Now().Unix).
			UpdateDefault(time.Now().Unix).
			Comment("更新时间"),
	}
}
