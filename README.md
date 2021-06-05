# Sample Application
- Cloud Spanner
- Cloud Function
- Go Lang

### start docker
```
$ sudo docker-compose up
$ sudo docker-compose exec golang bash
```

### start functon server
```
$ go run cmd/function/test_server.go
```

### insert initial spanner data
```
$ go run cmd/spanner/init_table.go
```

### in spanner container
```
$ sudo docker-compose exec spanner-cli spanner-cli -p test-project -i test-instance -d test-database
```
edit your .env file

### curl command
```
$ curl localhost:8080/list
$ curl localhost:8080/item?property_id=100001
```
