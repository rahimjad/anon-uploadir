Steps to get the API going locally:
1. Ensure you have postgres installed.
1. Create postgres db, in your CLI select your prefered user and run the following:
    `CREATE DATABASE anon_uploadir`.
1. Ensure you have golang installed, please see the [official docs]('https://golang.org/doc/install?download=go1.15.3.darwin-amd64.pkg').
1. You will need a YAML file to load configuration, please look at the [example config file](api/config/development.yml.example)
1. Run the API using `cd api && go run main.go`
    - You can also build the bin if preferred using `go build main.go` 
    - This will create an `api` bin file which you can execute
1. API code will be served on `localhost:8080`