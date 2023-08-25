Cool Music App!

Getting started!

Requirements 

1 - Have a spotify development API KEY
2 - Have a SQL Server instance running either on you env or the cloud to connect
3 - Have the lastest version of golang installed
  
In order to use the app follow the next steps!

1) Clone the repository
git clone https://gitlab.com/vozy/go-vozyengine-hangup-manager.git

2) Get and compile dependencies
go mod tidy

3) Set your env keys on a .env file using the following key names
SPOTIFY_CLIENT_ID=<YOUR SPOTIFY CLIENT ID>
SPOTIFY_CLIENT_KEY=<YOUR SPOTIFY CLIENT SECRET KEY>
APP_PORT=8080


4) Run the project!
go run main.go


//todo:
You can check the API documentation here: sample.txt
You can explore the postman collection here: https://documenter.getpostman.com/view/3678249/2s9Y5WyPUr