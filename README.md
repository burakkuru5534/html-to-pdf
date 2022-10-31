# User Management Service

## Introduction

In this project, we will be building a user management service.

### Languages and frameworks

Technologies used in this project:

Golang,
postgresql

Test Environments:

postman,
golang testing library,

### Database

Postgresql was used as the database language.

Tables created:

table name:usr
columns:
id: serial primary key
name: text
password: text
email: text (unique)

## Problem solution

We should be able to create, read, update, and delete users.
We should prevent users from creating duplicate emails.

### Create User

Create User request url example:

Method: POST

http://localhost:8080/users

request Body Example:

 ```json
{
  "name": "Burak KURU",
  "email": "burak.kuru@gmail.com",
  "password": "test123456"
}
 ```

response example:

for 200:

 ```json
{
  "id": 1,
  "name": "Burak KURU",
  "email": "burak.kuru@gmail.com"
}
 ```

for 400:
    
```json
{"error": "Bad request"}
```

for 403:
```json
{"error": "User with that email already exists"}
```

for 500:
```json
{"error": "server error"}
```

### Get User 

Get User request url example:

Method: GET

http://localhost:8080/user?id=1

id: this id should be one of the user's ids.

request Body:

response example:

for 200:
 ```json
{
  "id": 1,
  "name": "Burak KURU",
  "email": "burak.kuru@gmail.com"
}
 ```

for 400:

```json
{"error": "Bad request"}
```

for 404:
```json
{"error": "User with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```

### Update User

Update User request url example:
Method: PATCH

http://localhost:8080/users?id=1

id: this id should be one of the user's ids.
request Body Example:

 ```json
{
  "name":"Updated Name",
  "email":"updatedemail@gmail.com",
  "password":"updatedPass"
}
 ```

response example:

for 200:
 ```json
{
  "id": 1,
  "name":"Updated Name",
  "email":"updatedemail@gmail.com",
  "password":"updatedPass"
}
 ```

for 403:
```json
{"error": "User with that email already exists"}
```

for 400:

```json
{"error": "Bad request"}
```

for 404:
```json
{"error": "User with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```

### Delete User

Delete User request url example:
Method: DELETE

http://localhost:8080/users?id=1

id: this id should be one of the user's ids.

response example:

for 200: "ok"


for 400:

```json
{"error": "Bad request"}
```

for 404:
```json
{"error": "User with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```

### User List

User List request url example:
Method: GET

http://localhost:8080/users

response example:

for 200:
 ```json
[{
  "name":"Updated Name",
  "email":"updatedemail@gmail.com",
  "password":"updatedPass"
}]
 ```
for 400:

```json
{"error": "Bad request"}
```

for 500:
```json
{"error": "server error"}
```

### Test

I used postman and also golang testing libary to test these rest APIs

you can run test by typing:

go test -v

## Conclusion

We have successfully implemented the user management service.

note: I was going to use chi router to get id from url but I couldn't get it to work for testing library. So I used query params instead.



