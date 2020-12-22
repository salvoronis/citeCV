module main

go 1.15

replace config => ./config

replace handlers => ./handlers

replace root => ./handlers/root

require (
	config v0.0.0
	github.com/gorilla/mux v1.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	handlers v0.0.0
	root v0.0.0
)
