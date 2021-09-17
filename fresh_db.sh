bash ./migrate.sh down
bash ./migrate.sh up
cd db/seed || exit
go run .
cd ../..