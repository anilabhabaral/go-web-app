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
- See all the resources present in the namespace/project after the above s2i step:
```
$ oc get all
                      
NAME                        READY   STATUS      RESTARTS   AGE
pod/test-1-build            0/1     Completed   0          2m39s
pod/test-6bb97746cb-xlwnm   1/1     Running     0          112s

NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/test   1/1     1            1           2m41s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/test-649cd496cc   0         0         0       2m41s
replicaset.apps/test-6bb97746cb   1         1         1       114s

NAME                                  TYPE     FROM   LATEST
buildconfig.build.openshift.io/test   Source   Git    1

NAME                              TYPE     FROM          STATUS     STARTED         DURATION
build.build.openshift.io/test-1   Source   Git@b59f854   Complete   2 minutes ago   50s

NAME                                  IMAGE REPOSITORY                                                                                      TAGS     UPDATED
imagestream.image.openshift.io/test   default-route-openshift-image-registry.apps.xxxx.xxxx.xxxx.lab.eng.rdu2.redhat.com/abaral-go/test   latest   About a minute ago

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
curl <ROUTE_URL>/<PATH>
```
