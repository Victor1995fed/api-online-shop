# API for online store
## Description:
API for online store on golang

## Quickly start:
* Build docker image (into directory with app)

    ```git clone```

     ```make init```

## Other commands:
### Postgres:
 * Connect (inside the pgsql-container) :
 
    ```psql  restapi_dev user```

 * Up migrations:
 
    ```./migrate  -path migrations -database "postgres://pgsql?dbname=somedb&user=someuser&password=somepassword&sslmode=disable" up ``` 

 * Down  all migrations: 

    ```./migrate  -path migrations -database "postgres://pgsql?dbname=somedb&user=someuser&password=somepassword&sslmode=disable"  down```  

## Useful links:

* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

* [Codeship Golang Best Practices](https://github.com/codeship/go-best-practices)