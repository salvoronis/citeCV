module main

go 1.15

replace utils => ./src/utils

replace config => ./src/config

replace databaseutils => ./src/database

replace controllers => ./src/controllers

replace models => ./src/models

require (
	config v0.0.0
	controllers v0.0.0
	databaseutils v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lib/pq v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	models v0.0.0
	utils v0.0.0
)
