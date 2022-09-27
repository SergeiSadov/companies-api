# Companies API

This a REST API Microservices that handles Company entity with following attributes:
* Name
* Code
* Country
* Website
* Phone
For read operation, fetching one or many companies is available.
Each Company’s attribute is available as filtering in the CRUD operations.

## Getting Started
You can find the swagger.yml file with the detailed description of an API in api/ directory

## Running

To run an app using your local environment:

1. build the application using makefile commands (build/build_mac/build_windows)
2. inject env variables using example.env file as an example of required variables (for Goland see goland-env.txt)
3. use command ``./companies migrate up`` to run migrations

To run an app in kubernetes execute this command below (for example, in minikube):
```
kubectl apply -f k8s/postgres/postgres-configmap.yaml
kubectl apply -f k8s/postgres/postgres-storage.yaml  
kubectl apply -f k8s/postgres/postgres-deployment.yaml
kubectl apply -f k8s/postgres/postgres-service.yaml 

kubectl apply -f k8s/zookeeper/zookeeper-configmap.yaml
kubectl apply -f k8s/zookeeper/zookeeper-deployment.yaml
kubectl apply -f k8s/zookeeper/zookeeper-service.yaml 

kubectl apply -f k8s/kafka/kafka-configmap.yaml
kubectl apply -f k8s/kafka/kafka-deployment.yaml
kubectl apply -f k8s/kafka/kafka-service.yaml 

kubectl apply -f k8s/companies/companies-configmap.yaml
kubectl apply -f k8s/companies/companies-deployment.yaml
kubectl apply -f k8s/companies/companies-service.yaml 


```

App exposes 3000 port by default

### TODO List

- [ ] Add resilient kafka client. Current implementation is naive and error prone
- [ ] Add additional unit test for pkg dir
- [ ] Fix kafka for k8s config