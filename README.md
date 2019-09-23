# blog is a simple web site with blog functions

## TODO

* [x] add MySQL storage
* [x] separete api methods and web
* [ ] add MongoDB storage
* [ ] add using BeeGo framework
* [ ] add config app
* [ ] add clear logging
* [ ] add metrics (Prometheus), business and techics
* [ ] GUI refactoring
* [ ] add user register
* [ ] add content valid and moderate
* [ ] add authorization from social networks
* [ ] etc...

## HOW TO USE IN THE HW4

- use $GOPATH
- prepare database
    - use mysql-client `mysql < ./db/blog.sql`
    - OR use go `go run ./db/dbmigrate.sql`
- set MySQL connection string `export DATABASE_URL=root:master@tcp(localhost:3306)/blog`
- navigate to directory with `asset` folder  
- go build -o blog
- start `blog`

### How to use general

![important] 
you must have connection to MySQL server > 8.X  
for creating database `blog` you need privileges to drop and create database and tables  

- install to your $GOPATH `go get -u github.com/art-frela/blog`
- help `blog -h`
- prepare database
    - use mysql-client `mysql < ./db/blog.sql`
    - OR use go `go run ./db/dbmigrate.sql`
- set MySQL connection string `export DATABASE_URL=root:master@tcp(localhost:3306)/blog`
- navigate to directory with `asset` folder  
- start `blog`