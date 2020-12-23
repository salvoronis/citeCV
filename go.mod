module main

go 1.15

replace config => ./src/config

replace handlers => ./src/handlers

replace root => ./src/handlers/root

replace auth => ./src/handlers/auth

replace databaseutils => ./src/database

replace models => ./src/models

replace utils => ./src/utils

require (
	auth v0.0.0
	config v0.0.0
	databaseutils v0.0.0
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/lib/pq v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	handlers v0.0.0
	root v0.0.0
	models v0.0.0
	utils v0.0.0
)
