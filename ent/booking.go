// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"room-reservation/ent/booking"
	"room-reservation/ent/room"
	"room-reservation/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Booking is the model entity for the Booking schema.
type Booking struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Information holds the value of the "information" field.
	Information string `json:"information,omitempty"`
	// StartHour holds the value of the "start_hour" field.
	StartHour time.Time `json:"start_hour,omitempty"`
	// EndHour holds the value of the "end_hour" field.
	EndHour time.Time `json:"end_hour,omitempty"`
	// BookingDate holds the value of the "booking_date" field.
	BookingDate time.Time `json:"booking_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookingQuery when eager-loading is set.
	Edges         BookingEdges `json:"edges"`
	room_bookings *int
	user_bookings *int
	selectValues  sql.SelectValues
}

// BookingEdges holds the relations/edges for other nodes in the graph.
type BookingEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Room holds the value of the room edge.
	Room *Room `json:"room,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) RoomOrErr() (*Room, error) {
	if e.Room != nil {
		return e.Room, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: room.Label}
	}
	return nil, &NotLoadedError{edge: "room"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Booking) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case booking.FieldID:
			values[i] = new(sql.NullInt64)
		case booking.FieldInformation:
			values[i] = new(sql.NullString)
		case booking.FieldStartHour, booking.FieldEndHour, booking.FieldBookingDate:
			values[i] = new(sql.NullTime)
		case booking.ForeignKeys[0]: // room_bookings
			values[i] = new(sql.NullInt64)
		case booking.ForeignKeys[1]: // user_bookings
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Booking fields.
func (b *Booking) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case booking.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case booking.FieldInformation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field information", values[i])
			} else if value.Valid {
				b.Information = value.String
			}
		case booking.FieldStartHour:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_hour", values[i])
			} else if value.Valid {
				b.StartHour = value.Time
			}
		case booking.FieldEndHour:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_hour", values[i])
			} else if value.Valid {
				b.EndHour = value.Time
			}
		case booking.FieldBookingDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field booking_date", values[i])
			} else if value.Valid {
				b.BookingDate = value.Time
			}
		case booking.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field room_bookings", value)
			} else if value.Valid {
				b.room_bookings = new(int)
				*b.room_bookings = int(value.Int64)
			}
		case booking.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_bookings", value)
			} else if value.Valid {
				b.user_bookings = new(int)
				*b.user_bookings = int(value.Int64)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Booking.
// This includes values selected through modifiers, order, etc.
func (b *Booking) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Booking entity.
func (b *Booking) QueryUser() *UserQuery {
	return NewBookingClient(b.config).QueryUser(b)
}

// QueryRoom queries the "room" edge of the Booking entity.
func (b *Booking) QueryRoom() *RoomQuery {
	return NewBookingClient(b.config).QueryRoom(b)
}

// Update returns a builder for updating this Booking.
// Note that you need to call Booking.Unwrap() before calling this method if this Booking
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Booking) Update() *BookingUpdateOne {
	return NewBookingClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Booking entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Booking) Unwrap() *Booking {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Booking is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Booking) String() string {
	var builder strings.Builder
	builder.WriteString("Booking(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("information=")
	builder.WriteString(b.Information)
	builder.WriteString(", ")
	builder.WriteString("start_hour=")
	builder.WriteString(b.StartHour.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_hour=")
	builder.WriteString(b.EndHour.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("booking_date=")
	builder.WriteString(b.BookingDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Bookings is a parsable slice of Booking.
type Bookings []*Booking