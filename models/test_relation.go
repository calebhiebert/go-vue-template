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

// TestRelation is an object representing the database table.
type TestRelation struct {
	ID          string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	TestTableID null.String `boil:"test_table_id" json:"test_table_id,omitempty" toml:"test_table_id" yaml:"test_table_id,omitempty"`

	R *testRelationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L testRelationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TestRelationColumns = struct {
	ID          string
	Name        string
	TestTableID string
}{
	ID:          "id",
	Name:        "name",
	TestTableID: "test_table_id",
}

var TestRelationTableColumns = struct {
	ID          string
	Name        string
	TestTableID string
}{
	ID:          "test_relation.id",
	Name:        "test_relation.name",
	TestTableID: "test_relation.test_table_id",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var TestRelationWhere = struct {
	ID          whereHelperstring
	Name        whereHelperstring
	TestTableID whereHelpernull_String
}{
	ID:          whereHelperstring{field: "\"test_relation\".\"id\""},
	Name:        whereHelperstring{field: "\"test_relation\".\"name\""},
	TestTableID: whereHelpernull_String{field: "\"test_relation\".\"test_table_id\""},
}

// TestRelationRels is where relationship names are stored.
var TestRelationRels = struct {
	TestTable string
}{
	TestTable: "TestTable",
}

// testRelationR is where relationships are stored.
type testRelationR struct {
	TestTable *TestTable `boil:"TestTable" json:"TestTable" toml:"TestTable" yaml:"TestTable"`
}

// NewStruct creates a new relationship struct
func (*testRelationR) NewStruct() *testRelationR {
	return &testRelationR{}
}

// testRelationL is where Load methods for each relationship are stored.
type testRelationL struct{}

var (
	testRelationAllColumns            = []string{"id", "name", "test_table_id"}
	testRelationColumnsWithoutDefault = []string{"id", "name", "test_table_id"}
	testRelationColumnsWithDefault    = []string{}
	testRelationPrimaryKeyColumns     = []string{"id"}
)

type (
	// TestRelationSlice is an alias for a slice of pointers to TestRelation.
	// This should almost always be used instead of []TestRelation.
	TestRelationSlice []*TestRelation
	// TestRelationHook is the signature for custom TestRelation hook methods
	TestRelationHook func(context.Context, boil.ContextExecutor, *TestRelation) error

	testRelationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	testRelationType                 = reflect.TypeOf(&TestRelation{})
	testRelationMapping              = queries.MakeStructMapping(testRelationType)
	testRelationPrimaryKeyMapping, _ = queries.BindMapping(testRelationType, testRelationMapping, testRelationPrimaryKeyColumns)
	testRelationInsertCacheMut       sync.RWMutex
	testRelationInsertCache          = make(map[string]insertCache)
	testRelationUpdateCacheMut       sync.RWMutex
	testRelationUpdateCache          = make(map[string]updateCache)
	testRelationUpsertCacheMut       sync.RWMutex
	testRelationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var testRelationBeforeInsertHooks []TestRelationHook
var testRelationBeforeUpdateHooks []TestRelationHook
var testRelationBeforeDeleteHooks []TestRelationHook
var testRelationBeforeUpsertHooks []TestRelationHook

var testRelationAfterInsertHooks []TestRelationHook
var testRelationAfterSelectHooks []TestRelationHook
var testRelationAfterUpdateHooks []TestRelationHook
var testRelationAfterDeleteHooks []TestRelationHook
var testRelationAfterUpsertHooks []TestRelationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TestRelation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TestRelation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TestRelation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TestRelation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TestRelation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TestRelation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TestRelation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TestRelation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TestRelation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testRelationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTestRelationHook registers your hook function for all future operations.
func AddTestRelationHook(hookPoint boil.HookPoint, testRelationHook TestRelationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		testRelationBeforeInsertHooks = append(testRelationBeforeInsertHooks, testRelationHook)
	case boil.BeforeUpdateHook:
		testRelationBeforeUpdateHooks = append(testRelationBeforeUpdateHooks, testRelationHook)
	case boil.BeforeDeleteHook:
		testRelationBeforeDeleteHooks = append(testRelationBeforeDeleteHooks, testRelationHook)
	case boil.BeforeUpsertHook:
		testRelationBeforeUpsertHooks = append(testRelationBeforeUpsertHooks, testRelationHook)
	case boil.AfterInsertHook:
		testRelationAfterInsertHooks = append(testRelationAfterInsertHooks, testRelationHook)
	case boil.AfterSelectHook:
		testRelationAfterSelectHooks = append(testRelationAfterSelectHooks, testRelationHook)
	case boil.AfterUpdateHook:
		testRelationAfterUpdateHooks = append(testRelationAfterUpdateHooks, testRelationHook)
	case boil.AfterDeleteHook:
		testRelationAfterDeleteHooks = append(testRelationAfterDeleteHooks, testRelationHook)
	case boil.AfterUpsertHook:
		testRelationAfterUpsertHooks = append(testRelationAfterUpsertHooks, testRelationHook)
	}
}

// OneG returns a single testRelation record from the query using the global executor.
func (q testRelationQuery) OneG(ctx context.Context) (*TestRelation, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single testRelation record from the query.
func (q testRelationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TestRelation, error) {
	o := &TestRelation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for test_relation")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TestRelation records from the query using the global executor.
func (q testRelationQuery) AllG(ctx context.Context) (TestRelationSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TestRelation records from the query.
func (q testRelationQuery) All(ctx context.Context, exec boil.ContextExecutor) (TestRelationSlice, error) {
	var o []*TestRelation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TestRelation slice")
	}

	if len(testRelationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TestRelation records in the query, and panics on error.
func (q testRelationQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TestRelation records in the query.
func (q testRelationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count test_relation rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q testRelationQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q testRelationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if test_relation exists")
	}

	return count > 0, nil
}

// TestTable pointed to by the foreign key.
func (o *TestRelation) TestTable(mods ...qm.QueryMod) testTableQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TestTableID),
	}

	queryMods = append(queryMods, mods...)

	query := TestTables(queryMods...)
	queries.SetFrom(query.Query, "\"test_table\"")

	return query
}

// LoadTestTable allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (testRelationL) LoadTestTable(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTestRelation interface{}, mods queries.Applicator) error {
	var slice []*TestRelation
	var object *TestRelation

	if singular {
		object = maybeTestRelation.(*TestRelation)
	} else {
		slice = *maybeTestRelation.(*[]*TestRelation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &testRelationR{}
		}
		if !queries.IsNil(object.TestTableID) {
			args = append(args, object.TestTableID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &testRelationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.TestTableID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.TestTableID) {
				args = append(args, obj.TestTableID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`test_table`),
		qm.WhereIn(`test_table.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TestTable")
	}

	var resultSlice []*TestTable
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TestTable")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for test_table")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for test_table")
	}

	if len(testRelationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.TestTable = foreign
		if foreign.R == nil {
			foreign.R = &testTableR{}
		}
		foreign.R.TestRelations = append(foreign.R.TestRelations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.TestTableID, foreign.ID) {
				local.R.TestTable = foreign
				if foreign.R == nil {
					foreign.R = &testTableR{}
				}
				foreign.R.TestRelations = append(foreign.R.TestRelations, local)
				break
			}
		}
	}

	return nil
}

// SetTestTableG of the testRelation to the related item.
// Sets o.R.TestTable to related.
// Adds o to related.R.TestRelations.
// Uses the global database handle.
func (o *TestRelation) SetTestTableG(ctx context.Context, insert bool, related *TestTable) error {
	return o.SetTestTable(ctx, boil.GetContextDB(), insert, related)
}

// SetTestTable of the testRelation to the related item.
// Sets o.R.TestTable to related.
// Adds o to related.R.TestRelations.
func (o *TestRelation) SetTestTable(ctx context.Context, exec boil.ContextExecutor, insert bool, related *TestTable) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"test_relation\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"test_table_id"}),
		strmangle.WhereClause("\"", "\"", 2, testRelationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.TestTableID, related.ID)
	if o.R == nil {
		o.R = &testRelationR{
			TestTable: related,
		}
	} else {
		o.R.TestTable = related
	}

	if related.R == nil {
		related.R = &testTableR{
			TestRelations: TestRelationSlice{o},
		}
	} else {
		related.R.TestRelations = append(related.R.TestRelations, o)
	}

	return nil
}

// RemoveTestTableG relationship.
// Sets o.R.TestTable to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *TestRelation) RemoveTestTableG(ctx context.Context, related *TestTable) error {
	return o.RemoveTestTable(ctx, boil.GetContextDB(), related)
}

// RemoveTestTable relationship.
// Sets o.R.TestTable to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *TestRelation) RemoveTestTable(ctx context.Context, exec boil.ContextExecutor, related *TestTable) error {
	var err error

	queries.SetScanner(&o.TestTableID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("test_table_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.TestTable = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.TestRelations {
		if queries.Equal(o.TestTableID, ri.TestTableID) {
			continue
		}

		ln := len(related.R.TestRelations)
		if ln > 1 && i < ln-1 {
			related.R.TestRelations[i] = related.R.TestRelations[ln-1]
		}
		related.R.TestRelations = related.R.TestRelations[:ln-1]
		break
	}
	return nil
}

// TestRelations retrieves all the records using an executor.
func TestRelations(mods ...qm.QueryMod) testRelationQuery {
	mods = append(mods, qm.From("\"test_relation\""))
	return testRelationQuery{NewQuery(mods...)}
}

// FindTestRelationG retrieves a single record by ID.
func FindTestRelationG(ctx context.Context, iD string, selectCols ...string) (*TestRelation, error) {
	return FindTestRelation(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTestRelation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTestRelation(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*TestRelation, error) {
	testRelationObj := &TestRelation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"test_relation\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, testRelationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from test_relation")
	}

	if err = testRelationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return testRelationObj, err
	}

	return testRelationObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TestRelation) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TestRelation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no test_relation provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(testRelationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	testRelationInsertCacheMut.RLock()
	cache, cached := testRelationInsertCache[key]
	testRelationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			testRelationAllColumns,
			testRelationColumnsWithDefault,
			testRelationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(testRelationType, testRelationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(testRelationType, testRelationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"test_relation\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"test_relation\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into test_relation")
	}

	if !cached {
		testRelationInsertCacheMut.Lock()
		testRelationInsertCache[key] = cache
		testRelationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TestRelation record using the global executor.
// See Update for more documentation.
func (o *TestRelation) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TestRelation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TestRelation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	testRelationUpdateCacheMut.RLock()
	cache, cached := testRelationUpdateCache[key]
	testRelationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			testRelationAllColumns,
			testRelationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update test_relation, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"test_relation\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, testRelationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(testRelationType, testRelationMapping, append(wl, testRelationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update test_relation row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for test_relation")
	}

	if !cached {
		testRelationUpdateCacheMut.Lock()
		testRelationUpdateCache[key] = cache
		testRelationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q testRelationQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q testRelationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for test_relation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for test_relation")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TestRelationSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TestRelationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), testRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"test_relation\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, testRelationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in testRelation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all testRelation")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TestRelation) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TestRelation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no test_relation provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(testRelationColumnsWithDefault, o)

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

	testRelationUpsertCacheMut.RLock()
	cache, cached := testRelationUpsertCache[key]
	testRelationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			testRelationAllColumns,
			testRelationColumnsWithDefault,
			testRelationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			testRelationAllColumns,
			testRelationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert test_relation, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(testRelationPrimaryKeyColumns))
			copy(conflict, testRelationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"test_relation\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(testRelationType, testRelationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(testRelationType, testRelationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert test_relation")
	}

	if !cached {
		testRelationUpsertCacheMut.Lock()
		testRelationUpsertCache[key] = cache
		testRelationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TestRelation record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TestRelation) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TestRelation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TestRelation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TestRelation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), testRelationPrimaryKeyMapping)
	sql := "DELETE FROM \"test_relation\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from test_relation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for test_relation")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q testRelationQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q testRelationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no testRelationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from test_relation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for test_relation")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TestRelationSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TestRelationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(testRelationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), testRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"test_relation\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, testRelationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from testRelation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for test_relation")
	}

	if len(testRelationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TestRelation) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no TestRelation provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TestRelation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTestRelation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TestRelationSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TestRelationSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TestRelationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TestRelationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), testRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"test_relation\".* FROM \"test_relation\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, testRelationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TestRelationSlice")
	}

	*o = slice

	return nil
}

// TestRelationExistsG checks if the TestRelation row exists.
func TestRelationExistsG(ctx context.Context, iD string) (bool, error) {
	return TestRelationExists(ctx, boil.GetContextDB(), iD)
}

// TestRelationExists checks if the TestRelation row exists.
func TestRelationExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"test_relation\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if test_relation exists")
	}

	return exists, nil
}
