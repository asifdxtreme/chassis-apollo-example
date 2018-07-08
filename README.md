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
![Apollo Login](images/Login-Apollo.png)
By default the credentials is user is apollo/admin

- Create a new Project
![](images/NewProject.png)
Create a new project called ***apollo-integration*** which we will use through out this example.
![](images/Apollo-Integration.png)

- Note all the details regarding Project Name, AppID, NameSpace, ClusterName.
![](images/Project-Page.png)

- Add a new Configurations
![](images/AddNewConfig.png)
Once config is added you need to release the config so that it will be available to outside components.
![](images/ReleaseConfig.png)

#### Configure the chassis.yaml
You need to configure the chassis.yaml for both the client and server to use Apollo as the configurations server. (By default we have added these configurations to this example in client and server).
```go
cse:
  config:
    client:
      serverUri: http://127.0.0.1:8080          # This should be the address of your Apollo Server
      type: apollo                              # The type should be Apollo
      refreshMode: 1                            # Refresh Mode should be 1 so that Chassis-pulls the Configuration periodically
      refreshInterval: 10                       # Chassis retrives the configurations from Apollo at this interval
      serviceName: apollo-integration           # This the name of the project in Apollo Server
      env: DEV                                  # This is the name of environment to which configurations belong in Apollo
      cluster: apollo                           # This is the name of cluster to which your Project belongs in Apollo
      namespace: application                    # This is the NameSpace to which your configurations belong in the project.
```

#### Running the Client and Server 
Client and Server in this example can be run on the same machine
- Start the Service-Center
```go
docker run -d -p 30100:30100 servicecomb/service-center
```  
- Start the Client
```go
cd client
go build
./client
```
- Start the Server
```go
cd server
go build
./server
```
- Verify the Communication between client and server
```go

```

