# go http server
A simple Go http server built using Gin Framework https://gin-gonic.com/docs/

### Installation
-----------------
 - Clone repository:
  ```bash
    $ git clone https://github.com/iozerov/testtask.git
  ```
 - Run application:
  ```bash
    $ go run .
  ```

### Test
--------
To run unit tests execute this command:
```bash
$ go test -v

--- PASS: TestContains (0.00s)
=== RUN   TestDifference
--- PASS: TestDifference (0.00s)
=== RUN   TestFilter
--- PASS: TestFilter (0.00s)
=== RUN   TestDelete
--- PASS: TestDelete (0.00s)
PASS
ok      example/web-service-gin 0.009s
```

### Endpoints
-------------
You can use terminal and [Curl](https://curl.se/) to perform requests following endpoints:
 - GET `/refresh`
 Refresh all ip addresses from the resource
 ```bash
    $ curl http://localhost:8080/refresh
 ```
 - GET `/last-changes`
 Returns JSON with last changes (removed and added ip addresses)
 ```bash
    $ curl http://localhost:8080/last-changes
 ```
 - GET `/filter?query=`
  Returns list with ip addresses containing substring
 ```bash
    $ curl http://localhost:8080/filter?query=33
 ```
 - GET `/count`
 Returns count of all ip addresses
 ```bash
    $ curl http://localhost:8080/count
 ```
 - POST `/contains"`
 Returns bool value (true/false) if ip address exists
 ```bash
    $ curl http://localhost:8080/contains -H 'Content-Type: application/json' -d '{"ip_address": "74.82.204.124"}'
 ```
- DELETE `/delete`
 Delete all ip addresses
 ```bash
    $ curl -X DELETE http://localhost:8080/delete
 ```
