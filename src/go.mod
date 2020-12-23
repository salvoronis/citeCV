module main

go 1.15

replace config => ./config

replace handlers => ./handlers

replace root => ./handlers/root

replace auth => ./handlers/auth

replace databaseutils => ./database

replace models => ./models

require (
	auth v0.0.0
	config v0.0.0
	databaseutils v0.0.0
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lib/pq v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	handlers v0.0.0
	root v0.0.0
	models v0.0.0
)
