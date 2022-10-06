# golang-vercel
 Golang basic template and deployment with vercel 

[![Actions Status](https://github.com/teachmind/Auth-service/workflows/build/badge.svg)](https://github.com/teachmind/Auth-service/actions)

## Features 
-   Server Preparation for Running the project on localhost
-   Database Migration
-   Registration
-   Login
-   Token Validation


## Project Structure
    .
    |-- api                 # This is the default folder for vercel as entrypoint 
    |-- app                 # Contain services and features like handler, database etc
    |-- docs                # Contain swagger files
    |-- readme.md           # Explains project installation and other informations

## Tools and Technology
-   Golang
-   PostgreSQL

## Installation
-   **Step-1:** Copy/rename `.env.example` file as `.env`. Change the `APP_PORT`, `DB_PORT`, `DB_NAME`,`DB_HOST`, `DB_USER`, `DB_PASSWORD` value as per your DB and Project setup.
-   **Step-2:** Run migration command `go run main.go migrate` for Database migration
-   **Step-3:** To start server run `go run main.go server`
