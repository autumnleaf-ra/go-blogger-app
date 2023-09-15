# Blog Site Backend API

This is the backend REST API for a blog site, developed using Go and the Fiber web framework.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Deploy-API](#deploy-api)

## Introduction

This repository contains the source code for the backend of a blog site. It provides the necessary endpoints and functionality to manage blog posts, users, and other related data.

## Features

- Create, read, update, and delete blog posts.
- User authentication and authorization.
- CRUD operations for user accounts.
- Secure and efficient handling of data using the Fiber framework.
- RESTful API design.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/doc/install) installed on your machine.
- [Fiber](https://docs.gofiber.io/gofiber/getting-started) framework installed.
- Database system (e.g., PostgreSQL, MySQL) set up and configured.
- Configuration files (e.g., `.env`) with your database credentials and other sensitive information.

## Getting Started

1. Clone this repository:

   ```bash
   git clone https://github.com/autumnleaf-ra/go-blogger-app.git

2. Navigate to the project directory:
   ```bash
   cd blog-site-backend
3. Install dependencies:
   ```bash
   go mod tidy
4. Set up your configuration by creating an .env file with your environment-specific values.
5. Run the application
   ```bash
   go run main.go
   
## Usage
Visit http://localhost:your-port in your web browser or use a tool like Postman to interact with the API endpoints.

## Deploy-API
```bash
https://go-blogsite-api-ec193e4abeb7.herokuapp.com/
