# webservice-prototype

A prototype building out with GraphQL, Golang, and MongoDB

## Setup

### Go work

```txt
go 1.22.5

use (
    .
    ./db
    ./db/mongodb
    ./models
    ./web
)
```

### .env file

**On Windows**

```dotenv
echo MONGO_SERVER=mongodb://root:example@0.0.0.0:27017/ > .env
```

**on linux/unix**

```dotenv
echo "MONGO_SERVER=mongodb://root:example@0.0.0.0:27017/" >> .env
```

## Running

From the root of the application...

1. Start Docker mongo database instance, and mongo express instance

```sh
docker compose up -d
```
2. Make sure to install all dependencies

```sh
go get
```

3. Start the main application

```sh
go run main.go
```

4. test the application is up and healthy

```sh
curl http://localhost:8080/
```

should output

```sh
{ "status": "ok" }
```

## Building Docker Image

1. `docker build --tag docker-gs-ping .`

You can tag with different tags by doing

https://docs.docker.com/guides/golang/build-images/

```sh
docker image tag docker-gs-ping:latest docker-gs-ping:v1.0
```

## Testing

### Scenarios

#### Creating an organization resource


```GraphQL
mutation {
  createOrganization(dba: "first org", name: "first org", mailingAddress: {
    line1: "123 Fake St.",
    line2: "AB Section",
    line3: "Floor 3",
    region: "Minnesota",
    locale: "Minneapolis",
    countryCode: "US",
    postalCode: "44444"
  }, billingAddress: {
       line1: "123 Fake St.",
    line2: "AB Section",
    line3: "Floor 3",
    region: "Minnesota",
    locale: "Minneapolis",
    countryCode: "US",
    postalCode: "44444"
  }) {
    id
  }
}
```
