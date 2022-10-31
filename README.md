# Actuator Service

In the modern world, we need to leverage the new technologies that come through. Most of the APIs we see are JSON loaded and use REST interface. However, when performance is of utmost importance, the gRPC and Protobufs have been proved to be better.

The article will give the audience an understanding of how gRPC and Protobufs work by creating an Actuator service in Golang. 

## Design

1. Engine Service - The Maintainer of the client state (an end user) and clients will interact with engine. The maintainer will send the ideal state using REST to the engine. The engine will update the state.
2. Clients - The clients will interact with Engine using gRPC service and defined contracts in form of Protobufs. Upon getting the desired state, the client will incorporate it.

We need the interaction between service and engine to be highly performant and this is the reason for using Protobufs and gRPC. We need the maintainer interaction to be user friendly and so, REST will be used for that case.

Initial design will be simple request response in gRPC. We can modify the interaction to be bi-directional streams in the further releases.

### Request Format
``` curl http://localhost:9090/variable --request POST --data @body.json  -H "Content-Type: application/json" ```

## Nirvana (Expectations from the Project)
Create generic designs for client and engine for database interactions and database state handling. Deploy the engine service on Cloud and create client packages for databases like MySQL, PostgreSQL databases etc. to maintain custom deployments on cloud. State of engine should be stored in a reliable key value pair like Redis/minio to be durable.