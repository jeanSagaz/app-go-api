What is the app-go-api Project?
=====================
The app-go-api Project is a open-source project written in Golang

## Give a Star! :star:
If you liked the project or if app-go-api helped you, please give a star ;)

## How to use:
You will need the latest Visual Studio Code.  

Also you can run the app-go-api Project in Visual Studio Code (Windows, Linux or MacOS).

To know more about how to setup your enviroment visit the [go dev Download Guide](https://go.dev/learn/)

Execute the following command in the terminal to run project:  
```
go run .\cmd\server\main.go
```

Execute the following command in the terminal to run domain tests:  
```
go test .\internal\customer\domain\entity
```

Execute the following command in the terminal to run repository tests:  
```
go test .\internal\customer\infra\database
```

Execute the following command in the terminal to run docker:  
```
docker-compose up -d
```

## Technologies implemented:

- go 1.20
 - Routers with chi, gin and mux
 - DI
 - gorm
 - generics
 - validator
 - testing

## Architecture:

- Full architecture with responsibility separation concerns, SOLID and Clean Code
- Domain Driven Design (Layers and Domain Model Pattern)
- Repository
