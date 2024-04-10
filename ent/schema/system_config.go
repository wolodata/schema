package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SystemConfig struct {
	ent.Schema
}

func (SystemConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable(),
		field.String("name").Unique().Immutable(),
		field.String("description").Default(""),
		field.String("api_model").Default(""),
		field.String("api_url").Default(""),
		field.String("api_key").Default(""),
		field.Text("prompt_system").Default(""),
		field.Text("prompt_user").Default(""),
	}
}

func (SystemConfig) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SystemConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_system_config",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
