# Go Vue Template

This is a little template project to create apps using 

- Golang (With [Gin](https://github.com/gin-gonic/gin) web framework)
- PostgreSQL database using [SQLBoiler](https://github.com/volatiletech/sqlboiler)
- Database migrations using [Migrate](https://github.com/golang-migrate/migrate)
- Frontend using [Vue 2](https://vuejs.org/)

## Setup Instructions

1. Install golang packages
   1. `go get`
2. Make sure [Migrate](https://github.com/golang-migrate/migrate) is installed
3. Make sure [SQLBoiler](https://github.com/volatiletech/sqlboiler) is installed
   1. `go install github.com/volatiletech/sqlboiler/v4@latest`
   2. `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest`
4. Make sure [GQLGen](https://github.com/99designs/gqlgen) is installed
   1. `go install github.com/99designs/gqlgen`
5. Make sure [Swag](https://github.com/swaggo/swag) is installed
   1. `go install github.com/swaggo/swag/cmd/swag`
6. Make sure [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) is installed
   1. `go install golang.org/x/tools/cmd/goimports`
7. Copy `./env` to a new file `./.env` and make any neccecary changes
8. Start up the supporting services using the `docker-compose.dev.yml` docker-compose file
9. Migrate the database
   1. `./migrate.sh up`
10. (Optional) Freshen up the database. (This will totally reset the database, and run the seed script in db/seed/seed.go)
    1. `./fresh_db.sh`
11. Start up the frontend
    1. `cd ui`
    2. `yarn build --development --watch`
12. Start up the go service
    1. `go run .`
13. Two users are auto generated when migrating the database
    1. ```
       login: admin@admin.admin
       password: admin
       ```
    2. ```
       login: user@user.user
       password: bestuser
       ```


## About

The aim of this project is to remove a lot of the boilerplate from making an api. The general usage is
1. Create database models via migration (`./create_migration.sh <name of migration>`)
2. run `go generate` (if there is errors, try running `go generate` again, don't ask me why)
3. sqlboiler will generate db model structs, and orm-like code for accessing the database
4. sqlboiler will run some custom templates in the `./templates` directory
   1. `./templates/modelcrud` will generate a CRUD api that is automatically used with the admin interface
   2. `./templates/gql` will generate types for use in graphql, and helpers for creating/updating/deleting entities
5. gqlgen will use the `.graphql` and `.graphqls` files in the `./graph` directory to generate resolvers
   1. These resolvers can be found in `./graph/resolve`, and filled in with custom code
6. a swagger api document will be generated based on comments


## Usage


