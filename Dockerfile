# https://docs.docker.com/guides/golang/build-images/ 
FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./ # copy dependency files
RUN go mod download # install dependency files 

COPY *.go ./ # copy source code into image
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

CMD ["/docker-gs-ping"]

