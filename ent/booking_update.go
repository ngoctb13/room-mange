// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"room-reservation/ent/booking"
	"room-reservation/ent/predicate"
	"room-reservation/ent/room"
	"room-reservation/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// Where appends a list predicates to the BookingUpdate builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetInformation sets the "information" field.
func (bu *BookingUpdate) SetInformation(s string) *BookingUpdate {
	bu.mutation.SetInformation(s)
	return bu
}

// SetNillableInformation sets the "information" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableInformation(s *string) *BookingUpdate {
	if s != nil {
		bu.SetInformation(*s)
	}
	return bu
}

// SetStartHour sets the "start_hour" field.
func (bu *BookingUpdate) SetStartHour(t time.Time) *BookingUpdate {
	bu.mutation.SetStartHour(t)
	return bu
}

// SetNillableStartHour sets the "start_hour" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableStartHour(t *time.Time) *BookingUpdate {
	if t != nil {
		bu.SetStartHour(*t)
	}
	return bu
}

// SetEndHour sets the "end_hour" field.
func (bu *BookingUpdate) SetEndHour(t time.Time) *BookingUpdate {
	bu.mutation.SetEndHour(t)
	return bu
}

// SetNillableEndHour sets the "end_hour" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableEndHour(t *time.Time) *BookingUpdate {
	if t != nil {
		bu.SetEndHour(*t)
	}
	return bu
}

// SetBookingDate sets the "booking_date" field.
func (bu *BookingUpdate) SetBookingDate(t time.Time) *BookingUpdate {
	bu.mutation.SetBookingDate(t)
	return bu
}

// SetNillableBookingDate sets the "booking_date" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableBookingDate(t *time.Time) *BookingUpdate {
	if t != nil {
		bu.SetBookingDate(*t)
	}
	return bu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bu *BookingUpdate) SetUserID(id int) *BookingUpdate {
	bu.mutation.SetUserID(id)
	return bu
}

// SetUser sets the "user" edge to the User entity.
func (bu *BookingUpdate) SetUser(u *User) *BookingUpdate {
	return bu.SetUserID(u.ID)
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (bu *BookingUpdate) SetRoomID(id int) *BookingUpdate {
	bu.mutation.SetRoomID(id)
	return bu
}

// SetRoom sets the "room" edge to the Room entity.
func (bu *BookingUpdate) SetRoom(r *Room) *BookingUpdate {
	return bu.SetRoomID(r.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (bu *BookingUpdate) ClearUser() *BookingUpdate {
	bu.mutation.ClearUser()
	return bu
}

// ClearRoom clears the "room" edge to the Room entity.
func (bu *BookingUpdate) ClearRoom() *BookingUpdate {
	bu.mutation.ClearRoom()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookingUpdate) check() error {
	if _, ok := bu.mutation.UserID(); bu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Booking.user"`)
	}
	if _, ok := bu.mutation.RoomID(); bu.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Booking.room"`)
	}
	return nil
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Information(); ok {
		_spec.SetField(booking.FieldInformation, field.TypeString, value)
	}
	if value, ok := bu.mutation.StartHour(); ok {
		_spec.SetField(booking.FieldStartHour, field.TypeTime, value)
	}
	if value, ok := bu.mutation.EndHour(); ok {
		_spec.SetField(booking.FieldEndHour, field.TypeTime, value)
	}
	if value, ok := bu.mutation.BookingDate(); ok {
		_spec.SetField(booking.FieldBookingDate, field.TypeTime, value)
	}
	if bu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UserTable,
			Columns: []string{booking.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UserTable,
			Columns: []string{booking.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.RoomTable,
			Columns: []string{booking.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.RoomTable,
			Columns: []string{booking.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookingMutation
}

// SetInformation sets the "information" field.
func (buo *BookingUpdateOne) SetInformation(s string) *BookingUpdateOne {
	buo.mutation.SetInformation(s)
	return buo
}

// SetNillableInformation sets the "information" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableInformation(s *string) *BookingUpdateOne {
	if s != nil {
		buo.SetInformation(*s)
	}
	return buo
}

// SetStartHour sets the "start_hour" field.
func (buo *BookingUpdateOne) SetStartHour(t time.Time) *BookingUpdateOne {
	buo.mutation.SetStartHour(t)
	return buo
}

// SetNillableStartHour sets the "start_hour" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableStartHour(t *time.Time) *BookingUpdateOne {
	if t != nil {
		buo.SetStartHour(*t)
	}
	return buo
}

// SetEndHour sets the "end_hour" field.
func (buo *BookingUpdateOne) SetEndHour(t time.Time) *BookingUpdateOne {
	buo.mutation.SetEndHour(t)
	return buo
}

// SetNillableEndHour sets the "end_hour" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableEndHour(t *time.Time) *BookingUpdateOne {
	if t != nil {
		buo.SetEndHour(*t)
	}
	return buo
}

// SetBookingDate sets the "booking_date" field.
func (buo *BookingUpdateOne) SetBookingDate(t time.Time) *BookingUpdateOne {
	buo.mutation.SetBookingDate(t)
	return buo
}

// SetNillableBookingDate sets the "booking_date" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableBookingDate(t *time.Time) *BookingUpdateOne {
	if t != nil {
		buo.SetBookingDate(*t)
	}
	return buo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (buo *BookingUpdateOne) SetUserID(id int) *BookingUpdateOne {
	buo.mutation.SetUserID(id)
	return buo
}

// SetUser sets the "user" edge to the User entity.
func (buo *BookingUpdateOne) SetUser(u *User) *BookingUpdateOne {
	return buo.SetUserID(u.ID)
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (buo *BookingUpdateOne) SetRoomID(id int) *BookingUpdateOne {
	buo.mutation.SetRoomID(id)
	return buo
}

// SetRoom sets the "room" edge to the Room entity.
func (buo *BookingUpdateOne) SetRoom(r *Room) *BookingUpdateOne {
	return buo.SetRoomID(r.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (buo *BookingUpdateOne) ClearUser() *BookingUpdateOne {
	buo.mutation.ClearUser()
	return buo
}

// ClearRoom clears the "room" edge to the Room entity.
func (buo *BookingUpdateOne) ClearRoom() *BookingUpdateOne {
	buo.mutation.ClearRoom()
	return buo
}

// Where appends a list predicates to the BookingUpdate builder.
func (buo *BookingUpdateOne) Where(ps ...predicate.Booking) *BookingUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookingUpdateOne) Select(field string, fields ...string) *BookingUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Booking entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookingUpdateOne) check() error {
	if _, ok := buo.mutation.UserID(); buo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Booking.user"`)
	}
	if _, ok := buo.mutation.RoomID(); buo.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Booking.room"`)
	}
	return nil
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (_node *Booking, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Booking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for _, f := range fields {
			if !booking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Information(); ok {
		_spec.SetField(booking.FieldInformation, field.TypeString, value)
	}
	if value, ok := buo.mutation.StartHour(); ok {
		_spec.SetField(booking.FieldStartHour, field.TypeTime, value)
	}
	if value, ok := buo.mutation.EndHour(); ok {
		_spec.SetField(booking.FieldEndHour, field.TypeTime, value)
	}
	if value, ok := buo.mutation.BookingDate(); ok {
		_spec.SetField(booking.FieldBookingDate, field.TypeTime, value)
	}
	if buo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UserTable,
			Columns: []string{booking.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UserTable,
			Columns: []string{booking.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.RoomTable,
			Columns: []string{booking.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.RoomTable,
			Columns: []string{booking.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Booking{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
