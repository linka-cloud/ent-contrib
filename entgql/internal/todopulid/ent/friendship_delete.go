// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql/internal/todopulid/ent/friendship"
	"entgo.io/contrib/entgql/internal/todopulid/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipDelete is the builder for deleting a Friendship entity.
type FriendshipDelete struct {
	config
	hooks    []Hook
	mutation *FriendshipMutation
}

// Where appends a list predicates to the FriendshipDelete builder.
func (fd *FriendshipDelete) Where(ps ...predicate.Friendship) *FriendshipDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FriendshipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, FriendshipMutation](ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FriendshipDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FriendshipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: friendship.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: friendship.FieldID,
			},
		},
	}
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FriendshipDeleteOne is the builder for deleting a single Friendship entity.
type FriendshipDeleteOne struct {
	fd *FriendshipDelete
}

// Exec executes the deletion query.
func (fdo *FriendshipDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{friendship.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FriendshipDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}
