# salaries

Basic golang api to handle salaries and stats

Features to add
- Improve Auth Service and implement Auth Repository to verify tokens against database instead of using dummy user 
- Convert constants to environment variables and create a handler to get them
- Add swagger for api documentation

How to use

Public endpoint to get token
- First we need to get a token calling `auth/login` with my dummy user (`pmagnaghi` && `123456`) 
- Add that access_token as Authorization Bearer `access_token` for protected endpoints in `api/salaries`
````
curl --location --request POST 'http://localhost:8080/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "pmagnaghi",
    "password": "123456"
}'
````

Protected endpoints

- Create salary
```
curl --location --request POST 'http://localhost:8080/api/salaries' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2NzAzMjQ0MDIsImlhdCI6MTY3MDMyMDgwMiwiaWQiOjF9.SDJ3PVWtWsIIRX59xjepRzNOznITyZbqCIJAAJXmbKE' \
--header 'Content-Type: application/json' \
--data-raw '  {
    "name": "Anurag",
    "salary": "90000",
    "currency": "USD",
    "department": "Banking",
    "on_contract": "true",
    "sub_department": "Loan"
  }'
```
- Get all salaries
```
curl --location --request GET 'http://localhost:8080/api/salaries' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2NzAzMjQ0MDIsImlhdCI6MTY3MDMyMDgwMiwiaWQiOjF9.SDJ3PVWtWsIIRX59xjepRzNOznITyZbqCIJAAJXmbKE'
```
- Delete salary by id
```
curl --location --request DELETE 'http://localhost:8080/api/salaries/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2NzAzMjQ0MDIsImlhdCI6MTY3MDMyMDgwMiwiaWQiOjF9.SDJ3PVWtWsIIRX59xjepRzNOznITyZbqCIJAAJXmbKE'
```
- Get stats for entire datasets
```
curl --location --request GET 'http://localhost:8080/api/salaries/stats' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2NzAzMjQ0MDIsImlhdCI6MTY3MDMyMDgwMiwiaWQiOjF9.SDJ3PVWtWsIIRX59xjepRzNOznITyZbqCIJAAJXmbKE'
```
- Get stats for contracts
```
curl --location --request GET 'http://localhost:8080/api/salaries/stats/contracts' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2NzAzMjQ0MDIsImlhdCI6MTY3MDMyMDgwMiwiaWQiOjF9.SDJ3PVWtWsIIRX59xjepRzNOznITyZbqCIJAAJXmbKE'
```
- Get stats for departments
```
curl --location --request GET 'http://localhost:8080/api/salaries/stats/departments' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2NzAzMjQ0MDIsImlhdCI6MTY3MDMyMDgwMiwiaWQiOjF9.SDJ3PVWtWsIIRX59xjepRzNOznITyZbqCIJAAJXmbKE'
```
- Get stats for subDepartments
```
curl --location --request GET 'http://localhost:8080/api/salaries/stats/sub-departments' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2NzAzMjQ0MDIsImlhdCI6MTY3MDMyMDgwMiwiaWQiOjF9.SDJ3PVWtWsIIRX59xjepRzNOznITyZbqCIJAAJXmbKE'```
```

Run services
```
make up
```

Clean container in case of error
- There is a [bug](https://github.com/docker/compose-cli/issues/1537) in docker
- If you see it after run services, you should run
```
make clean
```


Stop services
```
make down
```

Testing container
```
make test-container
```

Run locally
```
make run
```

Testing locally
```
make test-locally
```

Final comments
- I left an initializer to show how I created the database from the dataset, this is only to test faster the endpoints
- To run (You do not need to run because salaries.db was created, it just to show how to run)
```
make initializer-dataset
```