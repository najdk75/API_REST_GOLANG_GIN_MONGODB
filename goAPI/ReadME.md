# Simple GO Lang REST API

> Simple REST API in Golang with Gin & MongoDB 


## Quick Start


Installation
```bash
$ go get -u github.com/gin-gonic/gin
$ go get go.mongodb.org/mongo-driver/mongo 
$ go get .
$ go build
$ ./go_api
```

## API Endpoints

- http://localhost:8080/create/users
    - `POST`: insert user(s) into the database
- http://localhost:8080/users/list
    - `GET`: get all the users
- http://localhost:8080/user/:id
    - `GET`: get a user 
- http://localhost:8080/delete/user/:id
    - `DELETE`: delete a user 
- http://localhost:8080/delete/users
    - `DELETE`: delete all the users
- http://localhost:8080/update/user/:id
    - `PUT`: update a user
- http://localhost:8080/login
    - `POST`: get all the information of a user after a successful log in




## Data Structure
`/create` &&`update/user/:id` : Use the following data structure if you want to insert a single user (you can check the all the fields in the models.user.go file)
```json
{
"_id": "1t5VsIBXpGl4s8",
"password" :"$2a$10$oGNG9nb5AM.3Ht5tHxhKIe9vo4kguhV9wAwkFzrzhbyZrWIifeXBC",
"name" :"Nikki Farley",
"balance" :"$2,547.50",
"age" :27,
"company" : "ANIVET",
"email": "nikkifarley@anivet.com",
"phone" : "+1 (868) 439-2675",
"address" :"588 Schaefer Street, Falconaire, Missouri, 9457",
"about" :"Commodo minim fugiat est fugiat sunt duis consectetur fugiat Lorem sun…",
"registered" :"2014-10-12T04:14:20 -02:00",
"latitude" : -10.647121,
"longitude" : 26.04006,
//...
}
```

`/login`: Use the following UserLogin data structure 
```json
{
  "id": "1t5VsIBXpGl4s8",
  "password": "012-345-6789",
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
