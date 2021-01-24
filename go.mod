module main

go 1.15

replace utils => ./src/utils

replace config => ./src/config

replace databaseutils => ./src/database

replace controllers => ./src/controllers

require (
	config v0.0.0
	databaseutils v0.0.0
	github.com/gorilla/mux v1.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	utils v0.0.0
	controllers v0.0.0
)
