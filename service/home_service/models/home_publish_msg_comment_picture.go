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

// HomeCommentPicture is an object representing the database table.
type HomeCommentPicture struct {
	ID           uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	OwnerID      string    `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`
	ThumbnailUrl string    `boil:"thumbnailUrl" json:"thumbnailUrl" toml:"thumbnailUrl" yaml:"thumbnailUrl"`
	MiddlePicUrl string    `boil:"middlePicUrl" json:"middlePicUrl" toml:"middlePicUrl" yaml:"middlePicUrl"`
	PicUrl       string    `boil:"picUrl" json:"picUrl" toml:"picUrl" yaml:"picUrl"`
	Format       string    `boil:"format" json:"format" toml:"format" yaml:"format"`
	Width        int       `boil:"width" json:"width" toml:"width" yaml:"width"`
	Height       int       `boil:"height" json:"height" toml:"height" yaml:"height"`
	Type         int8      `boil:"type" json:"type" toml:"type" yaml:"type"`
	CreatedAt    time.Time `boil:"createdAt" json:"createdAt" toml:"createdAt" yaml:"createdAt"`

	R *homeCommentPictureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L homeCommentPictureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HomeCommentPictureColumns = struct {
	ID           string
	OwnerID      string
	ThumbnailUrl string
	MiddlePicUrl string
	PicUrl       string
	Format       string
	Width        string
	Height       string
	Type         string
	CreatedAt    string
}{
	ID:           "id",
	OwnerID:      "owner_id",
	ThumbnailUrl: "thumbnailUrl",
	MiddlePicUrl: "middlePicUrl",
	PicUrl:       "picUrl",
	Format:       "format",
	Width:        "width",
	Height:       "height",
	Type:         "type",
	CreatedAt:    "createdAt",
}

// Generated where

var HomeCommentPictureWhere = struct {
	ID           whereHelperuint
	OwnerID      whereHelperstring
	ThumbnailUrl whereHelperstring
	MiddlePicUrl whereHelperstring
	PicUrl       whereHelperstring
	Format       whereHelperstring
	Width        whereHelperint
	Height       whereHelperint
	Type         whereHelperint8
	CreatedAt    whereHelpertime_Time
}{
	ID:           whereHelperuint{field: "`home_publish_msg_comment_picture`.`id`"},
	OwnerID:      whereHelperstring{field: "`home_publish_msg_comment_picture`.`owner_id`"},
	ThumbnailUrl: whereHelperstring{field: "`home_publish_msg_comment_picture`.`thumbnailUrl`"},
	MiddlePicUrl: whereHelperstring{field: "`home_publish_msg_comment_picture`.`middlePicUrl`"},
	PicUrl:       whereHelperstring{field: "`home_publish_msg_comment_picture`.`picUrl`"},
	Format:       whereHelperstring{field: "`home_publish_msg_comment_picture`.`format`"},
	Width:        whereHelperint{field: "`home_publish_msg_comment_picture`.`width`"},
	Height:       whereHelperint{field: "`home_publish_msg_comment_picture`.`height`"},
	Type:         whereHelperint8{field: "`home_publish_msg_comment_picture`.`type`"},
	CreatedAt:    whereHelpertime_Time{field: "`home_publish_msg_comment_picture`.`createdAt`"},
}

// HomeCommentPictureRels is where relationship names are stored.
var HomeCommentPictureRels = struct {
}{}

// homeCommentPictureR is where relationships are stored.
type homeCommentPictureR struct {
}

// NewStruct creates a new relationship struct
func (*homeCommentPictureR) NewStruct() *homeCommentPictureR {
	return &homeCommentPictureR{}
}

// homeCommentPictureL is where Load methods for each relationship are stored.
type homeCommentPictureL struct{}

var (
	homeCommentPictureAllColumns            = []string{"id", "owner_id", "thumbnailUrl", "middlePicUrl", "picUrl", "format", "width", "height", "type", "createdAt"}
	homeCommentPictureColumnsWithoutDefault = []string{"owner_id", "thumbnailUrl", "middlePicUrl", "picUrl", "format"}
	homeCommentPictureColumnsWithDefault    = []string{"id", "width", "height", "type", "createdAt"}
	homeCommentPicturePrimaryKeyColumns     = []string{"id"}
)

type (
	// HomeCommentPictureSlice is an alias for a slice of pointers to HomeCommentPicture.
	// This should generally be used opposed to []HomeCommentPicture.
	HomeCommentPictureSlice []*HomeCommentPicture
	// HomeCommentPictureHook is the signature for custom HomeCommentPicture hook methods
	HomeCommentPictureHook func(context.Context, boil.ContextExecutor, *HomeCommentPicture) error

	homeCommentPictureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	homeCommentPictureType                 = reflect.TypeOf(&HomeCommentPicture{})
	homeCommentPictureMapping              = queries.MakeStructMapping(homeCommentPictureType)
	homeCommentPicturePrimaryKeyMapping, _ = queries.BindMapping(homeCommentPictureType, homeCommentPictureMapping, homeCommentPicturePrimaryKeyColumns)
	homeCommentPictureInsertCacheMut       sync.RWMutex
	homeCommentPictureInsertCache          = make(map[string]insertCache)
	homeCommentPictureUpdateCacheMut       sync.RWMutex
	homeCommentPictureUpdateCache          = make(map[string]updateCache)
	homeCommentPictureUpsertCacheMut       sync.RWMutex
	homeCommentPictureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var homeCommentPictureBeforeInsertHooks []HomeCommentPictureHook
var homeCommentPictureBeforeUpdateHooks []HomeCommentPictureHook
var homeCommentPictureBeforeDeleteHooks []HomeCommentPictureHook
var homeCommentPictureBeforeUpsertHooks []HomeCommentPictureHook

var homeCommentPictureAfterInsertHooks []HomeCommentPictureHook
var homeCommentPictureAfterSelectHooks []HomeCommentPictureHook
var homeCommentPictureAfterUpdateHooks []HomeCommentPictureHook
var homeCommentPictureAfterDeleteHooks []HomeCommentPictureHook
var homeCommentPictureAfterUpsertHooks []HomeCommentPictureHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HomeCommentPicture) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HomeCommentPicture) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HomeCommentPicture) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HomeCommentPicture) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HomeCommentPicture) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HomeCommentPicture) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HomeCommentPicture) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HomeCommentPicture) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HomeCommentPicture) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range homeCommentPictureAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHomeCommentPictureHook registers your hook function for all future operations.
func AddHomeCommentPictureHook(hookPoint boil.HookPoint, homeCommentPictureHook HomeCommentPictureHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		homeCommentPictureBeforeInsertHooks = append(homeCommentPictureBeforeInsertHooks, homeCommentPictureHook)
	case boil.BeforeUpdateHook:
		homeCommentPictureBeforeUpdateHooks = append(homeCommentPictureBeforeUpdateHooks, homeCommentPictureHook)
	case boil.BeforeDeleteHook:
		homeCommentPictureBeforeDeleteHooks = append(homeCommentPictureBeforeDeleteHooks, homeCommentPictureHook)
	case boil.BeforeUpsertHook:
		homeCommentPictureBeforeUpsertHooks = append(homeCommentPictureBeforeUpsertHooks, homeCommentPictureHook)
	case boil.AfterInsertHook:
		homeCommentPictureAfterInsertHooks = append(homeCommentPictureAfterInsertHooks, homeCommentPictureHook)
	case boil.AfterSelectHook:
		homeCommentPictureAfterSelectHooks = append(homeCommentPictureAfterSelectHooks, homeCommentPictureHook)
	case boil.AfterUpdateHook:
		homeCommentPictureAfterUpdateHooks = append(homeCommentPictureAfterUpdateHooks, homeCommentPictureHook)
	case boil.AfterDeleteHook:
		homeCommentPictureAfterDeleteHooks = append(homeCommentPictureAfterDeleteHooks, homeCommentPictureHook)
	case boil.AfterUpsertHook:
		homeCommentPictureAfterUpsertHooks = append(homeCommentPictureAfterUpsertHooks, homeCommentPictureHook)
	}
}

// One returns a single homeCommentPicture record from the query.
func (q homeCommentPictureQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HomeCommentPicture, error) {
	o := &HomeCommentPicture{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "home: failed to execute a one query for home_publish_msg_comment_picture")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all HomeCommentPicture records from the query.
func (q homeCommentPictureQuery) All(ctx context.Context, exec boil.ContextExecutor) (HomeCommentPictureSlice, error) {
	var o []*HomeCommentPicture

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "home: failed to assign all query results to HomeCommentPicture slice")
	}

	if len(homeCommentPictureAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all HomeCommentPicture records in the query.
func (q homeCommentPictureQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to count home_publish_msg_comment_picture rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q homeCommentPictureQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "home: failed to check if home_publish_msg_comment_picture exists")
	}

	return count > 0, nil
}

// HomeCommentPictures retrieves all the records using an executor.
func HomeCommentPictures(mods ...qm.QueryMod) homeCommentPictureQuery {
	mods = append(mods, qm.From("`home_publish_msg_comment_picture`"))
	return homeCommentPictureQuery{NewQuery(mods...)}
}

// FindHomeCommentPicture retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHomeCommentPicture(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*HomeCommentPicture, error) {
	homeCommentPictureObj := &HomeCommentPicture{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `home_publish_msg_comment_picture` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, homeCommentPictureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "home: unable to select from home_publish_msg_comment_picture")
	}

	return homeCommentPictureObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HomeCommentPicture) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("home: no home_publish_msg_comment_picture provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(homeCommentPictureColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	homeCommentPictureInsertCacheMut.RLock()
	cache, cached := homeCommentPictureInsertCache[key]
	homeCommentPictureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			homeCommentPictureAllColumns,
			homeCommentPictureColumnsWithDefault,
			homeCommentPictureColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(homeCommentPictureType, homeCommentPictureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(homeCommentPictureType, homeCommentPictureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `home_publish_msg_comment_picture` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `home_publish_msg_comment_picture` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `home_publish_msg_comment_picture` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, homeCommentPicturePrimaryKeyColumns))
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
		return errors.Wrap(err, "home: unable to insert into home_publish_msg_comment_picture")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == homeCommentPictureMapping["ID"] {
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
		return errors.Wrap(err, "home: unable to populate default values for home_publish_msg_comment_picture")
	}

CacheNoHooks:
	if !cached {
		homeCommentPictureInsertCacheMut.Lock()
		homeCommentPictureInsertCache[key] = cache
		homeCommentPictureInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the HomeCommentPicture.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HomeCommentPicture) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	homeCommentPictureUpdateCacheMut.RLock()
	cache, cached := homeCommentPictureUpdateCache[key]
	homeCommentPictureUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			homeCommentPictureAllColumns,
			homeCommentPicturePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("home: unable to update home_publish_msg_comment_picture, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `home_publish_msg_comment_picture` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, homeCommentPicturePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(homeCommentPictureType, homeCommentPictureMapping, append(wl, homeCommentPicturePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "home: unable to update home_publish_msg_comment_picture row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by update for home_publish_msg_comment_picture")
	}

	if !cached {
		homeCommentPictureUpdateCacheMut.Lock()
		homeCommentPictureUpdateCache[key] = cache
		homeCommentPictureUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q homeCommentPictureQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to update all for home_publish_msg_comment_picture")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to retrieve rows affected for home_publish_msg_comment_picture")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HomeCommentPictureSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeCommentPicturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `home_publish_msg_comment_picture` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homeCommentPicturePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to update all in homeCommentPicture slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to retrieve rows affected all in update all homeCommentPicture")
	}
	return rowsAff, nil
}

var mySQLHomeCommentPictureUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HomeCommentPicture) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("home: no home_publish_msg_comment_picture provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(homeCommentPictureColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLHomeCommentPictureUniqueColumns, o)

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

	homeCommentPictureUpsertCacheMut.RLock()
	cache, cached := homeCommentPictureUpsertCache[key]
	homeCommentPictureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			homeCommentPictureAllColumns,
			homeCommentPictureColumnsWithDefault,
			homeCommentPictureColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			homeCommentPictureAllColumns,
			homeCommentPicturePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("home: unable to upsert home_publish_msg_comment_picture, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "home_publish_msg_comment_picture", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `home_publish_msg_comment_picture` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(homeCommentPictureType, homeCommentPictureMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(homeCommentPictureType, homeCommentPictureMapping, ret)
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
		return errors.Wrap(err, "home: unable to upsert for home_publish_msg_comment_picture")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == homeCommentPictureMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(homeCommentPictureType, homeCommentPictureMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "home: unable to retrieve unique values for home_publish_msg_comment_picture")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "home: unable to populate default values for home_publish_msg_comment_picture")
	}

CacheNoHooks:
	if !cached {
		homeCommentPictureUpsertCacheMut.Lock()
		homeCommentPictureUpsertCache[key] = cache
		homeCommentPictureUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single HomeCommentPicture record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HomeCommentPicture) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("home: no HomeCommentPicture provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), homeCommentPicturePrimaryKeyMapping)
	sql := "DELETE FROM `home_publish_msg_comment_picture` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete from home_publish_msg_comment_picture")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by delete for home_publish_msg_comment_picture")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q homeCommentPictureQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("home: no homeCommentPictureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete all from home_publish_msg_comment_picture")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by deleteall for home_publish_msg_comment_picture")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HomeCommentPictureSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(homeCommentPictureBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeCommentPicturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `home_publish_msg_comment_picture` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homeCommentPicturePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "home: unable to delete all from homeCommentPicture slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "home: failed to get rows affected by deleteall for home_publish_msg_comment_picture")
	}

	if len(homeCommentPictureAfterDeleteHooks) != 0 {
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
func (o *HomeCommentPicture) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHomeCommentPicture(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HomeCommentPictureSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HomeCommentPictureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), homeCommentPicturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `home_publish_msg_comment_picture`.* FROM `home_publish_msg_comment_picture` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, homeCommentPicturePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "home: unable to reload all in HomeCommentPictureSlice")
	}

	*o = slice

	return nil
}

// HomeCommentPictureExists checks if the HomeCommentPicture row exists.
func HomeCommentPictureExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `home_publish_msg_comment_picture` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "home: unable to check if home_publish_msg_comment_picture exists")
	}

	return exists, nil
}