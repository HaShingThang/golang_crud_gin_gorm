

- .env file
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=jostthang
DB_NAME=testing


TOKEN_EXPIRED_IN=60m
TOKEN_MAXAGE=60

TOKEN_SECRET=Bs5QWz4YH2J8h8KHZlbSOBY/VChWmfX2NV8V1/eM/Q8=
```

- RESTAPI For Auth
```
http://localhost:8080/api/auth/register
http://localhost:8080/api/auth/login
```
- RESTAPI For User
 ```
 http://localhost:8080/api/users
 http://localhost:8080/api/users/:id
 ```
 - RESTAPI For Post
 ```
 http://localhost:8080/api/posts
 http://localhost:8080/api/posts/:postId
 ```