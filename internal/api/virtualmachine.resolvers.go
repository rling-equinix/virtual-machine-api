package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"go.infratographer.com/virtual-machine-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// VirtualMachine is the resolver for the virtualMachine field.
func (r *queryResolver) VirtualMachine(ctx context.Context, id gidx.PrefixedID) (*generated.VirtualMachine, error) {
	// TODO check locations
	// if err := permissions.CheckAccess(ctx, id, actionVirtualMachineGet); err != nil {
	// 	return nil, err
	// }

	return r.client.VirtualMachine.Get(ctx, id)
}
