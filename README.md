# HexaShop

HexaShop is a simple Golang application built to learn and implement hexagonal architecture. It includes APIs for managing users and products, demonstrating the core principles of domain-driven design and separation of concerns.

## Features

- User management API (CRUD operations)
- Product management API (CRUD operations)
- Demonstrates hexagonal Architecture (ports & adapters pattern)

## Architecture Overview

This project uses the hexagonal architecture (also known as Ports and Adapters), which aims to create a more modular, loosely-coupled code structure. The business logic (core domain) is separated from the infrastructure (database, HTTP layer), making it easier to test and extend.

## Pre-requisites

- Must have installed `go 1.21.1` or latest.
- Must have installed [task](https://taskfile.dev/installation/) command

## How to Run

- Run the following command to execute the application

    ```bash
    task run
    ```

## Project Structure

```bash
./
├── cmd/
│   └── app/
│       └── main.go
├── conf/
│   └── conf.go
├── internal/
│   ├── adapter/
│   │   ├── http/
│   │   │   ├── api/
│   │   │   │   ├── product.go
│   │   │   │   └── user.go
│   │   │   ├── rq/
│   │   │   │   ├── product.go
│   │   │   │   └── user.go
│   │   │   └── router.go
│   │   └── storage/
│   │       ├── postgres/
│   │       │   ├── migration/
│   │       │   │   ├── 000001_init.down.sql
│   │       │   │   └── 000001_init.up.sql
│   │       │   ├── repo/
│   │       │   │   ├── product.go
│   │       │   │   └── user.go
│   │       │   ├── schema/
│   │       │   │   └── schema.sql
│   │       │   ├── db.go
│   │       │   └── migration.go
│   │       └── redis/
│   ├── core/
│   │   ├── domain/
│   │   │   ├── product.go
│   │   │   └── user.go
│   │   └── service/
│   │       ├── product.go
│   │       └── user.go
│   └── port/
│       ├── product.go
│       └── user.go
├── .env.sample
├── .gitignore
├── Taskfile.yml
├── go.mod
└── go.sum
```
