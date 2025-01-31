library-kata
-----

## Run 

run the app:
```
go run main.go
```


## APP API

* get all users
```
curl --request GET \
  --url http://localhost:3000/users \
  --header 'Content-Type: application/json' 
```

* add user
```
curl --request POST \
  --url http://localhost:3000/user \
  --header 'Content-Type: application/json' \
  --data '{"name":"Roman"}'
```

* get all books
```
curl --request GET \
  --url http://localhost:3000/books \
  --header 'Content-Type: application/json' 
```

* add book
```
curl --request POST \
  --url http://localhost:3000/book \
  --header 'Content-Type: application/json' \
  --data '{"title": "Harry Potter"}'
```

* borrow
```
curl --request POST \
  --url http://localhost:3000/borrow \
  --header 'Content-Type: application/json' \
  --data '{"user_id": "cb0aa0ba-3c04-4ee4-9c2c-a34bd3dbc7de", "book_id": "d9914b4f-d0ec-405b-b1cc-1387d489bc5e"}'
```


* return
```
curl --request POST \
  --url http://localhost:3000/return \
  --header 'Content-Type: application/json' \
  --data '{"user_id": "cb0aa0ba-3c04-4ee4-9c2c-a34bd3dbc7de", "book_id": "d9914b4f-d0ec-405b-b1cc-1387d489bc5e"}'
```

* get books borrowed by a user
```
curl --request GET \
  --url 'http://localhost:3000/user-books?user_id=cb0aa0ba-3c04-4ee4-9c2c-a34bd3dbc7de' \
  --header 'Content-Type: application/json' 
```
