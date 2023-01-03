# Companies API

This a REST API Microservices that handles Company entity with following attributes:
* Name
* Code
* Country
* Website
* Phone
For read operation, fetching one or many companies is available.
Each Company’s attribute is available as filtering in the CRUD operations.

For resiliency the outbox pattern was implemented into this app. Each create,update,delete operation is saved into "outbox" table.
CDC (Change Data Capture) is implemented using the Debezium application which monitors outbox table and pushes events to topics:
```
outbox.event.companies-create
outbox.event.companies-delete
outbox.event.companies-update
```
## Getting Started
You can find the swagger.yml file with the detailed description of an API in api/ directory

## Running

To run an app using your local environment:

1. build the application using makefile commands (build/build_mac/build_windows)
2. inject env variables using example.env file as an example of required variables (for Goland see goland-env.txt)
3. use command ``./companies migrate up`` to run migrations

To run the application, execute command below in your k8s cluster:
```
kubectl apply -f k8s/namespaces/ns.yaml

kubectl apply -f k8s/postgres/postgres-configmap.yaml
kubectl apply -f k8s/postgres/postgres-storage.yaml  
kubectl wait --for=jsonpath='{.status.phase}'=Bound pvc/postgres-pv-claim --namespace companies
kubectl apply -f k8s/postgres/postgres-deployment.yaml
kubectl apply -f k8s/postgres/postgres-service.yaml 

helm repo add strimzi https://strimzi.io/charts/
helm install strimzi-kafka strimzi/strimzi-kafka-operator --namespace companies

kubectl apply -f k8s/kafka/kafka-deployment.yaml

kubectl apply -f k8s/debezium/debezium-configmap.yaml
kubectl apply -f k8s/debezium/debezium-deployment.yaml
kubectl wait deployments/debezium --for=condition=Available=True --namespace companies
kubectl apply -f k8s/debezium/debezium-service.yaml
kubectl apply -f k8s/debezium/debezium-job.yaml

kubectl apply -f k8s/companies/companies-configmap.yaml
kubectl apply -f k8s/companies/companies-deployment.yaml
kubectl apply -f k8s/companies/companies-service.yaml 
```

App exposes 3000 port by default
