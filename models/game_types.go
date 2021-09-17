// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// GameType is an object representing the database table.
type GameType struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	IsCustom  bool      `boil:"is_custom" json:"is_custom" toml:"is_custom" yaml:"is_custom"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *gameTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gameTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GameTypeColumns = struct {
	ID        string
	Name      string
	IsCustom  string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	IsCustom:  "is_custom",
	DeletedAt: "deleted_at",
}

var GameTypeTableColumns = struct {
	ID        string
	Name      string
	IsCustom  string
	DeletedAt string
}{
	ID:        "game_types.id",
	Name:      "game_types.name",
	IsCustom:  "game_types.is_custom",
	DeletedAt: "game_types.deleted_at",
}

// Generated where

var GameTypeWhere = struct {
	ID        whereHelperstring
	Name      whereHelperstring
	IsCustom  whereHelperbool
	DeletedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"game_types\".\"id\""},
	Name:      whereHelperstring{field: "\"game_types\".\"name\""},
	IsCustom:  whereHelperbool{field: "\"game_types\".\"is_custom\""},
	DeletedAt: whereHelpernull_Time{field: "\"game_types\".\"deleted_at\""},
}

// GameTypeRels is where relationship names are stored.
var GameTypeRels = struct {
	Events string
}{
	Events: "Events",
}

// gameTypeR is where relationships are stored.
type gameTypeR struct {
	Events EventSlice `boil:"Events" json:"Events" toml:"Events" yaml:"Events"`
}

// NewStruct creates a new relationship struct
func (*gameTypeR) NewStruct() *gameTypeR {
	return &gameTypeR{}
}

// gameTypeL is where Load methods for each relationship are stored.
type gameTypeL struct{}

var (
	gameTypeAllColumns            = []string{"id", "name", "is_custom", "deleted_at"}
	gameTypeColumnsWithoutDefault = []string{"id", "name", "deleted_at"}
	gameTypeColumnsWithDefault    = []string{"is_custom"}
	gameTypePrimaryKeyColumns     = []string{"id"}
)

type (
	// GameTypeSlice is an alias for a slice of pointers to GameType.
	// This should almost always be used instead of []GameType.
	GameTypeSlice []*GameType
	// GameTypeHook is the signature for custom GameType hook methods
	GameTypeHook func(context.Context, boil.ContextExecutor, *GameType) error

	gameTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gameTypeType                 = reflect.TypeOf(&GameType{})
	gameTypeMapping              = queries.MakeStructMapping(gameTypeType)
	gameTypePrimaryKeyMapping, _ = queries.BindMapping(gameTypeType, gameTypeMapping, gameTypePrimaryKeyColumns)
	gameTypeInsertCacheMut       sync.RWMutex
	gameTypeInsertCache          = make(map[string]insertCache)
	gameTypeUpdateCacheMut       sync.RWMutex
	gameTypeUpdateCache          = make(map[string]updateCache)
	gameTypeUpsertCacheMut       sync.RWMutex
	gameTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gameTypeBeforeInsertHooks []GameTypeHook
var gameTypeBeforeUpdateHooks []GameTypeHook
var gameTypeBeforeDeleteHooks []GameTypeHook
var gameTypeBeforeUpsertHooks []GameTypeHook

var gameTypeAfterInsertHooks []GameTypeHook
var gameTypeAfterSelectHooks []GameTypeHook
var gameTypeAfterUpdateHooks []GameTypeHook
var gameTypeAfterDeleteHooks []GameTypeHook
var gameTypeAfterUpsertHooks []GameTypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GameType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GameType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GameType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GameType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GameType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GameType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GameType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GameType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GameType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGameTypeHook registers your hook function for all future operations.
func AddGameTypeHook(hookPoint boil.HookPoint, gameTypeHook GameTypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		gameTypeBeforeInsertHooks = append(gameTypeBeforeInsertHooks, gameTypeHook)
	case boil.BeforeUpdateHook:
		gameTypeBeforeUpdateHooks = append(gameTypeBeforeUpdateHooks, gameTypeHook)
	case boil.BeforeDeleteHook:
		gameTypeBeforeDeleteHooks = append(gameTypeBeforeDeleteHooks, gameTypeHook)
	case boil.BeforeUpsertHook:
		gameTypeBeforeUpsertHooks = append(gameTypeBeforeUpsertHooks, gameTypeHook)
	case boil.AfterInsertHook:
		gameTypeAfterInsertHooks = append(gameTypeAfterInsertHooks, gameTypeHook)
	case boil.AfterSelectHook:
		gameTypeAfterSelectHooks = append(gameTypeAfterSelectHooks, gameTypeHook)
	case boil.AfterUpdateHook:
		gameTypeAfterUpdateHooks = append(gameTypeAfterUpdateHooks, gameTypeHook)
	case boil.AfterDeleteHook:
		gameTypeAfterDeleteHooks = append(gameTypeAfterDeleteHooks, gameTypeHook)
	case boil.AfterUpsertHook:
		gameTypeAfterUpsertHooks = append(gameTypeAfterUpsertHooks, gameTypeHook)
	}
}

// OneG returns a single gameType record from the query using the global executor.
func (q gameTypeQuery) OneG(ctx context.Context) (*GameType, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single gameType record from the query.
func (q gameTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GameType, error) {
	o := &GameType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for game_types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all GameType records from the query using the global executor.
func (q gameTypeQuery) AllG(ctx context.Context) (GameTypeSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all GameType records from the query.
func (q gameTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (GameTypeSlice, error) {
	var o []*GameType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GameType slice")
	}

	if len(gameTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all GameType records in the query, and panics on error.
func (q gameTypeQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all GameType records in the query.
func (q gameTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count game_types rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q gameTypeQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q gameTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if game_types exists")
	}

	return count > 0, nil
}

// Events retrieves all the event's Events with an executor.
func (o *GameType) Events(mods ...qm.QueryMod) eventQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"events\".\"game_type_id\"=?", o.ID),
		qmhelper.WhereIsNull("\"events\".\"deleted_at\""),
	)

	query := Events(queryMods...)
	queries.SetFrom(query.Query, "\"events\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"events\".*"})
	}

	return query
}

// LoadEvents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (gameTypeL) LoadEvents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGameType interface{}, mods queries.Applicator) error {
	var slice []*GameType
	var object *GameType

	if singular {
		object = maybeGameType.(*GameType)
	} else {
		slice = *maybeGameType.(*[]*GameType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gameTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gameTypeR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`events`),
		qm.WhereIn(`events.game_type_id in ?`, args...),
		qmhelper.WhereIsNull(`events.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load events")
	}

	var resultSlice []*Event
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice events")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on events")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for events")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Events = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &eventR{}
			}
			foreign.R.GameType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GameTypeID {
				local.R.Events = append(local.R.Events, foreign)
				if foreign.R == nil {
					foreign.R = &eventR{}
				}
				foreign.R.GameType = local
				break
			}
		}
	}

	return nil
}

// AddEventsG adds the given related objects to the existing relationships
// of the game_type, optionally inserting them as new records.
// Appends related to o.R.Events.
// Sets related.R.GameType appropriately.
// Uses the global database handle.
func (o *GameType) AddEventsG(ctx context.Context, insert bool, related ...*Event) error {
	return o.AddEvents(ctx, boil.GetContextDB(), insert, related...)
}

// AddEvents adds the given related objects to the existing relationships
// of the game_type, optionally inserting them as new records.
// Appends related to o.R.Events.
// Sets related.R.GameType appropriately.
func (o *GameType) AddEvents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Event) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GameTypeID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"events\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"game_type_id"}),
				strmangle.WhereClause("\"", "\"", 2, eventPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GameTypeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &gameTypeR{
			Events: related,
		}
	} else {
		o.R.Events = append(o.R.Events, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &eventR{
				GameType: o,
			}
		} else {
			rel.R.GameType = o
		}
	}
	return nil
}

// GameTypes retrieves all the records using an executor.
func GameTypes(mods ...qm.QueryMod) gameTypeQuery {
	mods = append(mods, qm.From("\"game_types\""), qmhelper.WhereIsNull("\"game_types\".\"deleted_at\""))
	return gameTypeQuery{NewQuery(mods...)}
}

// FindGameTypeG retrieves a single record by ID.
func FindGameTypeG(ctx context.Context, iD string, selectCols ...string) (*GameType, error) {
	return FindGameType(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindGameType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGameType(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*GameType, error) {
	gameTypeObj := &GameType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"game_types\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gameTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from game_types")
	}

	if err = gameTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return gameTypeObj, err
	}

	return gameTypeObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *GameType) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GameType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no game_types provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gameTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gameTypeInsertCacheMut.RLock()
	cache, cached := gameTypeInsertCache[key]
	gameTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gameTypeAllColumns,
			gameTypeColumnsWithDefault,
			gameTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gameTypeType, gameTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gameTypeType, gameTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"game_types\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"game_types\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into game_types")
	}

	if !cached {
		gameTypeInsertCacheMut.Lock()
		gameTypeInsertCache[key] = cache
		gameTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single GameType record using the global executor.
// See Update for more documentation.
func (o *GameType) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the GameType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GameType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gameTypeUpdateCacheMut.RLock()
	cache, cached := gameTypeUpdateCache[key]
	gameTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gameTypeAllColumns,
			gameTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update game_types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"game_types\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gameTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gameTypeType, gameTypeMapping, append(wl, gameTypePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update game_types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for game_types")
	}

	if !cached {
		gameTypeUpdateCacheMut.Lock()
		gameTypeUpdateCache[key] = cache
		gameTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q gameTypeQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q gameTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for game_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for game_types")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GameTypeSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GameTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"game_types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gameTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in gameType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all gameType")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *GameType) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GameType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no game_types provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gameTypeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	gameTypeUpsertCacheMut.RLock()
	cache, cached := gameTypeUpsertCache[key]
	gameTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			gameTypeAllColumns,
			gameTypeColumnsWithDefault,
			gameTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			gameTypeAllColumns,
			gameTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert game_types, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(gameTypePrimaryKeyColumns))
			copy(conflict, gameTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"game_types\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(gameTypeType, gameTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gameTypeType, gameTypeMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert game_types")
	}

	if !cached {
		gameTypeUpsertCacheMut.Lock()
		gameTypeUpsertCache[key] = cache
		gameTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single GameType record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *GameType) DeleteG(ctx context.Context, hardDelete bool) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB(), hardDelete)
}

// Delete deletes a single GameType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GameType) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GameType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gameTypePrimaryKeyMapping)
		sql = "DELETE FROM \"game_types\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"game_types\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(gameTypeType, gameTypeMapping, append(wl, gameTypePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from game_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for game_types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q gameTypeQuery) DeleteAllG(ctx context.Context, hardDelete bool) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB(), hardDelete)
}

// DeleteAll deletes all matching rows.
func (q gameTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no gameTypeQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from game_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for game_types")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o GameTypeSlice) DeleteAllG(ctx context.Context, hardDelete bool) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB(), hardDelete)
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GameTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gameTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameTypePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"game_types\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gameTypePrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameTypePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"game_types\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, gameTypePrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gameType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for game_types")
	}

	if len(gameTypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *GameType) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no GameType provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GameType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGameType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GameTypeSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty GameTypeSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GameTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GameTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"game_types\".* FROM \"game_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gameTypePrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GameTypeSlice")
	}

	*o = slice

	return nil
}

// GameTypeExistsG checks if the GameType row exists.
func GameTypeExistsG(ctx context.Context, iD string) (bool, error) {
	return GameTypeExists(ctx, boil.GetContextDB(), iD)
}

// GameTypeExists checks if the GameType row exists.
func GameTypeExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"game_types\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if game_types exists")
	}

	return exists, nil
}
