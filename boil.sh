sqlboiler \
    --add-global-variants --tag-ignore pw_hash --add-soft-deletes \
    --config ./sqlboiler.toml \
    --templates ./templates \
    --templates $GOPATH/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.6.0/templates \
    --templates $GOPATH/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.6.0/templates_test \
    psql

goimports -w models/