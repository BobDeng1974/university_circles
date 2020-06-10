// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package home

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

// HomeLinkinfo is an object representing the database table.
type HomeLinkinfo struct {
	ID         uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	MSGID      string    `boil:"msg_id" json:"msg_id" toml:"msg_id" yaml:"msg_id"`
	Title      string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	PictureUrl string    `boil:"pictureUrl" json:"pictureUrl" toml:"pictureUrl" yaml:"pictureUrl"`
	LinkUrl    string    `boil:"linkUrl" json:"linkUrl" toml:"linkUrl" yaml:"linkUrl"`
	Source     string    `boil:"source" json:"source" toml:"source" yaml:"source"`
	VideoID    int       `boil:"video_id" json:"video_id" toml:"video_id" yaml:"video_id"`
	CreatedAt  time.Time `boil:"createdAt" json:"createdAt" toml:"createdAt" yaml:"createdAt"`

	R *homeLinkinfoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L homeLinkinfoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HomeLinkinfoColumns = struct {
	ID         string
	MSGID      string
	Title      string
	PictureUrl string
	LinkUrl    string
	Source     string
	VideoID    string
	CreatedAt  string
}{
	ID:         "id",
	MSGID:      "msg_id",
	Title:      "title",
	PictureUrl: "pictureUrl",
	LinkUrl:    "linkUrl",
	Source:     "source",
	VideoID:    "video_id",
	CreatedAt:  "createdAt",
}

// Generated where

var HomeLinkinfoWhere = struct {
	ID         whereHelperuint
	MSGID      whereHelperstring
	Title      whereHelperstring
	PictureUrl whereHelperstring
	LinkUrl    whereHelperstring
	Source     whereHelperstring
	VideoID    whereHelperint
	CreatedAt  whereHelpertime_Time
}{
	ID:         whereHelperuint{field: "`home_publish_msg_linkinfo`.`id`"},
	MSGID:      whereHelperstring{field: "`home_publish_msg_linkinfo`.`msg_id`"},
	Title:      whereHelperstring{field: "`home_publish_msg_linkinfo`.`title`"},
	PictureUrl: whereHelperstring{field: "`home_publish_msg_linkinfo`.`pictureUrl`"},
	LinkUrl:    whereHelperstring{field: "`home_publish_msg_linkinfo`.`linkUrl`"},
	Source:     whereHelperstring{field: "`home_publish_msg_linkinfo`.`source`"},
	VideoID:    whereHelperint{field: "`home_publish_msg_linkinfo`.`video_id`"},
	CreatedAt:  whereHelpertime_Time{field: "`home_publish_msg_linkinfo`.`createdAt`"},
}

// HomeLinkinfoRels is where relationship names are stored.
var HomeLinkinfoRels = struct {
}{}

// homeLinkinfoR is where relationships are stored.
type homeLinkinfoR struct {
}

// NewStruct creates a new relationship struct
func (*homeLinkinfoR) NewStruct() *homeLinkinfoR {
	return &homeLinkinfoR{}
}

// homeLinkinfoL is where Load methods for each relationship are stored.
type homeLinkinfoL struct{}

var (
	homeLinkinfoAllColumns            = []string{"id", "msg_id", "title", "pictureUrl", "linkUrl", "source", "video_id", "createdAt"}
	homeLinkinfoColumnsWithoutDefault = []string{"msg_id", "title", "pictureUrl", "linkUrl", "source"}
	homeLinkinfoColumnsWithDefault    = []string{"id", "video_id", "createdAt"}
	homeLinkinfoPrimaryKeyColumns     = []string{"id"}
)

type (
	// HomeLinkinfoSlice is an alias for a slice of pointers to HomeLinkinfo.
	// This should generally be used opposed to []HomeLinkinfo.
	HomeLinkinfoSlice []*HomeLinkinfo
	// HomeLinkinfoHook is the signature for custom HomeLinkinfo hook methods
	HomeLinkinfoHook func(context.Context, boil.ContextExecutor, *HomeLinkinfo) error

	homeLinkinfoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	homeLinkinfoType                 = reflect.TypeOf(&HomeLinkinfo{})
	homeLinkinfoMapping              = queries.MakeStructMapping(homeLinkinfoType)
	homeLinkinfoPrimaryKeyMapping, _ = queries.BindMapping(homeLinkinfoType, homeLinkinfoMapping, homeLinkinfoPrimaryKeyColumns)
	homeLinkinfoInsertCacheMut       sync.RWMutex
	homeLinkinfoInsertCache          = make(map[string]insertCache)
	homeLinkinfoUpdateCacheMut       sync.RWMutex
	homeLinkinfoUpdateCache          = make(map[string]updateCache)
	homeLinkinfoUpsertCacheMut       sync.RWMutex
	homeLinkinfoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var homeLinkinfoBeforeInsertHooks []HomeLinkinfoHook
var homeLinkinfoBeforeUpdateHooks []HomeLinkinfoHook
var homeLinkinfoBeforeDeleteHooks []HomeLinkinfoHook
var homeLinkinfoBeforeUpsertHooks []HomeLinkinfoHook

var homeLinkinfoAfterInsertHooks []HomeLinkinfoHook
var homeLinkinfoAfterSelectHooks []HomeLinkinfoHook
var homeLinkinfoAfterUpdateHooks []HomeLinkinfoHook
var homeLinkinfoAfterDeleteHooks []HomeLinkinfoHook
var homeLinkinfoAfterUpsertHooks []HomeLinkinfoHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HomeLinkinfo) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HomeLinkinfo) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HomeLinkinfo) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HomeLinkinfo) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HomeLinkinfo) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HomeLinkinfo) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HomeLinkinfo) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HomeLinkinfo) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HomeLinkinfo) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeLinkinfoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHomeLinkinfoHook registers your hook function for all future operations.
func AddHomeLinkinfoHook(hookPoint boil.HookPoint, homeLinkinfoHook HomeLinkinfoHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		homeLinkinfoBeforeInsertHooks = append(homeLinkinfoBeforeInsertHooks, homeLinkinfoHook)
	case boil.BeforeUpdateHook:
		homeLinkinfoBeforeUpdateHooks = append(homeLinkinfoBeforeUpdateHooks, homeLinkinfoHook)
	case boil.BeforeDeleteHook:
		homeLinkinfoBeforeDeleteHooks = append(homeLinkinfoBeforeDeleteHooks, homeLinkinfoHook)
	case boil.BeforeUpsertHook:
		homeLinkinfoBeforeUpsertHooks = append(homeLinkinfoBeforeUpsertHooks, homeLinkinfoHook)
	case boil.AfterInsertHook:
		homeLinkinfoAfterInsertHooks = append(homeLinkinfoAfterInsertHooks, homeLinkinfoHook)
	case boil.AfterSelectHook:
		homeLinkinfoAfterSelectHooks = append(homeLinkinfoAfterSelectHooks, homeLinkinfoHook)
	case boil.AfterUpdateHook:
		homeLinkinfoAfterUpdateHooks = append(homeLinkinfoAfterUpdateHooks, homeLinkinfoHook)
	case boil.AfterDeleteHook:
		homeLinkinfoAfterDeleteHooks = append(homeLinkinfoAfterDeleteHooks, homeLinkinfoHook)
	case boil.AfterUpsertHook:
		homeLinkinfoAfterUpsertHooks = append(homeLinkinfoAfterUpsertHooks, homeLinkinfoHook)
	}
}

// One returns a single homeLinkinfo record from the query.
func (q homeLinkinfoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HomeLinkinfo, error) {
	o := &HomeLinkinfo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "home: failed to execute a one query for home_publish_msg_linkinfo")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all HomeLinkinfo records from the query.
func (q homeLinkinfoQuery) All(ctx context.Context, exec boil.ContextExecutor) (HomeLinkinfoSlice, error) {
	var o []*HomeLinkinfo

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "home: failed to assign all query results to HomeLinkinfo slice")
	}

	if len(homeLinkinfoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all HomeLinkinfo records in the query.
func (q homeLinkinfoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to count home_publish_msg_linkinfo rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q homeLinkinfoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "home: failed to check if home_publish_msg_linkinfo exists")
	}

	return count > 0, nil
}

// HomeLinkinfos retrieves all the records using an executor.
func HomeLinkinfos(mods ...qm.QueryMod) homeLinkinfoQuery {
	mods = append(mods, qm.From("`home_publish_msg_linkinfo`"))
	return homeLinkinfoQuery{NewQuery(mods...)}
}

// FindHomeLinkinfo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHomeLinkinfo(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*HomeLinkinfo, error) {
	homeLinkinfoObj := &HomeLinkinfo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `home_publish_msg_linkinfo` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, homeLinkinfoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "home: unable to select from home_publish_msg_linkinfo")
	}

	return homeLinkinfoObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HomeLinkinfo) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("home: no home_publish_msg_linkinfo provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(homeLinkinfoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	homeLinkinfoInsertCacheMut.RLock()
	cache, cached := homeLinkinfoInsertCache[key]
	homeLinkinfoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			homeLinkinfoAllColumns,
			homeLinkinfoColumnsWithDefault,
			homeLinkinfoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(homeLinkinfoType, homeLinkinfoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(homeLinkinfoType, homeLinkinfoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `home_publish_msg_linkinfo` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `home_publish_msg_linkinfo` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `home_publish_msg_linkinfo` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, homeLinkinfoPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "home: unable to insert into home_publish_msg_linkinfo")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == homeLinkinfoMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "home: unable to populate default values for home_publish_msg_linkinfo")
	}

CacheNoHooks:
	if !cached {
		homeLinkinfoInsertCacheMut.Lock()
		homeLinkinfoInsertCache[key] = cache
		homeLinkinfoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the HomeLinkinfo.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HomeLinkinfo) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	homeLinkinfoUpdateCacheMut.RLock()
	cache, cached := homeLinkinfoUpdateCache[key]
	homeLinkinfoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			homeLinkinfoAllColumns,
			homeLinkinfoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("home: unable to update home_publish_msg_linkinfo, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `home_publish_msg_linkinfo` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, homeLinkinfoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(homeLinkinfoType, homeLinkinfoMapping, append(wl, homeLinkinfoPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "home: unable to update home_publish_msg_linkinfo row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by update for home_publish_msg_linkinfo")
	}

	if !cached {
		homeLinkinfoUpdateCacheMut.Lock()
		homeLinkinfoUpdateCache[key] = cache
		homeLinkinfoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q homeLinkinfoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to update all for home_publish_msg_linkinfo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to retrieve rows affected for home_publish_msg_linkinfo")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HomeLinkinfoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("home: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeLinkinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `home_publish_msg_linkinfo` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homeLinkinfoPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to update all in homeLinkinfo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to retrieve rows affected all in update all homeLinkinfo")
	}
	return rowsAff, nil
}

var mySQLHomeLinkinfoUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HomeLinkinfo) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("home: no home_publish_msg_linkinfo provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(homeLinkinfoColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLHomeLinkinfoUniqueColumns, o)

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

	homeLinkinfoUpsertCacheMut.RLock()
	cache, cached := homeLinkinfoUpsertCache[key]
	homeLinkinfoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			homeLinkinfoAllColumns,
			homeLinkinfoColumnsWithDefault,
			homeLinkinfoColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			homeLinkinfoAllColumns,
			homeLinkinfoPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("home: unable to upsert home_publish_msg_linkinfo, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "home_publish_msg_linkinfo", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `home_publish_msg_linkinfo` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(homeLinkinfoType, homeLinkinfoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(homeLinkinfoType, homeLinkinfoMapping, ret)
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

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "home: unable to upsert for home_publish_msg_linkinfo")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == homeLinkinfoMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(homeLinkinfoType, homeLinkinfoMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "home: unable to retrieve unique values for home_publish_msg_linkinfo")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "home: unable to populate default values for home_publish_msg_linkinfo")
	}

CacheNoHooks:
	if !cached {
		homeLinkinfoUpsertCacheMut.Lock()
		homeLinkinfoUpsertCache[key] = cache
		homeLinkinfoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single HomeLinkinfo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HomeLinkinfo) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("home: no HomeLinkinfo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), homeLinkinfoPrimaryKeyMapping)
	sql := "DELETE FROM `home_publish_msg_linkinfo` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete from home_publish_msg_linkinfo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by delete for home_publish_msg_linkinfo")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q homeLinkinfoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("home: no homeLinkinfoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete all from home_publish_msg_linkinfo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by deleteall for home_publish_msg_linkinfo")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HomeLinkinfoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(homeLinkinfoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeLinkinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `home_publish_msg_linkinfo` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homeLinkinfoPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete all from homeLinkinfo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by deleteall for home_publish_msg_linkinfo")
	}

	if len(homeLinkinfoAfterDeleteHooks) != 0 {
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
func (o *HomeLinkinfo) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHomeLinkinfo(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HomeLinkinfoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HomeLinkinfoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeLinkinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `home_publish_msg_linkinfo`.* FROM `home_publish_msg_linkinfo` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homeLinkinfoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "home: unable to reload all in HomeLinkinfoSlice")
	}

	*o = slice

	return nil
}

// HomeLinkinfoExists checks if the HomeLinkinfo row exists.
func HomeLinkinfoExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `home_publish_msg_linkinfo` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "home: unable to check if home_publish_msg_linkinfo exists")
	}

	return exists, nil
}