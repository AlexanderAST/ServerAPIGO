# мой учебный проект в нём я изучил технологию REST API,MongoDB, проработал ошибки приложения
# user-service
# на базе REST API 

GET /users --list of users -- 200,404 ,500 
GER /users/:id --user by id --200,404 ,500
Post /users/:id --create user --204 ,4xx Header Location: url 
PUT /users/:id -- fully update user --204/200,400 404 ,500
PATCH  /users/:id -- partially update user --204/200 400 404 ,500 
DELETE /users/:id -- delete user by id --204 400 404  