Cool Music App! - Challenge application for spotify API

1# Description

The application connects to the spotify developer API in order to extract metadata via a unic spotify ID
related to a song, afterwards the application stores it on a sqlite database for further usage

the application is divided on 3 routes

1) Post - ISRC: The route gets data from the api and saves on the local database

2) GET By ISRC: Get perfect match on the local database

3) GET LIKE: Get any matches by song name on the local database



2# Requirements 

1 - Have a spotify development API KEY
2 - Have a SQL Server instance running either on you env or the cloud to connect
3 - Have the lastest version of golang installed
  
In order to run the app follow the next steps

1) Clone the repository
``` bash
git clone https://gitlab.com/vozy/go-vozyengine-hangup-manager.git
```


3) Get and compile dependencies
``` bash
go mod tidy
```

5) Set your env keys on a .env file using the following key names (the .env file must be on your project root / also in your test/unit folder if you want to run tests properly)

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
./CoolMusicApp
```


Tests

1) To run tests go to, /test/unit and run the default test command
``` bash
cd ./test/unit
```

``` bash
go test  || go test -run ''
```


🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎🔎

You can explore the Postman collection && documentation here: https://documenter.getpostman.com/view/3678249/2s9Y5WyPUr

# Author

Andrew Alizaga

# Main dependencies 

Golang Version v1.20
Gin Version v1.9.1
Godotenv Version v1.5.1
GORM / SQLITE Version v1.5.3
zmb3 - spotify library Version v1.3.0 

