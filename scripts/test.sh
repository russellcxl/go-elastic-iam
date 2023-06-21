# create videos
curl -X POST -H "Content-Type: application/json" -H "authorization: Basic dXNlcjE6MTIz" -d '{"title":"Power Rangers","description": "Go Go Power Rangers!", "url":"https://power-rangers.com","author": {"name": "user1","email": "user1@gmail.com"}}' http://localhost:8080/api/save
curl -X POST -H "Content-Type: application/json" -H "authorization: Basic dXNlcjE6MTIz" -d '{"title":"Power Rangers","description": "Go Go Power Rangers!", "url":"https://power-rangers.com","author": {"name": "user1","email": "user1@gmail.com"}}' http://localhost:8080/api/save
curl -X POST -H "Content-Type: application/json" -H "authorization: Basic dXNlcjE6MTIz" -d '{"title":"Power Rangers","description": "Go Go Power Rangers!", "url":"https://power-rangers.com","author": {"name": "user1","email": "user1@gmail.com"}}' http://localhost:8080/api/save

# create author
curl -X POST -H "Content-Type: application/json" -H "authorization: Basic dXNlcjE6MTIz" -d '{"name": "Andy","email": "andy@gmail.com"}' http://localhost:8080/api/author

# get author
curl -X GET -H "Content-Type: application/json" -H "authorization: Basic dXNlcjE6MTIz" -d '{"id": "1"}' http://localhost:8080/api/author