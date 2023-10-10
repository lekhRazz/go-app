# go-application
## stack used
1. go as main text type 
2. echo api framework
3. PostgreSQL
4. Elasticsearch for storing application logs 
5. Docker for containerizing

## What this app offers and ensures
1. redix tree like advace and fastest api routing
2. stores application logs in elasticsearch for advance log analysis
3. well and good panic recovery
4. code quality checked with sonarqube 
5. incorporates 12 factor app(software) principles
6. psql custom enum types are introduced
7. psql trigger function is used to archive deleted data

## Steps to run this application with docker

1. clone the repo with ssh link
```
git@github.com:lekhRazz/sample_go_application.git
```
2. build docker image with these command
```
docker build -t sample-go-app .
``` 
3. run docker image in container
```
docker run -d -p 3000:3000 --name sample-go-app sample-go-app
```
4. watch server log
```
 docker logs -f sample-go-app
```