# How start Go-Lange Project

1. Create your project folder
2. Go mod init github.com/user_name/repository_name
3. install router framework in this case I use Chi
   <br>
   `go get github.com/go-chi/chi/v5`
   <br>
   `go get github.com/go-chi/chi/v5/middleware`
   <br>
   `go get github.com/go-chi/cors`
   <br>

# How to connect Postgres database

1. Install postgres driver
   <br>

```
go get github.com/jackc/pgconn
<br>
got get github.com/jackc/pgx/v4
<br>
got get github.com/jackc/pgx/v4/stdlib
```
