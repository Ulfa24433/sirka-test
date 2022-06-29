# Documentation

## 1. Clone or extract file

extract this zip file into your computer

## 2. Installation

setup your local machine by install 

- Docker 

## 3. Setup Envvar

in your local machine, make sure that .env file is already created. this is to load certain values for database connection.

```
  RDS_POSTGRES_HOST=localhost
  RDS_POSTGRES_PORT=5432
  RDS_POSTGRES_DATABASE=test_case
  RDS_POSTGRES_USERNAME=root
  RDS_POSTGRES_PASSWORD=secret
  RDS_POSTGRES_SSL_CERT=disable
  TIMEZONE=Asia/Indonesia
  PORT=50050
```
## 4. Run Command

once everyhing is done please execute below command: 
`docker-compose up` this command is to create 2 container for postgres and golang

## 5. Test Program

- from your browser or postman, please run below endpoint
    - ```localhost:50050/v1/user/displayUser`` to display user by input userid in body request.

      sample curl to display user

      ```
        curl --location --request POST 'localhost:50050/v1/user/displayUser' \
        --header 'Content-Type: application/json' \
        --data-raw '{
          "userid": "01"
          }'

      ```

    - ```localhost:50050/v1/user/displayAllUser``` to display all user

      sample curl to display all user

      ```
        ccurl --location --request GET 'localhost:50050/v1/user/displayAllUser' \
        --data-raw ''

      ```





