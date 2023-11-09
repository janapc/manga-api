<div align="center">
  <h1>Manga API</h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/manga-api"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/manga-api"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/manga-api"/>
  <img alt="CI/CD Pipeline" src="https://github.com/janapc/manga-api/actions/workflows/tests.yml/badge.svg"/>

<a href="#project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#run-project">Run Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#request-api">Request API</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#technologies">Technologies</a>

</div>

## Project

Api to register mangas and manager those mangas.

## Requirement

To this project your need:

- golang v1.21 [Golang](https://go.dev/)
- docker [Docker](https://www.docker.com/)

In _cmd/server_ folder create a file **.env** with:

```env
DB_URL= //database connection
JWT_SECRET= // your secret
JWT_EXPIRES_IN= // your time expires
BASE_URL_V1=http://localhost:3000/api/v1 // url of api
```

## Run Project

Start Docker in your machine and run this commands in your terminal:

```sh
## up mongodb, kafka and mysql
❯ docker compose up -d

## run this command to install dependencies:
❯ go mod tidy

## run this command inside cmd/server to start api(localhost:3000):
❯ go run main.go

```

## Request API

[local-swagger](http://localhost:3000/api/v1/swagger/index.html)

## Technologies

- golang
- postgres
- docker
- gorm
- go-chi
- viper

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
