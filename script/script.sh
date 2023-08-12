#!/bin/bash

# Get MySQL database connection parameters from user
read -p "Enter MySQL host: " DB_HOST
read -p "Enter MySQL username: " DB_USERNAME
read -s -p "Enter MySQL password: " DB_PASSWORD
echo
read -p "Enter MySQL database name: " DB_NAME

migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp(localhost:3306)/$DB_NAME" -verbose up

cat <<EOF > config.yaml
DB_USERNAME: $DB_USERNAME
DB_PASSWORD: '$DB_PASSWORD'
DB_HOST: $DB_HOST
DB_NAME: $DB_NAME
EOF

# Open a MySQL database connection
mysql -h $DB_HOST -u $DB_USERNAME -p$DB_PASSWORD $DB_NAME -e "Select * from users"


# Build the server binary
go mod vendor
go mod tidy
go build -o lms ./cmd/main.go
./lms
 echo "Server running ..."