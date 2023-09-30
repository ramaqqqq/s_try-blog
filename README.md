
# claim : Developed by Lutfi M

- Golang (gorilla mux)
- MySQL

## Application Manual

1. Please make DB with name "blogs".
2. Please adjust the .env file for DB_HOST(localhost) and DB_PASSWORD according to your preferences.
2. How to Run :

   ```
   go run . / go run main.go
   ```

## Or by Dockerfile and Docker Compose

- docker build -t backend-go-sagara:v1 .
- docker-compose up -d
- optional : sudo docker-compose restart backend-go

## Endpoints List

- USER
[POST]       `localhost:7000/api/login` : Login user </br>
[POST]       `localhost:7000/api/register` : Register </br>
[GET]       `localhost:7000/api/user/{user_id}/single` : Get User by Id </br>
[PUT]       `localhost:7000/api/user/{user_id}/edit` : Edit User by Id </br>
[DELETE]       `localhost:7000/api/user/{user_id}/delete` : Delete User by Id </br>
</br>

- BLOG
[POST]      `localhost:7000/api/blog` : Add Blog</br>
[GET]       `localhost:7000/api/blog` : All Blog</br>
[GET]       `localhost:7000/api/blog/paginated?page=1&limit=3` : Blog by paginated </br>
[GET]       `localhost:7000/api/blog/{blog_id}/single` : Get One Blog by blog_id</br>
[PUT]       `localhost:7000/api/blog/{blog_id}/edit` : Edit Blog by blog_id</br>
[DELETE]    `localhost:7000/api/blog/{blog_id}/delete` : Delete Blog by blog_id</br>
</br>
