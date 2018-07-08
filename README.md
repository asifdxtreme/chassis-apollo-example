# Chassis-Apollo Example [WIP]

This is a example illustrates the integration of go-chassis with Ctrip Apollo Configuration Center.

### Objective
This examples showcases the following things:
1. Start Ctrip Apollo and make Projects, NameSpace and Configurations.
2. Configure chassis.yaml to use Apollo Configuration Server.
3. Use Apollo to dynamically change configurations of chassis.
4. Configure the QPS limit for microservices by using Apollo.
5. Configure Fault tolerance for microservices by using Apollo.


#### Quick start Apollo

Easiest way to bring up Apollo configuration is using docker compose.

```go
git clone https://github.com/ctripcorp/apollo $GOPATH/src/github.com/ctripcorp/apollo

cd $GOPATH/src/github.com/ctripcorp/apollo/scripts/docker-quick-start

docker-compose up
```
This brings up the Apollo, follow the below steps to configure the project and namespace.

- Login to Apollo 


