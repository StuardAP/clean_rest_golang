# Clean REST API with Go and MySQL

## Description

This project is a REST API built with Go and MySQL, specifically designed to handle the operations of an ecommerce. It follows the clean architecture approach, allowing a well structured and easy to maintain code. This API provides endpoints to manage users, products and inventory, and focuses on scalability and modularity.

## Project Structure The project is organized as follows

```markdown
└── 📁clean_rest_golang
    └── 📁cmd
        └── 📁api
            └── main.go
    └── 📁infrastructure
        └── 📁database
            └── 📁queries
                └── 📁product
                    └── product_repository_query.go
                └── 📁user
                    └── user_repository_query.go
        └── 📁http
            └── 📁handlers
                └── 📁product
                    └── product_hander.go
                └── 📁user
                    └── user_handler.go
            └── 📁middleware
            └── 📁routes
                └── product_route.go
                └── status_route.go
                └── user_route.go
            └── router.go
    └── 📁internal
        └── 📁domain
            └── 📁entities
                └── inventory_entity.go
                └── order_entity.go
                └── product_entity.go
                └── user_entity.go
            └── 📁repositories
                └── product_repository.go
                └── user_repository.go
            └── 📁usecases
                └── 📁product
                    └── product_usacase.go
                └── 📁user
                    └── user_usecase.go
    └── 📁migrations
        └── 20241018233645_add_user_table.down.sql
        └── 20241018233645_add_user_table.up.sql
    └── 📁pkg
        └── 📁config
            └── config.go
            └── database.go
            └── jwt.go
            └── server.go
        └── 📁database
            └── 📁mysql
                └── setup_mysql.go
        └── 📁utilities
            └── 📁auth
                └── jwt.go
                └── password_encrypt.go
    └── .air.toml
    └── .dockerignore
    └── .env
    └── .env.example
    └── .gitignore
    └── docker-compose.yml
    └── Dockerfile
    └── go.mod
    └── go.sum
    └── Makefile
    └── README.md
    └── script.sh
```
