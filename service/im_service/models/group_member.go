// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package im

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// GroupMember is an object representing the database table.
type GroupMember struct {
	GroupID   int64  `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	UID       int64  `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Timestamp int    `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Nickname  string `boil:"nickname" json:"nickname" toml:"nickname" yaml:"nickname"`
	Mute      bool   `boil:"mute" json:"mute" toml:"mute" yaml:"mute"`

	R *groupMemberR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L groupMemberL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GroupMemberColumns = struct {
	GroupID   string
	UID       string
	Timestamp string
	Nickname  string
	Mute      string
}{
	GroupID:   "group_id",
	UID:       "uid",
	Timestamp: "timestamp",
	Nickname:  "nickname",
	Mute:      "mute",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var GroupMemberWhere = struct {
	GroupID   whereHelperint64
	UID       whereHelperint64
	Timestamp whereHelperint
	Nickname  whereHelperstring
	Mute      whereHelperbool
}{
	GroupID:   whereHelperint64{field: "`group_member`.`group_id`"},
	UID:       whereHelperint64{field: "`group_member`.`uid`"},
	Timestamp: whereHelperint{field: "`group_member`.`timestamp`"},
	Nickname:  whereHelperstring{field: "`group_member`.`nickname`"},
	Mute:      whereHelperbool{field: "`group_member`.`mute`"},
}

// GroupMemberRels is where relationship names are stored.
var GroupMemberRels = struct {
}{}

// groupMemberR is where relationships are stored.
type groupMemberR struct {
}

// NewStruct creates a new relationship struct
func (*groupMemberR) NewStruct() *groupMemberR {
	return &groupMemberR{}
}

// groupMemberL is where Load methods for each relationship are stored.
type groupMemberL struct{}

var (
	groupMemberAllColumns            = []string{"group_id", "uid", "timestamp", "nickname", "mute"}
	groupMemberColumnsWithoutDefault = []string{"timestamp", "nickname"}
	groupMemberColumnsWithDefault    = []string{"group_id", "uid", "mute"}
	groupMemberPrimaryKeyColumns     = []string{"group_id", "uid"}
)

type (
	// GroupMemberSlice is an alias for a slice of pointers to GroupMember.
	// This should generally be used opposed to []GroupMember.
	GroupMemberSlice []*GroupMember
	// GroupMemberHook is the signature for custom GroupMember hook methods
	GroupMemberHook func(context.Context, boil.ContextExecutor, *GroupMember) error

	groupMemberQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	groupMemberType                 = reflect.TypeOf(&GroupMember{})
	groupMemberMapping              = queries.MakeStructMapping(groupMemberType)
	groupMemberPrimaryKeyMapping, _ = queries.BindMapping(groupMemberType, groupMemberMapping, groupMemberPrimaryKeyColumns)
	groupMemberInsertCacheMut       sync.RWMutex
	groupMemberInsertCache          = make(map[string]insertCache)
	groupMemberUpdateCacheMut       sync.RWMutex
	groupMemberUpdateCache          = make(map[string]updateCache)
	groupMemberUpsertCacheMut       sync.RWMutex
	groupMemberUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var groupMemberBeforeInsertHooks []GroupMemberHook
var groupMemberBeforeUpdateHooks []GroupMemberHook
var groupMemberBeforeDeleteHooks []GroupMemberHook
var groupMemberBeforeUpsertHooks []GroupMemberHook

var groupMemberAfterInsertHooks []GroupMemberHook
var groupMemberAfterSelectHooks []GroupMemberHook
var groupMemberAfterUpdateHooks []GroupMemberHook
var groupMemberAfterDeleteHooks []GroupMemberHook
var groupMemberAfterUpsertHooks []GroupMemberHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GroupMember) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GroupMember) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GroupMember) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GroupMember) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GroupMember) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GroupMember) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GroupMember) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GroupMember) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GroupMember) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupMemberAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGroupMemberHook registers your hook function for all future operations.
func AddGroupMemberHook(hookPoint boil.HookPoint, groupMemberHook GroupMemberHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		groupMemberBeforeInsertHooks = append(groupMemberBeforeInsertHooks, groupMemberHook)
	case boil.BeforeUpdateHook:
		groupMemberBeforeUpdateHooks = append(groupMemberBeforeUpdateHooks, groupMemberHook)
	case boil.BeforeDeleteHook:
		groupMemberBeforeDeleteHooks = append(groupMemberBeforeDeleteHooks, groupMemberHook)
	case boil.BeforeUpsertHook:
		groupMemberBeforeUpsertHooks = append(groupMemberBeforeUpsertHooks, groupMemberHook)
	case boil.AfterInsertHook:
		groupMemberAfterInsertHooks = append(groupMemberAfterInsertHooks, groupMemberHook)
	case boil.AfterSelectHook:
		groupMemberAfterSelectHooks = append(groupMemberAfterSelectHooks, groupMemberHook)
	case boil.AfterUpdateHook:
		groupMemberAfterUpdateHooks = append(groupMemberAfterUpdateHooks, groupMemberHook)
	case boil.AfterDeleteHook:
		groupMemberAfterDeleteHooks = append(groupMemberAfterDeleteHooks, groupMemberHook)
	case boil.AfterUpsertHook:
		groupMemberAfterUpsertHooks = append(groupMemberAfterUpsertHooks, groupMemberHook)
	}
}

// One returns a single groupMember record from the query.
func (q groupMemberQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GroupMember, error) {
	o := &GroupMember{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "im: failed to execute a one query for group_member")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GroupMember records from the query.
func (q groupMemberQuery) All(ctx context.Context, exec boil.ContextExecutor) (GroupMemberSlice, error) {
	var o []*GroupMember

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "im: failed to assign all query results to GroupMember slice")
	}

	if len(groupMemberAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GroupMember records in the query.
func (q groupMemberQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to count group_member rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q groupMemberQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "im: failed to check if group_member exists")
	}

	return count > 0, nil
}

// GroupMembers retrieves all the records using an executor.
func GroupMembers(mods ...qm.QueryMod) groupMemberQuery {
	mods = append(mods, qm.From("`group_member`"))
	return groupMemberQuery{NewQuery(mods...)}
}

// FindGroupMember retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGroupMember(ctx context.Context, exec boil.ContextExecutor, groupID int64, uID int64, selectCols ...string) (*GroupMember, error) {
	groupMemberObj := &GroupMember{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `group_member` where `group_id`=? AND `uid`=?", sel,
	)

	q := queries.Raw(query, groupID, uID)

	err := q.Bind(ctx, exec, groupMemberObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "im: unable to select from group_member")
	}

	return groupMemberObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GroupMember) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("im: no group_member provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(groupMemberColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	groupMemberInsertCacheMut.RLock()
	cache, cached := groupMemberInsertCache[key]
	groupMemberInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			groupMemberAllColumns,
			groupMemberColumnsWithDefault,
			groupMemberColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(groupMemberType, groupMemberMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(groupMemberType, groupMemberMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `group_member` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `group_member` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `group_member` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, groupMemberPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "im: unable to insert into group_member")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.GroupID,
		o.UID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "im: unable to populate default values for group_member")
	}

CacheNoHooks:
	if !cached {
		groupMemberInsertCacheMut.Lock()
		groupMemberInsertCache[key] = cache
		groupMemberInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GroupMember.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GroupMember) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	groupMemberUpdateCacheMut.RLock()
	cache, cached := groupMemberUpdateCache[key]
	groupMemberUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			groupMemberAllColumns,
			groupMemberPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("im: unable to update group_member, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `group_member` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, groupMemberPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(groupMemberType, groupMemberMapping, append(wl, groupMemberPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to update group_member row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to get rows affected by update for group_member")
	}

	if !cached {
		groupMemberUpdateCacheMut.Lock()
		groupMemberUpdateCache[key] = cache
		groupMemberUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q groupMemberQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to update all for group_member")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to retrieve rows affected for group_member")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GroupMemberSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("im: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `group_member` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupMemberPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to update all in groupMember slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to retrieve rows affected all in update all groupMember")
	}
	return rowsAff, nil
}

var mySQLGroupMemberUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GroupMember) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("im: no group_member provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(groupMemberColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGroupMemberUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	groupMemberUpsertCacheMut.RLock()
	cache, cached := groupMemberUpsertCache[key]
	groupMemberUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			groupMemberAllColumns,
			groupMemberColumnsWithDefault,
			groupMemberColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			groupMemberAllColumns,
			groupMemberPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("im: unable to upsert group_member, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "group_member", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `group_member` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(groupMemberType, groupMemberMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(groupMemberType, groupMemberMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "im: unable to upsert for group_member")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(groupMemberType, groupMemberMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "im: unable to retrieve unique values for group_member")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "im: unable to populate default values for group_member")
	}

CacheNoHooks:
	if !cached {
		groupMemberUpsertCacheMut.Lock()
		groupMemberUpsertCache[key] = cache
		groupMemberUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GroupMember record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GroupMember) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("im: no GroupMember provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), groupMemberPrimaryKeyMapping)
	sql := "DELETE FROM `group_member` WHERE `group_id`=? AND `uid`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to delete from group_member")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to get rows affected by delete for group_member")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q groupMemberQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("im: no groupMemberQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to delete all from group_member")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to get rows affected by deleteall for group_member")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GroupMemberSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(groupMemberBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `group_member` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupMemberPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to delete all from groupMember slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to get rows affected by deleteall for group_member")
	}

	if len(groupMemberAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GroupMember) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGroupMember(ctx, exec, o.GroupID, o.UID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GroupMemberSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GroupMemberSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `group_member`.* FROM `group_member` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupMemberPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "im: unable to reload all in GroupMemberSlice")
	}

	*o = slice

	return nil
}

// GroupMemberExists checks if the GroupMember row exists.
func GroupMemberExists(ctx context.Context, exec boil.ContextExecutor, groupID int64, uID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `group_member` where `group_id`=? AND `uid`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, groupID, uID)
	}

	row := exec.QueryRowContext(ctx, sql, groupID, uID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "im: unable to check if group_member exists")
	}

	return exists, nil
}
