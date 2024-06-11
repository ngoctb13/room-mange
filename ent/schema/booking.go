package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.String("information"),
		field.Time("start_hour").
			SchemaType(map[string]string{
				"postgres": "TIME",
			}),
		field.Time("end_hour").
			SchemaType(map[string]string{
				"postgres": "TIME",
			}),
		field.Time("booking_date").Default(time.Now()),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("bookings").Unique().Required(),
		edge.From("room", Room.Type).Ref("bookings").Unique().Required(),
	}
}
