# Application for receiving transaction information from a CSV file

## The REST API has the following endpointers:
- Downloading transactions from a *.csv file (example.csv), parsing it and saving the parsing results to the database
- Filtering and uploading previously saved data in JSON format in the response
- The same as endpoint 2, but only download data not in JSON, but in the form of a CSV file.
- Get transactions

## Prerequisites
- go 1.20
- docker & docker-compose
- <a href="https://github.com/swaggo/gin-swagger">swag</a> 

## Run Project
Create .env file in root directory and add following values:
```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=postgres
DB_SSLMODE=disable
DB_USERNAME=postgres
DB_PASSWORD=postgres

HTTP_HOST=localhost
HTTP_PORT=8080
```
Definition migrating to database

```make migrate-up```

Use `make run` to build&run project

```make run``

Swagger

<a href="http://localhost:8080/swagger/index.html">http://localhost:8080/swagger/index.html</a>

![swagger-image](../main/assets/swagger-image.png)

#### ALL TRANSACTIONS IS A FAKE