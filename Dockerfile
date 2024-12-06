# https://docs.docker.com/guides/golang/build-images/ 
FROM golang:1.23

WORKDIR /usr/local/go/src

RUN mkdir organizations

WORKDIR /usr/local/go/src/organizations

COPY go.mod go.sum ./
RUN go mod download 
COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

CMD ["/docker-gs-ping"]

