// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BookingsColumns holds the columns for the "bookings" table.
	BookingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "information", Type: field.TypeString},
		{Name: "start_hour", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "TIME"}},
		{Name: "end_hour", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "TIME"}},
		{Name: "booking_date", Type: field.TypeTime},
		{Name: "room_bookings", Type: field.TypeInt},
		{Name: "user_bookings", Type: field.TypeInt},
	}
	// BookingsTable holds the schema information for the "bookings" table.
	BookingsTable = &schema.Table{
		Name:       "bookings",
		Columns:    BookingsColumns,
		PrimaryKey: []*schema.Column{BookingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bookings_rooms_bookings",
				Columns:    []*schema.Column{BookingsColumns[5]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "bookings_users_bookings",
				Columns:    []*schema.Column{BookingsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "capacity", Type: field.TypeInt},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BookingsTable,
		RoomsTable,
		UsersTable,
	}
)

func init() {
	BookingsTable.ForeignKeys[0].RefTable = RoomsTable
	BookingsTable.ForeignKeys[1].RefTable = UsersTable
}