// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"room-reservation/ent/booking"
	"room-reservation/ent/predicate"
	"room-reservation/ent/room"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomUpdate is the builder for updating Room entities.
type RoomUpdate struct {
	config
	hooks    []Hook
	mutation *RoomMutation
}

// Where appends a list predicates to the RoomUpdate builder.
func (ru *RoomUpdate) Where(ps ...predicate.Room) *RoomUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RoomUpdate) SetName(s string) *RoomUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableName(s *string) *RoomUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetCapacity sets the "capacity" field.
func (ru *RoomUpdate) SetCapacity(i int) *RoomUpdate {
	ru.mutation.ResetCapacity()
	ru.mutation.SetCapacity(i)
	return ru
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableCapacity(i *int) *RoomUpdate {
	if i != nil {
		ru.SetCapacity(*i)
	}
	return ru
}

// AddCapacity adds i to the "capacity" field.
func (ru *RoomUpdate) AddCapacity(i int) *RoomUpdate {
	ru.mutation.AddCapacity(i)
	return ru
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (ru *RoomUpdate) AddBookingIDs(ids ...int) *RoomUpdate {
	ru.mutation.AddBookingIDs(ids...)
	return ru
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (ru *RoomUpdate) AddBookings(b ...*Booking) *RoomUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ru.AddBookingIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ru *RoomUpdate) Mutation() *RoomMutation {
	return ru.mutation
}

// ClearBookings clears all "bookings" edges to the Booking entity.
func (ru *RoomUpdate) ClearBookings() *RoomUpdate {
	ru.mutation.ClearBookings()
	return ru
}

// RemoveBookingIDs removes the "bookings" edge to Booking entities by IDs.
func (ru *RoomUpdate) RemoveBookingIDs(ids ...int) *RoomUpdate {
	ru.mutation.RemoveBookingIDs(ids...)
	return ru
}

// RemoveBookings removes "bookings" edges to Booking entities.
func (ru *RoomUpdate) RemoveBookings(b ...*Booking) *RoomUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ru.RemoveBookingIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoomUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoomUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := room.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Room.name": %w`, err)}
		}
	}
	return nil
}

func (ru *RoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(room.Table, room.Columns, sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(room.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Capacity(); ok {
		_spec.SetField(room.FieldCapacity, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedCapacity(); ok {
		_spec.AddField(room.FieldCapacity, field.TypeInt, value)
	}
	if ru.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.BookingsTable,
			Columns: []string{room.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedBookingsIDs(); len(nodes) > 0 && !ru.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.BookingsTable,
			Columns: []string{room.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.BookingsTable,
			Columns: []string{room.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoomUpdateOne is the builder for updating a single Room entity.
type RoomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoomMutation
}

// SetName sets the "name" field.
func (ruo *RoomUpdateOne) SetName(s string) *RoomUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableName(s *string) *RoomUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetCapacity sets the "capacity" field.
func (ruo *RoomUpdateOne) SetCapacity(i int) *RoomUpdateOne {
	ruo.mutation.ResetCapacity()
	ruo.mutation.SetCapacity(i)
	return ruo
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableCapacity(i *int) *RoomUpdateOne {
	if i != nil {
		ruo.SetCapacity(*i)
	}
	return ruo
}

// AddCapacity adds i to the "capacity" field.
func (ruo *RoomUpdateOne) AddCapacity(i int) *RoomUpdateOne {
	ruo.mutation.AddCapacity(i)
	return ruo
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (ruo *RoomUpdateOne) AddBookingIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.AddBookingIDs(ids...)
	return ruo
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (ruo *RoomUpdateOne) AddBookings(b ...*Booking) *RoomUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ruo.AddBookingIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ruo *RoomUpdateOne) Mutation() *RoomMutation {
	return ruo.mutation
}

// ClearBookings clears all "bookings" edges to the Booking entity.
func (ruo *RoomUpdateOne) ClearBookings() *RoomUpdateOne {
	ruo.mutation.ClearBookings()
	return ruo
}

// RemoveBookingIDs removes the "bookings" edge to Booking entities by IDs.
func (ruo *RoomUpdateOne) RemoveBookingIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.RemoveBookingIDs(ids...)
	return ruo
}

// RemoveBookings removes "bookings" edges to Booking entities.
func (ruo *RoomUpdateOne) RemoveBookings(b ...*Booking) *RoomUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ruo.RemoveBookingIDs(ids...)
}

// Where appends a list predicates to the RoomUpdate builder.
func (ruo *RoomUpdateOne) Where(ps ...predicate.Room) *RoomUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoomUpdateOne) Select(field string, fields ...string) *RoomUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Room entity.
func (ruo *RoomUpdateOne) Save(ctx context.Context) (*Room, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomUpdateOne) SaveX(ctx context.Context) *Room {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoomUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := room.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Room.name": %w`, err)}
		}
	}
	return nil
}

func (ruo *RoomUpdateOne) sqlSave(ctx context.Context) (_node *Room, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(room.Table, room.Columns, sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Room.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, room.FieldID)
		for _, f := range fields {
			if !room.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != room.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(room.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Capacity(); ok {
		_spec.SetField(room.FieldCapacity, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedCapacity(); ok {
		_spec.AddField(room.FieldCapacity, field.TypeInt, value)
	}
	if ruo.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.BookingsTable,
			Columns: []string{room.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedBookingsIDs(); len(nodes) > 0 && !ruo.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.BookingsTable,
			Columns: []string{room.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.BookingsTable,
			Columns: []string{room.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Room{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}