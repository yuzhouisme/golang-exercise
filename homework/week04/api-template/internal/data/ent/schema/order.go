package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("client_id"),
		field.String("config_type"),
		field.String("room_no"),
		field.String("buyer"),
		field.Float32 ("final_price"),
		field.Float32("actual_price"),
		field.Time("deposit_date").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("transaction_date").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Float32("fee_scale"),
		field.Float32("discharge_fee"),
		field.Float32("forward_fee"),
		field.Float32("receivable_fee"),
		field.Float32("invoiced"),
		field.Float32("not_invoiced"),
		field.Float32("received"),
		field.Float32("not_received"),
		field.String("status"),
		field.String("reserve"),
		field.Float32("reserve_price"),
		field.String("is_delete"),
		field.String("remark"),
		field.String("paper"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.String("created_by"),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.String("updated_by"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("orders"),
	}
}
