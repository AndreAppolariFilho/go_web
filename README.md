# Instructions
To run just create a .env folder with the following parameters

PORT=XXXX
DB_URL=SOME_DB_CONNECTION

The first is the port that you want to build, and the second is the connection string with your database

## Applying migrations
Under the folder sql/schema use the command `goose database_system_name database_connection_url up`, e.g. goose postgres postgres://username:@localhost:DB_PORT/DB up

## Running the server

go build -> to construct the binaries

./rss_aggregator.exe -> to execute