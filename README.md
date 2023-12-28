  
# Quick-go project

I chose to use quick-go package and implemented required solution. There is few way to run project.
## How to run services
To run services you need run programs in separate terminal tabs.

 1. Run server ```❯ go run main.go```
 2. Run subscriber demo  ```❯ go run ./clientSubscriber/main.go```
 3. Run publisher demo ```❯ go run ./clientSubscriber/main.go```

You may run several subscribers/publishers if you required. All program output will be provided to terminal.

## All in one example.
❯ go run ./example/main.go
To make it easier to see how server, publishers and subscribers communicate is made all in mini app.

It will launch server, 2 publishers and 2 subscribers. All communication via quick results will be outputted to *os.Stdout* which is visible in terminal.
Expected output will be somethig like this
```
[ Pub__1 ] 12:10:06 Pub__1 Is launched
[ Pub__2 ] 12:10:06 Pub__2 Is launched
[ Server ] 12:10:06 [ Publisher ] Listening on port 0.0.0.0:6002
[ Server ] 12:10:06 [ Subscriber ] Listening on port 0.0.0.0:6001
[ Server ] 12:10:07 New publisher connected [::1]:39812
[ Server ] 12:10:07 New publisher connected [::1]:41194
[ Pub__1 ] 12:10:10 <-- Hello from app at 1703758210
[ Pub__1 ] 12:10:10 --> [Server]: No subscribers connected [ 1703758210 ]
[ Pub__1 ] 12:10:13 <-- Hello from app at 1703758213
[ Pub__2 ] 12:10:14 <-- Hello from app at 1703758214
[ Pub__2 ] 12:10:14 --> [Server]: No subscribers connected [ 1703758214 ]
[ Sub__2 ] 12:10:14 Sub__2 Is launched
[ Sub__1 ] 12:10:14 Sub__1 Is launched
[ Server ] 12:10:14 Subscriber connected:  [::1]:44007 . Total now: 1
[ Server ] 12:10:14 Subscriber connected:  [::1]:44320 . Total now: 1
[ Pub__1 ] 12:10:16 <-- Hello from app at 1703758216
[ Pub__1 ] 12:10:16 --> [Server]: New subscriber connected! [ 1703758216 ]
[ Sub__2 ] 12:10:16 --> [Pub__1]: Hello from app [ 1703758216 ]
[ Sub__1 ] 12:10:16 --> [Pub__1]: Hello from app [ 1703758216 ]
[ Pub__1 ] 12:10:19 <-- Hello from app at 1703758219
[ Sub__1 ] 12:10:19 --> [Pub__1]: Hello from app [ 1703758219 ]
[ Sub__2 ] 12:10:19 --> [Pub__1]: Hello from app [ 1703758219 ] 
```

## Testing

To run tests ```go test ./...```

## Configuration
App supports configuration. You may change ports for ports.

```
❯ go run main.go help
Usage of /tmp/go-build2853282965/b001/exe/main:
  -name string
        Name (default "Server")
  -port int
        PortSub number for clientSubscriber (default 6001)
  -portPub int
        PortSub number for publisher (default 6002)
exit status 3
```
