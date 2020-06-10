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

// HomePoi is an object representing the database table.
type HomePoi struct {
	ID               uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	OwnerID          string    `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`
	Location         string    `boil:"location" json:"location" toml:"location" yaml:"location"`
	PoiId            string    `boil:"poiId" json:"poiId" toml:"poiId" yaml:"poiId"`
	Countryname      string    `boil:"countryname" json:"countryname" toml:"countryname" yaml:"countryname"`
	Pname            string    `boil:"pname" json:"pname" toml:"pname" yaml:"pname"`
	Cityname         string    `boil:"cityname" json:"cityname" toml:"cityname" yaml:"cityname"`
	Name             string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	FormattedAddress string    `boil:"formattedAddress" json:"formattedAddress" toml:"formattedAddress" yaml:"formattedAddress"`
	CreatedAt        time.Time `boil:"createdAt" json:"createdAt" toml:"createdAt" yaml:"createdAt"`

	R *homePoiR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L homePoiL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HomePoiColumns = struct {
	ID               string
	OwnerID          string
	Location         string
	PoiId            string
	Countryname      string
	Pname            string
	Cityname         string
	Name             string
	FormattedAddress string
	CreatedAt        string
}{
	ID:               "id",
	OwnerID:          "owner_id",
	Location:         "location",
	PoiId:            "poiId",
	Countryname:      "countryname",
	Pname:            "pname",
	Cityname:         "cityname",
	Name:             "name",
	FormattedAddress: "formattedAddress",
	CreatedAt:        "createdAt",
}

// Generated where

var HomePoiWhere = struct {
	ID               whereHelperuint
	OwnerID          whereHelperstring
	Location         whereHelperstring
	PoiId            whereHelperstring
	Countryname      whereHelperstring
	Pname            whereHelperstring
	Cityname         whereHelperstring
	Name             whereHelperstring
	FormattedAddress whereHelperstring
	CreatedAt        whereHelpertime_Time
}{
	ID:               whereHelperuint{field: "`home_publish_msg_poi`.`id`"},
	OwnerID:          whereHelperstring{field: "`home_publish_msg_poi`.`owner_id`"},
	Location:         whereHelperstring{field: "`home_publish_msg_poi`.`location`"},
	PoiId:            whereHelperstring{field: "`home_publish_msg_poi`.`poiId`"},
	Countryname:      whereHelperstring{field: "`home_publish_msg_poi`.`countryname`"},
	Pname:            whereHelperstring{field: "`home_publish_msg_poi`.`pname`"},
	Cityname:         whereHelperstring{field: "`home_publish_msg_poi`.`cityname`"},
	Name:             whereHelperstring{field: "`home_publish_msg_poi`.`name`"},
	FormattedAddress: whereHelperstring{field: "`home_publish_msg_poi`.`formattedAddress`"},
	CreatedAt:        whereHelpertime_Time{field: "`home_publish_msg_poi`.`createdAt`"},
}

// HomePoiRels is where relationship names are stored.
var HomePoiRels = struct {
}{}

// homePoiR is where relationships are stored.
type homePoiR struct {
}

// NewStruct creates a new relationship struct
func (*homePoiR) NewStruct() *homePoiR {
	return &homePoiR{}
}

// homePoiL is where Load methods for each relationship are stored.
type homePoiL struct{}

var (
	homePoiAllColumns            = []string{"id", "owner_id", "location", "poiId", "countryname", "pname", "cityname", "name", "formattedAddress", "createdAt"}
	homePoiColumnsWithoutDefault = []string{"owner_id", "location", "poiId", "countryname", "pname", "cityname", "name", "formattedAddress"}
	homePoiColumnsWithDefault    = []string{"id", "createdAt"}
	homePoiPrimaryKeyColumns     = []string{"id"}
)

type (
	// HomePoiSlice is an alias for a slice of pointers to HomePoi.
	// This should generally be used opposed to []HomePoi.
	HomePoiSlice []*HomePoi
	// HomePoiHook is the signature for custom HomePoi hook methods
	HomePoiHook func(context.Context, boil.ContextExecutor, *HomePoi) error

	homePoiQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	homePoiType                 = reflect.TypeOf(&HomePoi{})
	homePoiMapping              = queries.MakeStructMapping(homePoiType)
	homePoiPrimaryKeyMapping, _ = queries.BindMapping(homePoiType, homePoiMapping, homePoiPrimaryKeyColumns)
	homePoiInsertCacheMut       sync.RWMutex
	homePoiInsertCache          = make(map[string]insertCache)
	homePoiUpdateCacheMut       sync.RWMutex
	homePoiUpdateCache          = make(map[string]updateCache)
	homePoiUpsertCacheMut       sync.RWMutex
	homePoiUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var homePoiBeforeInsertHooks []HomePoiHook
var homePoiBeforeUpdateHooks []HomePoiHook
var homePoiBeforeDeleteHooks []HomePoiHook
var homePoiBeforeUpsertHooks []HomePoiHook

var homePoiAfterInsertHooks []HomePoiHook
var homePoiAfterSelectHooks []HomePoiHook
var homePoiAfterUpdateHooks []HomePoiHook
var homePoiAfterDeleteHooks []HomePoiHook
var homePoiAfterUpsertHooks []HomePoiHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HomePoi) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HomePoi) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HomePoi) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HomePoi) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HomePoi) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HomePoi) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HomePoi) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HomePoi) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HomePoi) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homePoiAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHomePoiHook registers your hook function for all future operations.
func AddHomePoiHook(hookPoint boil.HookPoint, homePoiHook HomePoiHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		homePoiBeforeInsertHooks = append(homePoiBeforeInsertHooks, homePoiHook)
	case boil.BeforeUpdateHook:
		homePoiBeforeUpdateHooks = append(homePoiBeforeUpdateHooks, homePoiHook)
	case boil.BeforeDeleteHook:
		homePoiBeforeDeleteHooks = append(homePoiBeforeDeleteHooks, homePoiHook)
	case boil.BeforeUpsertHook:
		homePoiBeforeUpsertHooks = append(homePoiBeforeUpsertHooks, homePoiHook)
	case boil.AfterInsertHook:
		homePoiAfterInsertHooks = append(homePoiAfterInsertHooks, homePoiHook)
	case boil.AfterSelectHook:
		homePoiAfterSelectHooks = append(homePoiAfterSelectHooks, homePoiHook)
	case boil.AfterUpdateHook:
		homePoiAfterUpdateHooks = append(homePoiAfterUpdateHooks, homePoiHook)
	case boil.AfterDeleteHook:
		homePoiAfterDeleteHooks = append(homePoiAfterDeleteHooks, homePoiHook)
	case boil.AfterUpsertHook:
		homePoiAfterUpsertHooks = append(homePoiAfterUpsertHooks, homePoiHook)
	}
}

// One returns a single homePoi record from the query.
func (q homePoiQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HomePoi, error) {
	o := &HomePoi{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "home: failed to execute a one query for home_publish_msg_poi")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all HomePoi records from the query.
func (q homePoiQuery) All(ctx context.Context, exec boil.ContextExecutor) (HomePoiSlice, error) {
	var o []*HomePoi

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "home: failed to assign all query results to HomePoi slice")
	}

	if len(homePoiAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all HomePoi records in the query.
func (q homePoiQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to count home_publish_msg_poi rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q homePoiQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "home: failed to check if home_publish_msg_poi exists")
	}

	return count > 0, nil
}

// HomePois retrieves all the records using an executor.
func HomePois(mods ...qm.QueryMod) homePoiQuery {
	mods = append(mods, qm.From("`home_publish_msg_poi`"))
	return homePoiQuery{NewQuery(mods...)}
}

// FindHomePoi retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHomePoi(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*HomePoi, error) {
	homePoiObj := &HomePoi{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `home_publish_msg_poi` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, homePoiObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "home: unable to select from home_publish_msg_poi")
	}

	return homePoiObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HomePoi) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("home: no home_publish_msg_poi provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(homePoiColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	homePoiInsertCacheMut.RLock()
	cache, cached := homePoiInsertCache[key]
	homePoiInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			homePoiAllColumns,
			homePoiColumnsWithDefault,
			homePoiColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(homePoiType, homePoiMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(homePoiType, homePoiMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `home_publish_msg_poi` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `home_publish_msg_poi` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `home_publish_msg_poi` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, homePoiPrimaryKeyColumns))
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
		return errors.Wrap(err, "home: unable to insert into home_publish_msg_poi")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == homePoiMapping["ID"] {
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
		return errors.Wrap(err, "home: unable to populate default values for home_publish_msg_poi")
	}

CacheNoHooks:
	if !cached {
		homePoiInsertCacheMut.Lock()
		homePoiInsertCache[key] = cache
		homePoiInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the HomePoi.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HomePoi) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	homePoiUpdateCacheMut.RLock()
	cache, cached := homePoiUpdateCache[key]
	homePoiUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			homePoiAllColumns,
			homePoiPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("home: unable to update home_publish_msg_poi, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `home_publish_msg_poi` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, homePoiPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(homePoiType, homePoiMapping, append(wl, homePoiPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "home: unable to update home_publish_msg_poi row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by update for home_publish_msg_poi")
	}

	if !cached {
		homePoiUpdateCacheMut.Lock()
		homePoiUpdateCache[key] = cache
		homePoiUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q homePoiQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to update all for home_publish_msg_poi")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to retrieve rows affected for home_publish_msg_poi")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HomePoiSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homePoiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `home_publish_msg_poi` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homePoiPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to update all in homePoi slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to retrieve rows affected all in update all homePoi")
	}
	return rowsAff, nil
}

var mySQLHomePoiUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HomePoi) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("home: no home_publish_msg_poi provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(homePoiColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLHomePoiUniqueColumns, o)

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

	homePoiUpsertCacheMut.RLock()
	cache, cached := homePoiUpsertCache[key]
	homePoiUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			homePoiAllColumns,
			homePoiColumnsWithDefault,
			homePoiColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			homePoiAllColumns,
			homePoiPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("home: unable to upsert home_publish_msg_poi, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "home_publish_msg_poi", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `home_publish_msg_poi` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(homePoiType, homePoiMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(homePoiType, homePoiMapping, ret)
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
		return errors.Wrap(err, "home: unable to upsert for home_publish_msg_poi")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == homePoiMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(homePoiType, homePoiMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "home: unable to retrieve unique values for home_publish_msg_poi")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "home: unable to populate default values for home_publish_msg_poi")
	}

CacheNoHooks:
	if !cached {
		homePoiUpsertCacheMut.Lock()
		homePoiUpsertCache[key] = cache
		homePoiUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single HomePoi record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HomePoi) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("home: no HomePoi provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), homePoiPrimaryKeyMapping)
	sql := "DELETE FROM `home_publish_msg_poi` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete from home_publish_msg_poi")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by delete for home_publish_msg_poi")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q homePoiQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("home: no homePoiQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete all from home_publish_msg_poi")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by deleteall for home_publish_msg_poi")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HomePoiSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(homePoiBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homePoiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `home_publish_msg_poi` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homePoiPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete all from homePoi slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by deleteall for home_publish_msg_poi")
	}

	if len(homePoiAfterDeleteHooks) != 0 {
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
func (o *HomePoi) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHomePoi(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HomePoiSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HomePoiSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homePoiPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `home_publish_msg_poi`.* FROM `home_publish_msg_poi` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homePoiPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "home: unable to reload all in HomePoiSlice")
	}

	*o = slice

	return nil
}

// HomePoiExists checks if the HomePoi row exists.
func HomePoiExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `home_publish_msg_poi` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "home: unable to check if home_publish_msg_poi exists")
	}

	return exists, nil
}
