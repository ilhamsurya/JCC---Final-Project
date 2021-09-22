# Jabar Coding Camp Final Project

## Movie Review API (IMDB Benchmark)

- Problem Description :
  - [ ] Get movie list that contain (title, description, duration, genre, etc)
  - [ ] Post a movie that contain the same attributes
  - [ ] Update a movie that contain the same attributes
  - [ ] Delete a movie from database
  - [ ] Basic authentication (Login & Register) as a writer and guest
  - [ ] As a writer, i want to write a review about movie
  - [ ] As a guest, i want to give a rating

## Prerequisites to run the application

- Development

  - Golang [installation](https://golang.org/doc/install)
  - MySQL :

    ```sh
    # My SQL Instalation:

    $ go get github.com/julienschmidt/httprouter

    $ go get github.com/go-sql-driver/mysql
    ```

  - Go Mod :

    ```sh
    # Go Mod Initialization:

    $ go mod init example.com/projectName


    ```

  - Gin-Gonic :

    ```sh
    # Gin gonic instalation:

    $ go get -u github.com/gin-gonic/gin

    $ import "github.com/gin-gonic/gin"
    ```

  - Gorm :

    ```sh
    # Gin gonic instalation:

    $ go get -u gorm.io/gorm

    $ import (
        "gorm.io/gorm"
      )
    ```

- Deployment
