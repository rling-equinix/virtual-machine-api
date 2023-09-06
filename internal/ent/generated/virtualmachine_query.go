// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/predicate"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachine"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachinecpuconfig"
	"go.infratographer.com/x/gidx"
)

// VirtualMachineQuery is the builder for querying VirtualMachine entities.
type VirtualMachineQuery struct {
	config
	ctx                         *QueryContext
	order                       []virtualmachine.OrderOption
	inters                      []Interceptor
	predicates                  []predicate.VirtualMachine
	withVirtualMachineCPUConfig *VirtualMachineCPUConfigQuery
	modifiers                   []func(*sql.Selector)
	loadTotal                   []func(context.Context, []*VirtualMachine) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VirtualMachineQuery builder.
func (vmq *VirtualMachineQuery) Where(ps ...predicate.VirtualMachine) *VirtualMachineQuery {
	vmq.predicates = append(vmq.predicates, ps...)
	return vmq
}

// Limit the number of records to be returned by this query.
func (vmq *VirtualMachineQuery) Limit(limit int) *VirtualMachineQuery {
	vmq.ctx.Limit = &limit
	return vmq
}

// Offset to start from.
func (vmq *VirtualMachineQuery) Offset(offset int) *VirtualMachineQuery {
	vmq.ctx.Offset = &offset
	return vmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vmq *VirtualMachineQuery) Unique(unique bool) *VirtualMachineQuery {
	vmq.ctx.Unique = &unique
	return vmq
}

// Order specifies how the records should be ordered.
func (vmq *VirtualMachineQuery) Order(o ...virtualmachine.OrderOption) *VirtualMachineQuery {
	vmq.order = append(vmq.order, o...)
	return vmq
}

// QueryVirtualMachineCPUConfig chains the current query on the "virtual_machine_cpu_config" edge.
func (vmq *VirtualMachineQuery) QueryVirtualMachineCPUConfig() *VirtualMachineCPUConfigQuery {
	query := (&VirtualMachineCPUConfigClient{config: vmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(virtualmachine.Table, virtualmachine.FieldID, selector),
			sqlgraph.To(virtualmachinecpuconfig.Table, virtualmachinecpuconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, virtualmachine.VirtualMachineCPUConfigTable, virtualmachine.VirtualMachineCPUConfigColumn),
		)
		fromU = sqlgraph.SetNeighbors(vmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VirtualMachine entity from the query.
// Returns a *NotFoundError when no VirtualMachine was found.
func (vmq *VirtualMachineQuery) First(ctx context.Context) (*VirtualMachine, error) {
	nodes, err := vmq.Limit(1).All(setContextOp(ctx, vmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{virtualmachine.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vmq *VirtualMachineQuery) FirstX(ctx context.Context) *VirtualMachine {
	node, err := vmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VirtualMachine ID from the query.
// Returns a *NotFoundError when no VirtualMachine ID was found.
func (vmq *VirtualMachineQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = vmq.Limit(1).IDs(setContextOp(ctx, vmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{virtualmachine.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vmq *VirtualMachineQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := vmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VirtualMachine entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VirtualMachine entity is found.
// Returns a *NotFoundError when no VirtualMachine entities are found.
func (vmq *VirtualMachineQuery) Only(ctx context.Context) (*VirtualMachine, error) {
	nodes, err := vmq.Limit(2).All(setContextOp(ctx, vmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{virtualmachine.Label}
	default:
		return nil, &NotSingularError{virtualmachine.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vmq *VirtualMachineQuery) OnlyX(ctx context.Context) *VirtualMachine {
	node, err := vmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VirtualMachine ID in the query.
// Returns a *NotSingularError when more than one VirtualMachine ID is found.
// Returns a *NotFoundError when no entities are found.
func (vmq *VirtualMachineQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = vmq.Limit(2).IDs(setContextOp(ctx, vmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{virtualmachine.Label}
	default:
		err = &NotSingularError{virtualmachine.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vmq *VirtualMachineQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := vmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VirtualMachines.
func (vmq *VirtualMachineQuery) All(ctx context.Context) ([]*VirtualMachine, error) {
	ctx = setContextOp(ctx, vmq.ctx, "All")
	if err := vmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VirtualMachine, *VirtualMachineQuery]()
	return withInterceptors[[]*VirtualMachine](ctx, vmq, qr, vmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vmq *VirtualMachineQuery) AllX(ctx context.Context) []*VirtualMachine {
	nodes, err := vmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VirtualMachine IDs.
func (vmq *VirtualMachineQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if vmq.ctx.Unique == nil && vmq.path != nil {
		vmq.Unique(true)
	}
	ctx = setContextOp(ctx, vmq.ctx, "IDs")
	if err = vmq.Select(virtualmachine.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vmq *VirtualMachineQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := vmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vmq *VirtualMachineQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vmq.ctx, "Count")
	if err := vmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vmq, querierCount[*VirtualMachineQuery](), vmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vmq *VirtualMachineQuery) CountX(ctx context.Context) int {
	count, err := vmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vmq *VirtualMachineQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vmq.ctx, "Exist")
	switch _, err := vmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vmq *VirtualMachineQuery) ExistX(ctx context.Context) bool {
	exist, err := vmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VirtualMachineQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vmq *VirtualMachineQuery) Clone() *VirtualMachineQuery {
	if vmq == nil {
		return nil
	}
	return &VirtualMachineQuery{
		config:                      vmq.config,
		ctx:                         vmq.ctx.Clone(),
		order:                       append([]virtualmachine.OrderOption{}, vmq.order...),
		inters:                      append([]Interceptor{}, vmq.inters...),
		predicates:                  append([]predicate.VirtualMachine{}, vmq.predicates...),
		withVirtualMachineCPUConfig: vmq.withVirtualMachineCPUConfig.Clone(),
		// clone intermediate query.
		sql:  vmq.sql.Clone(),
		path: vmq.path,
	}
}

// WithVirtualMachineCPUConfig tells the query-builder to eager-load the nodes that are connected to
// the "virtual_machine_cpu_config" edge. The optional arguments are used to configure the query builder of the edge.
func (vmq *VirtualMachineQuery) WithVirtualMachineCPUConfig(opts ...func(*VirtualMachineCPUConfigQuery)) *VirtualMachineQuery {
	query := (&VirtualMachineCPUConfigClient{config: vmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vmq.withVirtualMachineCPUConfig = query
	return vmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VirtualMachine.Query().
//		GroupBy(virtualmachine.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (vmq *VirtualMachineQuery) GroupBy(field string, fields ...string) *VirtualMachineGroupBy {
	vmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VirtualMachineGroupBy{build: vmq}
	grbuild.flds = &vmq.ctx.Fields
	grbuild.label = virtualmachine.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.VirtualMachine.Query().
//		Select(virtualmachine.FieldCreatedAt).
//		Scan(ctx, &v)
func (vmq *VirtualMachineQuery) Select(fields ...string) *VirtualMachineSelect {
	vmq.ctx.Fields = append(vmq.ctx.Fields, fields...)
	sbuild := &VirtualMachineSelect{VirtualMachineQuery: vmq}
	sbuild.label = virtualmachine.Label
	sbuild.flds, sbuild.scan = &vmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VirtualMachineSelect configured with the given aggregations.
func (vmq *VirtualMachineQuery) Aggregate(fns ...AggregateFunc) *VirtualMachineSelect {
	return vmq.Select().Aggregate(fns...)
}

func (vmq *VirtualMachineQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vmq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vmq); err != nil {
				return err
			}
		}
	}
	for _, f := range vmq.ctx.Fields {
		if !virtualmachine.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if vmq.path != nil {
		prev, err := vmq.path(ctx)
		if err != nil {
			return err
		}
		vmq.sql = prev
	}
	return nil
}

func (vmq *VirtualMachineQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VirtualMachine, error) {
	var (
		nodes       = []*VirtualMachine{}
		_spec       = vmq.querySpec()
		loadedTypes = [1]bool{
			vmq.withVirtualMachineCPUConfig != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VirtualMachine).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VirtualMachine{config: vmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vmq.modifiers) > 0 {
		_spec.Modifiers = vmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vmq.withVirtualMachineCPUConfig; query != nil {
		if err := vmq.loadVirtualMachineCPUConfig(ctx, query, nodes, nil,
			func(n *VirtualMachine, e *VirtualMachineCPUConfig) { n.Edges.VirtualMachineCPUConfig = e }); err != nil {
			return nil, err
		}
	}
	for i := range vmq.loadTotal {
		if err := vmq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vmq *VirtualMachineQuery) loadVirtualMachineCPUConfig(ctx context.Context, query *VirtualMachineCPUConfigQuery, nodes []*VirtualMachine, init func(*VirtualMachine), assign func(*VirtualMachine, *VirtualMachineCPUConfig)) error {
	ids := make([]gidx.PrefixedID, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID][]*VirtualMachine)
	for i := range nodes {
		fk := nodes[i].VMCPUConfigID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(virtualmachinecpuconfig.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "vm_cpu_config_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vmq *VirtualMachineQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vmq.querySpec()
	if len(vmq.modifiers) > 0 {
		_spec.Modifiers = vmq.modifiers
	}
	_spec.Node.Columns = vmq.ctx.Fields
	if len(vmq.ctx.Fields) > 0 {
		_spec.Unique = vmq.ctx.Unique != nil && *vmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vmq.driver, _spec)
}

func (vmq *VirtualMachineQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(virtualmachine.Table, virtualmachine.Columns, sqlgraph.NewFieldSpec(virtualmachine.FieldID, field.TypeString))
	_spec.From = vmq.sql
	if unique := vmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vmq.path != nil {
		_spec.Unique = true
	}
	if fields := vmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, virtualmachine.FieldID)
		for i := range fields {
			if fields[i] != virtualmachine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if vmq.withVirtualMachineCPUConfig != nil {
			_spec.Node.AddColumnOnce(virtualmachine.FieldVMCPUConfigID)
		}
	}
	if ps := vmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vmq *VirtualMachineQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vmq.driver.Dialect())
	t1 := builder.Table(virtualmachine.Table)
	columns := vmq.ctx.Fields
	if len(columns) == 0 {
		columns = virtualmachine.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vmq.sql != nil {
		selector = vmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vmq.ctx.Unique != nil && *vmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vmq.predicates {
		p(selector)
	}
	for _, p := range vmq.order {
		p(selector)
	}
	if offset := vmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VirtualMachineGroupBy is the group-by builder for VirtualMachine entities.
type VirtualMachineGroupBy struct {
	selector
	build *VirtualMachineQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vmgb *VirtualMachineGroupBy) Aggregate(fns ...AggregateFunc) *VirtualMachineGroupBy {
	vmgb.fns = append(vmgb.fns, fns...)
	return vmgb
}

// Scan applies the selector query and scans the result into the given value.
func (vmgb *VirtualMachineGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vmgb.build.ctx, "GroupBy")
	if err := vmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VirtualMachineQuery, *VirtualMachineGroupBy](ctx, vmgb.build, vmgb, vmgb.build.inters, v)
}

func (vmgb *VirtualMachineGroupBy) sqlScan(ctx context.Context, root *VirtualMachineQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vmgb.fns))
	for _, fn := range vmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vmgb.flds)+len(vmgb.fns))
		for _, f := range *vmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VirtualMachineSelect is the builder for selecting fields of VirtualMachine entities.
type VirtualMachineSelect struct {
	*VirtualMachineQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vms *VirtualMachineSelect) Aggregate(fns ...AggregateFunc) *VirtualMachineSelect {
	vms.fns = append(vms.fns, fns...)
	return vms
}

// Scan applies the selector query and scans the result into the given value.
func (vms *VirtualMachineSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vms.ctx, "Select")
	if err := vms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VirtualMachineQuery, *VirtualMachineSelect](ctx, vms.VirtualMachineQuery, vms, vms.inters, v)
}

func (vms *VirtualMachineSelect) sqlScan(ctx context.Context, root *VirtualMachineQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vms.fns))
	for _, fn := range vms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
