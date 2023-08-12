# Library Management System

It is a library management system written in Golang using the MVC architecture.

## Setup

Clone the repository. In the root use the following commands 

```
go mod vendor 
go mod tidy
```
### Mysql

1. Run this command : `migrate -path database/migration/ -database "mysql://your_db_username:your_db_password@tcp(localhost:3306)/DB_NAME" -verbose up`
2. If this command gives error run command : `migrate -path database/migration/ -database "mysql://your_db_username:your_db_password@tcp(localhost:3306)/DB_NAME" force version`

### Running the server:
1. `go build -o mvc ./cmd/main.go`
2.  Run the binary file: `./mvc`

### For testing 
 
 Run command `go test ./pkg/models`

### Accessing the website
Open localhost:9000 on your browser


## To use Virtual Hosting on Ubuntu

1. Install apache2 : `sudo apt install apache2`
2. `sudo a2enmod proxy proxy_http`
3. `sudo nano your_domain_name.conf` 
4. Copy and paste the virtual host file.
5. `sudo a2ensite your_domain_name.conf`
6. `sudo a2dissite 000-default.conf`
7. `sudo apache2ctl configtest`
8. `sudo nano /etc/hosts`
Add:
```
127.0.0.1  your_domain_name
```
9. `sudo systemctl restart apache2`
10. `sudo systemctl status apache2`
 Check your_domain_name on your browser

 ## Use Script
Run in the terminal
 ```
 chmod +x script/script.sh
 ./script/script.sh
 ```