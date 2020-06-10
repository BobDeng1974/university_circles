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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// UTeacher is an object representing the database table.
type UTeacher struct {
	ID            uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	TeachNo       string      `boil:"teach_no" json:"teach_no" toml:"teach_no" yaml:"teach_no"`
	Password      string      `boil:"password" json:"password" toml:"password" yaml:"password"`
	RealName      string      `boil:"real_name" json:"real_name" toml:"real_name" yaml:"real_name"`
	ScreenName    string      `boil:"screen_name" json:"screen_name" toml:"screen_name" yaml:"screen_name"`
	Gender        string      `boil:"gender" json:"gender" toml:"gender" yaml:"gender"`
	Phone         string      `boil:"phone" json:"phone" toml:"phone" yaml:"phone"`
	Email         string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	Avatar        string      `boil:"avatar" json:"avatar" toml:"avatar" yaml:"avatar"`
	Birthday      string      `boil:"birthday" json:"birthday" toml:"birthday" yaml:"birthday"`
	UniversityID  int         `boil:"university_id" json:"university_id" toml:"university_id" yaml:"university_id"`
	CollegeID     int         `boil:"college_id" json:"college_id" toml:"college_id" yaml:"college_id"`
	ProfessionID  int         `boil:"profession_id" json:"profession_id" toml:"profession_id" yaml:"profession_id"`
	Bio           string      `boil:"bio" json:"bio" toml:"bio" yaml:"bio"`
	Deleted       uint8       `boil:"deleted" json:"deleted" toml:"deleted" yaml:"deleted"`
	CreatedUser   null.String `boil:"created_user" json:"created_user,omitempty" toml:"created_user" yaml:"created_user,omitempty"`
	CreatedAt     time.Time   `boil:"createdAt" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	ModifiedUser  null.String `boil:"modified_user" json:"modified_user,omitempty" toml:"modified_user" yaml:"modified_user,omitempty"`
	UpdatedAt     time.Time   `boil:"updatedAt" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	IsVerified    int8        `boil:"isVerified" json:"isVerified" toml:"isVerified" yaml:"isVerified"`
	VerifyMessage string      `boil:"verifyMessage" json:"verifyMessage" toml:"verifyMessage" yaml:"verifyMessage"`
	Mid           int64       `boil:"mid" json:"mid" toml:"mid" yaml:"mid"`
	UID           string      `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	IDCardNumber  string      `boil:"id_card_number" json:"id_card_number" toml:"id_card_number" yaml:"id_card_number"`
	Address       string      `boil:"address" json:"address" toml:"address" yaml:"address"`
	Zodiac        string      `boil:"zodiac" json:"zodiac" toml:"zodiac" yaml:"zodiac"`
	Detail        string      `boil:"detail" json:"detail" toml:"detail" yaml:"detail"`
	ForbiStranger int8        `boil:"forbi_stranger" json:"forbi_stranger" toml:"forbi_stranger" yaml:"forbi_stranger"`
	VerifyImage   string      `boil:"verify_image" json:"verify_image" toml:"verify_image" yaml:"verify_image"`

	R *uTeacherR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L uTeacherL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UTeacherColumns = struct {
	ID            string
	TeachNo       string
	Password      string
	RealName      string
	ScreenName    string
	Gender        string
	Phone         string
	Email         string
	Avatar        string
	Birthday      string
	UniversityID  string
	CollegeID     string
	ProfessionID  string
	Bio           string
	Deleted       string
	CreatedUser   string
	CreatedAt     string
	ModifiedUser  string
	UpdatedAt     string
	IsVerified    string
	VerifyMessage string
	Mid           string
	UID           string
	IDCardNumber  string
	Address       string
	Zodiac        string
	Detail        string
	ForbiStranger string
	VerifyImage   string
}{
	ID:            "id",
	TeachNo:       "teach_no",
	Password:      "password",
	RealName:      "real_name",
	ScreenName:    "screen_name",
	Gender:        "gender",
	Phone:         "phone",
	Email:         "email",
	Avatar:        "avatar",
	Birthday:      "birthday",
	UniversityID:  "university_id",
	CollegeID:     "college_id",
	ProfessionID:  "profession_id",
	Bio:           "bio",
	Deleted:       "deleted",
	CreatedUser:   "created_user",
	CreatedAt:     "createdAt",
	ModifiedUser:  "modified_user",
	UpdatedAt:     "updatedAt",
	IsVerified:    "isVerified",
	VerifyMessage: "verifyMessage",
	Mid:           "mid",
	UID:           "uid",
	IDCardNumber:  "id_card_number",
	Address:       "address",
	Zodiac:        "zodiac",
	Detail:        "detail",
	ForbiStranger: "forbi_stranger",
	VerifyImage:   "verify_image",
}

// Generated where

type whereHelperuint8 struct{ field string }

func (w whereHelperuint8) EQ(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint8) NEQ(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint8) LT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint8) LTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint8) GT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint8) GTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var UTeacherWhere = struct {
	ID            whereHelperuint
	TeachNo       whereHelperstring
	Password      whereHelperstring
	RealName      whereHelperstring
	ScreenName    whereHelperstring
	Gender        whereHelperstring
	Phone         whereHelperstring
	Email         whereHelperstring
	Avatar        whereHelperstring
	Birthday      whereHelperstring
	UniversityID  whereHelperint
	CollegeID     whereHelperint
	ProfessionID  whereHelperint
	Bio           whereHelperstring
	Deleted       whereHelperuint8
	CreatedUser   whereHelpernull_String
	CreatedAt     whereHelpertime_Time
	ModifiedUser  whereHelpernull_String
	UpdatedAt     whereHelpertime_Time
	IsVerified    whereHelperint8
	VerifyMessage whereHelperstring
	Mid           whereHelperint64
	UID           whereHelperstring
	IDCardNumber  whereHelperstring
	Address       whereHelperstring
	Zodiac        whereHelperstring
	Detail        whereHelperstring
	ForbiStranger whereHelperint8
	VerifyImage   whereHelperstring
}{
	ID:            whereHelperuint{field: "`university_teacher`.`id`"},
	TeachNo:       whereHelperstring{field: "`university_teacher`.`teach_no`"},
	Password:      whereHelperstring{field: "`university_teacher`.`password`"},
	RealName:      whereHelperstring{field: "`university_teacher`.`real_name`"},
	ScreenName:    whereHelperstring{field: "`university_teacher`.`screen_name`"},
	Gender:        whereHelperstring{field: "`university_teacher`.`gender`"},
	Phone:         whereHelperstring{field: "`university_teacher`.`phone`"},
	Email:         whereHelperstring{field: "`university_teacher`.`email`"},
	Avatar:        whereHelperstring{field: "`university_teacher`.`avatar`"},
	Birthday:      whereHelperstring{field: "`university_teacher`.`birthday`"},
	UniversityID:  whereHelperint{field: "`university_teacher`.`university_id`"},
	CollegeID:     whereHelperint{field: "`university_teacher`.`college_id`"},
	ProfessionID:  whereHelperint{field: "`university_teacher`.`profession_id`"},
	Bio:           whereHelperstring{field: "`university_teacher`.`bio`"},
	Deleted:       whereHelperuint8{field: "`university_teacher`.`deleted`"},
	CreatedUser:   whereHelpernull_String{field: "`university_teacher`.`created_user`"},
	CreatedAt:     whereHelpertime_Time{field: "`university_teacher`.`createdAt`"},
	ModifiedUser:  whereHelpernull_String{field: "`university_teacher`.`modified_user`"},
	UpdatedAt:     whereHelpertime_Time{field: "`university_teacher`.`updatedAt`"},
	IsVerified:    whereHelperint8{field: "`university_teacher`.`isVerified`"},
	VerifyMessage: whereHelperstring{field: "`university_teacher`.`verifyMessage`"},
	Mid:           whereHelperint64{field: "`university_teacher`.`mid`"},
	UID:           whereHelperstring{field: "`university_teacher`.`uid`"},
	IDCardNumber:  whereHelperstring{field: "`university_teacher`.`id_card_number`"},
	Address:       whereHelperstring{field: "`university_teacher`.`address`"},
	Zodiac:        whereHelperstring{field: "`university_teacher`.`zodiac`"},
	Detail:        whereHelperstring{field: "`university_teacher`.`detail`"},
	ForbiStranger: whereHelperint8{field: "`university_teacher`.`forbi_stranger`"},
	VerifyImage:   whereHelperstring{field: "`university_teacher`.`verify_image`"},
}

// UTeacherRels is where relationship names are stored.
var UTeacherRels = struct {
}{}

// uTeacherR is where relationships are stored.
type uTeacherR struct {
}

// NewStruct creates a new relationship struct
func (*uTeacherR) NewStruct() *uTeacherR {
	return &uTeacherR{}
}

// uTeacherL is where Load methods for each relationship are stored.
type uTeacherL struct{}

var (
	uTeacherAllColumns            = []string{"id", "teach_no", "password", "real_name", "screen_name", "gender", "phone", "email", "avatar", "birthday", "university_id", "college_id", "profession_id", "bio", "deleted", "created_user", "createdAt", "modified_user", "updatedAt", "isVerified", "verifyMessage", "mid", "uid", "id_card_number", "address", "zodiac", "detail", "forbi_stranger", "verify_image"}
	uTeacherColumnsWithoutDefault = []string{"teach_no", "password", "real_name", "screen_name", "phone", "email", "avatar", "birthday", "bio", "created_user", "modified_user", "isVerified", "verifyMessage", "mid", "uid", "id_card_number", "address", "zodiac", "detail", "forbi_stranger", "verify_image"}
	uTeacherColumnsWithDefault    = []string{"id", "gender", "university_id", "college_id", "profession_id", "deleted", "createdAt", "updatedAt"}
	uTeacherPrimaryKeyColumns     = []string{"id"}
)

type (
	// UTeacherSlice is an alias for a slice of pointers to UTeacher.
	// This should generally be used opposed to []UTeacher.
	UTeacherSlice []*UTeacher
	// UTeacherHook is the signature for custom UTeacher hook methods
	UTeacherHook func(context.Context, boil.ContextExecutor, *UTeacher) error

	uTeacherQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	uTeacherType                 = reflect.TypeOf(&UTeacher{})
	uTeacherMapping              = queries.MakeStructMapping(uTeacherType)
	uTeacherPrimaryKeyMapping, _ = queries.BindMapping(uTeacherType, uTeacherMapping, uTeacherPrimaryKeyColumns)
	uTeacherInsertCacheMut       sync.RWMutex
	uTeacherInsertCache          = make(map[string]insertCache)
	uTeacherUpdateCacheMut       sync.RWMutex
	uTeacherUpdateCache          = make(map[string]updateCache)
	uTeacherUpsertCacheMut       sync.RWMutex
	uTeacherUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var uTeacherBeforeInsertHooks []UTeacherHook
var uTeacherBeforeUpdateHooks []UTeacherHook
var uTeacherBeforeDeleteHooks []UTeacherHook
var uTeacherBeforeUpsertHooks []UTeacherHook

var uTeacherAfterInsertHooks []UTeacherHook
var uTeacherAfterSelectHooks []UTeacherHook
var uTeacherAfterUpdateHooks []UTeacherHook
var uTeacherAfterDeleteHooks []UTeacherHook
var uTeacherAfterUpsertHooks []UTeacherHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UTeacher) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UTeacher) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UTeacher) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UTeacher) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UTeacher) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UTeacher) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UTeacher) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UTeacher) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UTeacher) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range uTeacherAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUTeacherHook registers your hook function for all future operations.
func AddUTeacherHook(hookPoint boil.HookPoint, uTeacherHook UTeacherHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		uTeacherBeforeInsertHooks = append(uTeacherBeforeInsertHooks, uTeacherHook)
	case boil.BeforeUpdateHook:
		uTeacherBeforeUpdateHooks = append(uTeacherBeforeUpdateHooks, uTeacherHook)
	case boil.BeforeDeleteHook:
		uTeacherBeforeDeleteHooks = append(uTeacherBeforeDeleteHooks, uTeacherHook)
	case boil.BeforeUpsertHook:
		uTeacherBeforeUpsertHooks = append(uTeacherBeforeUpsertHooks, uTeacherHook)
	case boil.AfterInsertHook:
		uTeacherAfterInsertHooks = append(uTeacherAfterInsertHooks, uTeacherHook)
	case boil.AfterSelectHook:
		uTeacherAfterSelectHooks = append(uTeacherAfterSelectHooks, uTeacherHook)
	case boil.AfterUpdateHook:
		uTeacherAfterUpdateHooks = append(uTeacherAfterUpdateHooks, uTeacherHook)
	case boil.AfterDeleteHook:
		uTeacherAfterDeleteHooks = append(uTeacherAfterDeleteHooks, uTeacherHook)
	case boil.AfterUpsertHook:
		uTeacherAfterUpsertHooks = append(uTeacherAfterUpsertHooks, uTeacherHook)
	}
}

// One returns a single uTeacher record from the query.
func (q uTeacherQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UTeacher, error) {
	o := &UTeacher{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "im: failed to execute a one query for university_teacher")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UTeacher records from the query.
func (q uTeacherQuery) All(ctx context.Context, exec boil.ContextExecutor) (UTeacherSlice, error) {
	var o []*UTeacher

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "im: failed to assign all query results to UTeacher slice")
	}

	if len(uTeacherAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UTeacher records in the query.
func (q uTeacherQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to count university_teacher rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q uTeacherQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "im: failed to check if university_teacher exists")
	}

	return count > 0, nil
}

// UTeachers retrieves all the records using an executor.
func UTeachers(mods ...qm.QueryMod) uTeacherQuery {
	mods = append(mods, qm.From("`university_teacher`"))
	return uTeacherQuery{NewQuery(mods...)}
}

// FindUTeacher retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUTeacher(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*UTeacher, error) {
	uTeacherObj := &UTeacher{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `university_teacher` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, uTeacherObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "im: unable to select from university_teacher")
	}

	return uTeacherObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UTeacher) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("im: no university_teacher provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(uTeacherColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	uTeacherInsertCacheMut.RLock()
	cache, cached := uTeacherInsertCache[key]
	uTeacherInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			uTeacherAllColumns,
			uTeacherColumnsWithDefault,
			uTeacherColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(uTeacherType, uTeacherMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(uTeacherType, uTeacherMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `university_teacher` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `university_teacher` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `university_teacher` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, uTeacherPrimaryKeyColumns))
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
		return errors.Wrap(err, "im: unable to insert into university_teacher")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == uTeacherMapping["ID"] {
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
		return errors.Wrap(err, "im: unable to populate default values for university_teacher")
	}

CacheNoHooks:
	if !cached {
		uTeacherInsertCacheMut.Lock()
		uTeacherInsertCache[key] = cache
		uTeacherInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UTeacher.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UTeacher) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	uTeacherUpdateCacheMut.RLock()
	cache, cached := uTeacherUpdateCache[key]
	uTeacherUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			uTeacherAllColumns,
			uTeacherPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("im: unable to update university_teacher, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `university_teacher` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, uTeacherPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(uTeacherType, uTeacherMapping, append(wl, uTeacherPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "im: unable to update university_teacher row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to get rows affected by update for university_teacher")
	}

	if !cached {
		uTeacherUpdateCacheMut.Lock()
		uTeacherUpdateCache[key] = cache
		uTeacherUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q uTeacherQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to update all for university_teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to retrieve rows affected for university_teacher")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UTeacherSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), uTeacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `university_teacher` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, uTeacherPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to update all in uTeacher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to retrieve rows affected all in update all uTeacher")
	}
	return rowsAff, nil
}

var mySQLUTeacherUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UTeacher) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("im: no university_teacher provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(uTeacherColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUTeacherUniqueColumns, o)

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

	uTeacherUpsertCacheMut.RLock()
	cache, cached := uTeacherUpsertCache[key]
	uTeacherUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			uTeacherAllColumns,
			uTeacherColumnsWithDefault,
			uTeacherColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			uTeacherAllColumns,
			uTeacherPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("im: unable to upsert university_teacher, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "university_teacher", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `university_teacher` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(uTeacherType, uTeacherMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(uTeacherType, uTeacherMapping, ret)
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
		return errors.Wrap(err, "im: unable to upsert for university_teacher")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == uTeacherMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(uTeacherType, uTeacherMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "im: unable to retrieve unique values for university_teacher")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "im: unable to populate default values for university_teacher")
	}

CacheNoHooks:
	if !cached {
		uTeacherUpsertCacheMut.Lock()
		uTeacherUpsertCache[key] = cache
		uTeacherUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UTeacher record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UTeacher) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("im: no UTeacher provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uTeacherPrimaryKeyMapping)
	sql := "DELETE FROM `university_teacher` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to delete from university_teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to get rows affected by delete for university_teacher")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q uTeacherQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("im: no uTeacherQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to delete all from university_teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to get rows affected by deleteall for university_teacher")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UTeacherSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(uTeacherBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), uTeacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `university_teacher` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, uTeacherPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "im: unable to delete all from uTeacher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "im: failed to get rows affected by deleteall for university_teacher")
	}

	if len(uTeacherAfterDeleteHooks) != 0 {
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
func (o *UTeacher) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUTeacher(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UTeacherSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UTeacherSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), uTeacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `university_teacher`.* FROM `university_teacher` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, uTeacherPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "im: unable to reload all in UTeacherSlice")
	}

	*o = slice

	return nil
}

// UTeacherExists checks if the UTeacher row exists.
func UTeacherExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `university_teacher` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "im: unable to check if university_teacher exists")
	}

	return exists, nil
}
