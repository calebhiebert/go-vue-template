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
   1. `go get github.com/99designs/gqlgen`
5. Make sure [Swag](https://github.com/swaggo/swag) is installed
   1. `go install github.com/swaggo/swag/cmd/swag`
6. Copy `./env` to a new file `./.env` and make any neccecary changes
7. Start up the supporting services using the `docker-compose.dev.yml` docker-compose file
8. Start up the go service
9. Start up the frontend
10. ???
11. Profit