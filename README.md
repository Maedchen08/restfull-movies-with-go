# homework-rakamin-golang-sql

## setup
1. clone this repository https://github.com/Maedchen08/restfull-movies-with-go.git
2. type ```make run``` to run the apps
3. Hit the Server `localhost:8080/movies/titanic`


### Tasks 
We define routes for handling operations:

| Method        | Route                  | Action                                              |
|---------------|------------------------|-----------------------------------------------------|

| POST          | /movie                 | create movie                                        |
| GET           | /movie/:slug           | get movie by slug                                   |
| PUT           | /movie/:slug           | update movie by slug                                |
| DELETE        | /movie/:slug           | delete movie by slug                                |

Access API via ```http://localhost:8080/{route}```



1. POST ```/movie ```

Authorization: Bearer {token} 

Request Body: 
```
{
   "title":"boy",
   "slug":"boy",
   "description":"lorem ipsum",
   "duration": 60,
   "image":"titanic poster url"
}
```

Response:
status code : 201
```
{
    "error": false,
    "msg": "success create data",
    "result": {
        "id": 13,
        "title": "boy",
        "Slug": "boy",
        "Description": "lorem ipsum",
        "Duration": 60,
        "Image": "titanic poster url"
    }
}
```

3. GET ```/movie/boy```

Authorization: Bearer {token} 

Response:
status code: 200
```
{
    "error": false,
    "msg": "success retrieve data",
    "result": {
        "id": 13,
        "title": "boy",
        "Slug": "boy",
        "Description": "lorem ipsum",
        "Duration": 60,
        "Image": "titanic poster url"
    }
}
```

4. PUT ```/movie/boy```

Authorization: Bearer {token} 

Response:
status code: 200
```
{
   "title":"boy",
   "slug":"titanic",
   "description":"lorem ipsum",
   "duration": 60,
   "image":"titanic poster url"
}
```

5. DELETE ```/movie/titanic```

Authorization: Bearer {token} 

Response:
status code: 200
```
{
    "error": false,
    "msg": "success"
}
```

### Tech Stack
* [Golang] - programming language
* [Fiber] - web framework with zero memory allocation and performance
* [Gorm] - Library untuk mendukung penggunaan SQL

[Golang] : <https://golang.org/>
[Fiber] : <https://github.com/gofiber/fiber/>
[Gorm] : <https://gorm.io/docs/>

nb: 
to do myhomework sql and data modelling, I make 2 repository to practice alone in understanding the material with and without using biolerplate, please check both of them and give me suggestions. 

[without_biolerplate] : <https://github.com/Maedchen08/crud-movies.git>
[with_boilerplate] : <https://github.com/Maedchen08/restfull-movies-with-go.git>