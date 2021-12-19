echo "Waiting for DB to start..."
./wait-for database:33060 -- echo "Database Has Started..."
# https://github.com/eficode/wait-for

echo "Preparing Database..."
./app/go_cli_app ./app/data_dump.csv

echo "Running Server..."
./app/go_server_app