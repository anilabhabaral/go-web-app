# Simple Hello World webapp using go

Run the application:
```
go build .
```

```
./go-web-app
```

Access the application using:
```
curl localhost:8080/greet
```

### Deploy this app in OpenShift using s2i
- Deploy application using s2i:
```
oc new-app --name=test https://github.com/anilabhabaral/go-web-app.git
```
- Create a service for the deployment:
```
oc expose deploy test --port 8080
```
- Create a ROUTE for the above service:
```
oc expose svc test --path "/greet"
```

### Access the app deployed in OpenShift
- Get the route:
```
oc get route
```
- Access the application endpoint:
```
curl <ROUTE_URL>
```
