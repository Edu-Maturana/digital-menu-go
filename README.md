# digital-menu-go
This is a simple CRUD API to create a digital menu, in this case a coffees menu.

##### Clone repository
`git clone https://github.com/Edu-Maturana/digital-menu-go.git`

`cd digital-menu-go`

##### Add the dsn
`rename .env.example to -> .env`

`add your dsn (Postgres database) to DB_URL variable`

##### Run the app
`go run main.go`

##### Run the app as a container
`docker build -t app .`

`docker run -it -p 3000:3000 app`
