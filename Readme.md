**This is a HTTP service that exposes an endpoint "/numbers". This endpoint receives a list of URLs through a "GET" query
parameter and return a slice of numbers as a response**

To start the application, run the following command in your terminal:

```bash 
make up
```

You can also start the application by directly running the following command in your terminal:

```bash
docker-compose up 
```
This will build the application and start the server. The server will be available at http://localhost:8080

You can also check the api doc after running the swagger on :

```copy
http://localhost:8080/swagger/index.html
```
 
Requests can be made to the server using the following curl command:

```copy
curl --location 'http://localhost:8080/numbers?u=http%3A%2F%2Fservers%3A8090%2Fodd&u=http%3A%2F%2Fservers%3A8090%2Fprimes'
```

To stop the application, run the following command in your terminal:

```bash
make stop
```
