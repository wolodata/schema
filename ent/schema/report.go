package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Report struct {
	ent.Schema
}

func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Immutable(),
		field.String("report_type").NotEmpty().Immutable(),
		field.Uint64("trigger_user_id").Optional().Immutable(),
		field.Time("trigger_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.JSON("related_article_ids", []uint64{}).Default([]uint64{}).Immutable(),
		field.Text("content").Default(""),
		field.Text("reason").Default(""),
		field.Time("generated_at").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")),
	}
}

func (Report) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("report_type"),
		index.Fields("trigger_user_id"),
		index.Fields("trigger_at"),
		index.Fields("generated_at"),
	}
}

func (Report) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_report",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
