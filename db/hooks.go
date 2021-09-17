package db

import (
	"context"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/calebhiebert/go-vue-template/util"
	"github.com/nyaruka/phonenumbers"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func addHooks() {
	models.AddUserHook(boil.BeforeUpdateHook, BeforeUserUpdate)

	models.AddUserHook(boil.BeforeInsertHook, func(ctx context.Context, executor boil.ContextExecutor, user *models.User) error {
		if !user.Phone.IsZero() {
			pn, err := phonenumbers.Parse(user.Phone.String, "CA")
			if err != nil {
				return err
			}

			user.Phone = null.StringFrom(phonenumbers.Format(pn, phonenumbers.E164))
		}

		return nil
	})
}

func BeforeUserUpdate(ctx context.Context, executor boil.ContextExecutor, user *models.User) error {
	// Validate user permissions
	err := util.ACAdminOr(ctx, func(u *models.User) (bool, error) {
		return u.ID == user.ID, nil
	})
	if err != nil {
		return err
	}

	if !user.Phone.IsZero() {
		pn, err := phonenumbers.Parse(user.Phone.String, "CA")
		if err != nil {
			return err
		}

		user.Phone = null.StringFrom(phonenumbers.Format(pn, phonenumbers.E164))
	}

	return nil
}
