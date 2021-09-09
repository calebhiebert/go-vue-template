# Usage: ./migrate.sh [up|down|etc...]
# See https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

# Load environment variables from .env file
export $(egrep -v '^#' .env | xargs)

# Migrate the database to the latest version
migrate -database ${DATABASE_URL} -path ./db/migrations $1