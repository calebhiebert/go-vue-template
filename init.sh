mkdir ui/dist
touch ui/dist/placeholder.txt
go get github.com/99designs/gqlgen/cmd@v0.13.0
go get github.com/swaggo/swag/gen@v1.7.1
go get
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go install github.com/99designs/gqlgen
go install github.com/swaggo/swag/cmd/swag
go install golang.org/x/tools/cmd/goimports
cp env .env