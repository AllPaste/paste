// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnsibleColumns holds the columns for the "ansible" table.
	AnsibleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "creator", Type: field.TypeString},
	}
	// AnsibleTable holds the schema information for the "ansible" table.
	AnsibleTable = &schema.Table{
		Name:       "ansible",
		Columns:    AnsibleColumns,
		PrimaryKey: []*schema.Column{AnsibleColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnsibleTable,
	}
)

func init() {
	AnsibleTable.Annotation = &entsql.Annotation{
		Table: "ansible",
	}
}
