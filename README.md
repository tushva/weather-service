# weather-service
A simple service to retrieve today's forecast and how it feels like. To use this service, execute 

`go run cmd/forecast/main.go`

To get today's short forecast along with how it feels like, execute 

`curl --location 'localhost:8080/forecast/{x},{y}'`

where `x` and `y` are integers making up Grid Points. If you want to execute using Docker, then clone this repository and execute the following:

`docker build -t weather-service/latest .`

`docker run -p 8080:8080 weather-service:latest`