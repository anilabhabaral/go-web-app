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
```
oc new-app --name=test https://github.com/anilabhabaral/go-web-app.git
```

```
oc expose deploy test --port 8080
```

```
oc expose svc test --path "/greet"
```
