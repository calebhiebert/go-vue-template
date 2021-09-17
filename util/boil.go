package util

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func CombineQM(mod ...[]qm.QueryMod) (mods []qm.QueryMod) {
	for _, modSet := range mod {
		mods = append(mods, modSet...)
	}

	return
}

func SQM(mod ...qm.QueryMod) []qm.QueryMod {
	return mod
}

func DoTX(ctx context.Context, fn func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(ctx, tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return rbErr
		}

		return err
	} else {
		cmErr := tx.Commit()
		if cmErr != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				return rbErr
			}

			return cmErr
		}
	}

	return nil
}
