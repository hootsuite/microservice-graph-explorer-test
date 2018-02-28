
- [Introduction](#introduction)
- [Running the Services](#running-the-services)
- [Exploring the Service Graph](#exploring-the-service-graph)
- [Advanced Testing](#advanced-testing)
- [Stopping the Services](#stopping-the-services)

# Introduction

This project uses [Docker](https://docs.docker.com/install/) and [Docker Compose](https://docs.docker.com/compose/install/) to build 
and run test services that can be used to demo the [Health Checks API](https://github.com/hootsuite/health-checks-api) and 
[Microservice Graph Explorer](https://github.com/hootsuite/microservice-graph-explorer).


# Running the Services

### Install the Build Tools
To build and run the test services, you must install the following:

#### Docker
Install Docker from [https://docs.docker.com/install/](https://docs.docker.com/install/). 

#### Docker Compose
Install Docker Compose from [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/).

### Run the Services
Build and run the services using Docker compose by running the follwing command:
```ssh
docker-compose up -d --build
```

### Check Services are Running
Once you have the services running locally in Docker, you can do a sanity check to make sure each service is running by 
loading the `/status/about` endpoint for [demo app](http://localhost:8080/status/about), [service 1](http://localhost:8081/status/about) 
and [service 2](http://localhost:8081/status/about).


# Exploring the Service Graph
If you haven't already, head on over the [Microservice Graph Explorer](https://github.com/hootsuite/microservice-graph-explorer) 
project in github and follow the instructions to [Build](https://github.com/hootsuite/microservice-graph-explorer#install-build-tools) 
and [Run the app](https://github.com/hootsuite/microservice-graph-explorer#running-the-app).

Once the Microservice Graph Explorer app is running, load [http://localhost:9000](http://localhost:9000) in your favorite 
browser. At this point you should see the homepage of the Microservice Graph Explorer with a link to `Test Service Graph` 
under `Quick Links`.

![Microservice Graph Explorer Homepage](/img/Microservice-Graph-Explorer-Test-Homepage.png?raw=true "Microservice Graph Explorer Homepage")

To start exploring, click on `Test Service Graph` and you'll see a page like:
[![For more info, watch the demo video](/img/microservice-graph-explorer.png?raw=true "Microservice Graph Explorer Dashboard")](https://youtu.be/JAoSkddOIC8?t=25m29s)
[For more info, watch the demo video](https://youtu.be/JAoSkddOIC8?t=25m29s)

# Advanced Testing
If you want to simulate other failure types, like a warning or a cascading failure, you can do this by modifying the `local.yml` conf file 
for a service and then rebuild/deploy the services using Docker Compose.

### Example - A Cascading Failure
To make the `Mongo` dependency of `Service 1` fail, open `docker/service-1/conf/local.yml` and update the `MONGO` item under 
`checks` from `result: "OK"` to `result: "CRIT"` and enter an error message like `details: "Can't connect"` and then run:

```ssh
docker-compose up -d --build
```

To make `Mongo` dependency fail randomly, change the value `randomlyfail: "false"` to ` randomlyfail: "true"` and then 
build/deploy again using:

```ssh
docker-compose up -d --build
```

# Stopping the Services
To stop the services running on your machine, run:
```ssh
docker-compose down --rmi all
```

TEST