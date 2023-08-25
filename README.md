Cool Music App!

Getting started!

Requirements 

1 - Have a spotify development API KEY
2 - Have a SQL Server instance running either on you env or the cloud to connect
3 - Have the lastest version of golang installed
  
In order to run the app follow the next steps!

1) Clone the repository
``` bash
git clone https://gitlab.com/vozy/go-vozyengine-hangup-manager.git
```


3) Get and compile dependencies
``` bash
go mod tidy
```

5) Set your env keys on a .env file using the following key names
``` 
SPOTIFY_CLIENT_ID="<YOUR SPOTIFY CLIENT ID>"
SPOTIFY_CLIENT_KEY="<YOUR SPOTIFY CLIENT SECRET KEY>"
APP_PORT=8080
``` 

7) Run the project!, or build it and run the compiled!

``` bash
go run main.go
```
``` bash
go build
```
``` bash 
./project-name
```




Tests

1) To run tests go to, /test/unit and run the default test command
``` bash
go test  || go test -run ''
```


🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎
You can check the API documentation here: sample.txt
You can explore the Postman collection here: https://documenter.getpostman.com/view/3678249/2s9Y5WyPUr
