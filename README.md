# golang-vercel
 Golang basic template and deployment with vercel serverless environment. 

 **What Is Vercel:**
 ​Vercel is a cloud platform for static sites and Serverless Functions that fits perfectly with workflow. It enables developers to host websites and web services that deploy instantly, scale automatically, and requires no supervision, all with no configuration.

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
-   Vercel
-   Node, NPM

## Installation
-   Install node and npm using `brew install node`
-   Run `npm i -g vercel` and check if vercel install properly by `vercel --version`
-   Run `vercel dev` for development server
-   Run `vercel --prod` for production server

**Vercel Production URL**: https://golang-vercel.vercel.app/
