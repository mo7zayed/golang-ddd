### Golang & Nuxt JS Chat Application
This is a simple chat application using [golang](https://golang.org/) and [nuxtjs](https://nuxtjs.org/).
I have implemented DDD (Domain Driven Design) [read more](http://dddsample.sourceforge.net/architecture.html).
The files structure needs enhancement it can be better.

#### Features
- JWT Authentication
- Database migrations (Makefile)
- Database seeds (Makefile)
- Single Page Application
- Postman collection created `postman_collection.json`, `postman_environment.json`

#### Installation
##### Server Side
- `git clone repo.url.here ./folder`
- `cd folder`
- `go list -e $(go list -f '{{.Path}}' -m all)`
- `make serve` or `make dev` for development NOTE: if you are gonna run `make dev` see https://github.com/gravityblast/fresh first.
- set your vars in .env file
- run the migrations `make migrate`
- run database seeds to generate fake data `make seed`

##### Client Side
- `cd folder/client`
- see https://nuxtjs.org/guide/installation first
- npm install
- npm run dev
- have fun ...

##### Run the application
- make serve
- cd client && npm run dev