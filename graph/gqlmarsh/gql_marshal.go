package gqlmarsh

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/volatiletech/null/v8"
	"io"
	"strconv"
	"time"
)

// MarshalNullString is a custom marshaller.
func MarshalNullString(ns null.String) graphql.Marshaler {
	if !ns.Valid {
		// this is also important, so we can detect if this scalar is used in a not null context and return an appropriate error
		return graphql.Null
	}
	return graphql.MarshalString(ns.String)
}

// UnmarshalNullString is a custom unmarshaller.
func UnmarshalNullString(v interface{}) (null.String, error) {
	if v == nil {
		return null.String{Valid: false}, nil
	}
	// again you can delegate to the default implementation to save yourself some work.
	s, err := graphql.UnmarshalString(v)
	return null.String{String: s, Valid: true}, err
}

// MarshalNullInt is a custom marshaller.
func MarshalNullInt(ns null.Int) graphql.Marshaler {
	if !ns.Valid {
		// this is also important, so we can detect if this scalar is used in a not null context and return an appropriate error
		return graphql.Null
	}
	return graphql.MarshalInt(ns.Int)
}

// UnmarshalNullInt is a custom unmarshaller.
func UnmarshalNullInt(v interface{}) (null.Int, error) {
	if v == nil {
		return null.Int{Valid: false}, nil
	}
	// again you can delegate to the default implementation to save yourself some work.
	s, err := graphql.UnmarshalInt(v)
	return null.Int{Int: s, Valid: true}, err
}

// MarshalNullBoolean is a custom marshaller.
func MarshalNullBoolean(ns null.Bool) graphql.Marshaler {
	if !ns.Valid {
		// this is also important, so we can detect if this scalar is used in a not null context and return an appropriate error
		return graphql.Null
	}
	return graphql.MarshalBoolean(ns.Bool)
}

// UnmarshalNullBoolean is a custom unmarshaller.
func UnmarshalNullBoolean(v interface{}) (null.Bool, error) {
	if v == nil {
		return null.Bool{Valid: false}, nil
	}

	// again you can delegate to the default implementation to save yourself some work.
	s, err := graphql.UnmarshalBoolean(v)
	return null.Bool{Bool: s, Valid: true}, err
}

func MarshalNullTime(t null.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Time.Format(time.RFC3339)))
	})
}

func UnmarshalNullTime(v interface{}) (null.Time, error) {
	if tmpStr, ok := v.(string); ok {
		t, err := time.Parse(time.RFC3339, tmpStr)
		if err != nil {
			return null.Time{}, err
		}

		return null.TimeFrom(t), nil
	}

	return null.Time{}, errors.New("time should be RFC3339 formatted string")
}