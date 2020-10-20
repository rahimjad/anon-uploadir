## Steps to get the API go-ing locally:

1. Ensure you have postgres installed.
1. Create postgres db, in your CLI select your prefered user and run the following:
    `CREATE DATABASE anon_uploadir`.
1. Ensure you have golang installed, please see the [official docs]('https://golang.org/doc/install?download=go1.15.3.darwin-amd64.pkg').
1. You will need a YAML file to load configuration, please look at the [example config file](api/config/development.yml.example). **This file must be located in `config/{ENV}.yml`**
1. Run the API using `cd api && go run main.go`
    - You can also build the bin if preferred using `go build main.go` 
    - This will create an `api` bin file which you can execute
    - You may also supply `ENV` variable to define how to boot the app. Currently only `development` is supported.
1. API code will be served on `localhost:8080` (unless you change ports in the YAML)

---
## Steps to get the Client running locally:

I chose to use [react-boilerplate](https://github.com/react-boilerplate/react-boilerplate) for the client. Mostly because I have not tried it before and I've heard good things about the dev experience. Start up instructions as per the original README below:

1.  Make sure that you have Node.js v8.15.1 and npm v5 or above installed.
1.  Run `npm run setup` in order to install dependencies and clean the git repo.
1.  Run `npm start` to get the server going
1.  Navigate to `http://localhost:3000` to get going!
