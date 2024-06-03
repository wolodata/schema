package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type PromotConfig struct {
	ent.Schema
}

func (PromotConfig) Fields() []ent.Field {
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

func (PromotConfig) Indexes() []ent.Index {
	return []ent.Index{}
}

func (PromotConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_prompt_config",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
