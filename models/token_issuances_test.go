// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTokenIssuances(t *testing.T) {
	t.Parallel()

	query := TokenIssuances()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTokenIssuancesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTokenIssuancesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TokenIssuances().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTokenIssuancesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TokenIssuanceSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTokenIssuancesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TokenIssuanceExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TokenIssuance exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TokenIssuanceExists to return true, but got false.")
	}
}

func testTokenIssuancesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	tokenIssuanceFound, err := FindTokenIssuance(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if tokenIssuanceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTokenIssuancesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TokenIssuances().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTokenIssuancesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TokenIssuances().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTokenIssuancesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tokenIssuanceOne := &TokenIssuance{}
	tokenIssuanceTwo := &TokenIssuance{}
	if err = randomize.Struct(seed, tokenIssuanceOne, tokenIssuanceDBTypes, false, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}
	if err = randomize.Struct(seed, tokenIssuanceTwo, tokenIssuanceDBTypes, false, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tokenIssuanceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tokenIssuanceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TokenIssuances().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTokenIssuancesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tokenIssuanceOne := &TokenIssuance{}
	tokenIssuanceTwo := &TokenIssuance{}
	if err = randomize.Struct(seed, tokenIssuanceOne, tokenIssuanceDBTypes, false, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}
	if err = randomize.Struct(seed, tokenIssuanceTwo, tokenIssuanceDBTypes, false, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tokenIssuanceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tokenIssuanceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func tokenIssuanceBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func tokenIssuanceAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func tokenIssuanceAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func tokenIssuanceBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func tokenIssuanceAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func tokenIssuanceBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func tokenIssuanceAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func tokenIssuanceBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func tokenIssuanceAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TokenIssuance) error {
	*o = TokenIssuance{}
	return nil
}

func testTokenIssuancesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TokenIssuance{}
	o := &TokenIssuance{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TokenIssuance object: %s", err)
	}

	AddTokenIssuanceHook(boil.BeforeInsertHook, tokenIssuanceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceBeforeInsertHooks = []TokenIssuanceHook{}

	AddTokenIssuanceHook(boil.AfterInsertHook, tokenIssuanceAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceAfterInsertHooks = []TokenIssuanceHook{}

	AddTokenIssuanceHook(boil.AfterSelectHook, tokenIssuanceAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceAfterSelectHooks = []TokenIssuanceHook{}

	AddTokenIssuanceHook(boil.BeforeUpdateHook, tokenIssuanceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceBeforeUpdateHooks = []TokenIssuanceHook{}

	AddTokenIssuanceHook(boil.AfterUpdateHook, tokenIssuanceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceAfterUpdateHooks = []TokenIssuanceHook{}

	AddTokenIssuanceHook(boil.BeforeDeleteHook, tokenIssuanceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceBeforeDeleteHooks = []TokenIssuanceHook{}

	AddTokenIssuanceHook(boil.AfterDeleteHook, tokenIssuanceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceAfterDeleteHooks = []TokenIssuanceHook{}

	AddTokenIssuanceHook(boil.BeforeUpsertHook, tokenIssuanceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceBeforeUpsertHooks = []TokenIssuanceHook{}

	AddTokenIssuanceHook(boil.AfterUpsertHook, tokenIssuanceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	tokenIssuanceAfterUpsertHooks = []TokenIssuanceHook{}
}

func testTokenIssuancesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTokenIssuancesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(tokenIssuanceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTokenIssuanceToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TokenIssuance
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, tokenIssuanceDBTypes, false, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TokenIssuanceSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*TokenIssuance)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTokenIssuanceToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TokenIssuance
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tokenIssuanceDBTypes, false, strmangle.SetComplement(tokenIssuancePrimaryKeyColumns, tokenIssuanceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TokenIssuances[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testTokenIssuancesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTokenIssuancesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TokenIssuanceSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTokenIssuancesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TokenIssuances().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tokenIssuanceDBTypes = map[string]string{`ID`: `uuid`, `UserID`: `uuid`, `IPAddress`: `character varying`, `CreatedAt`: `timestamp without time zone`}
	_                    = bytes.MinRead
)

func testTokenIssuancesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(tokenIssuancePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(tokenIssuanceAllColumns) == len(tokenIssuancePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuancePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTokenIssuancesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tokenIssuanceAllColumns) == len(tokenIssuancePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TokenIssuance{}
	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tokenIssuanceDBTypes, true, tokenIssuancePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tokenIssuanceAllColumns, tokenIssuancePrimaryKeyColumns) {
		fields = tokenIssuanceAllColumns
	} else {
		fields = strmangle.SetComplement(
			tokenIssuanceAllColumns,
			tokenIssuancePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TokenIssuanceSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTokenIssuancesUpsert(t *testing.T) {
	t.Parallel()

	if len(tokenIssuanceAllColumns) == len(tokenIssuancePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TokenIssuance{}
	if err = randomize.Struct(seed, &o, tokenIssuanceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TokenIssuance: %s", err)
	}

	count, err := TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, tokenIssuanceDBTypes, false, tokenIssuancePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TokenIssuance struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TokenIssuance: %s", err)
	}

	count, err = TokenIssuances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}