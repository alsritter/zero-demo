
user:
	cd service/user/api && go run user.go -f etc/user-api.yaml

user-test:
	curl -i -X POST \
  http://127.0.0.1:9999/user/login \
  -H 'Content-Type: application/json' \
  -d '{ "username":"666", "password":"123456" }'

search:
	cd service/search/api && go run search.go -f etc/search-api.yaml

search-test:
	curl -i -X GET \
  'http://127.0.0.1:8888/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0' \
	-H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTI4NjcwNzQsImlhdCI6MTYxMjc4MDY3NCwidXNlcklkIjoxfQ.JKa83g9BlEW84IiCXFGwP2aSd0xF3tMnxrOzVebbt80'