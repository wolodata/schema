package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable(),
		field.String("username").Unique().Immutable(),
		field.String("password").Immutable(),
		field.Bool("is_admin").Default(false).Immutable(),
		field.Enum("level").Values("Air", "Pro", "Max").Default("Air"),
		field.Bool("enabled").Default(true),
		field.String("nickname"),
		field.String("description").Default(""),
		field.Time("created_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("is_admin"),
		index.Fields("level"),
		index.Fields("enabled"),
		index.Fields("created_at"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_user",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
