# FreightCMS

## Organization Web Application

### Building Docker Image

1. `docker build --tag docker-gs-ping .`

You can tag with different tags by doing

https://docs.docker.com/guides/golang/build-images/

```sh
docker image tag docker-gs-ping:latest docker-gs-ping:v1.0
```

### Publishing Docker Image

In the terminal run the below commands to login to docker hub using the docker cli

1. Create a gpg key using `gpg --generate-key`
2. Follow the prompts to complete the necessary steps
3. Run [pass](https://www.passwordstore.org/) init . to initialize the password store on your system
4. Run docker login and follow the prompts
5. After successfully login in you can run `docker build . -t squishedfox/organization-web-service
6. Publish the new docker image with `squishedfox/organization-web-service:latest`

### Deploying

TODO
