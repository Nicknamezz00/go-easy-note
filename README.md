# go-easy-note

## Features
- Middleware、Rate Limiting、Request Retry、Timeout Control、Connection Multiplexing
- Tracing
    - use jaeger to tracing
- Customized BoundHandler
    - achieve CPU utilization rate customized bound handler
- Service Discovery and Register
    - use [registry-etcd](https://github.com/kitex-contrib/registry-etcd) to discovery and register service

```
                           ┌────────────────────────┐
 ┌─────────────────────────┤                        ├───────────────────────────────┐
 │                         │          api           │                               │
 │      ┌──────────────────►                        │◄──────────────────────┐       │
 │      │                  └───────────▲────────────┘                       │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                           resolve                                 │       │
 │      │                              │                                    │       │
req    resp                            │                                   resp    req
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                   ┌──────────▼─────────┐                          │       │
 │      │                   │                    │                          │       │
 │      │       ┌───────────►       etcd         ◄─────────────────┐        │       │
 │      │       │           │                    │                 │        │       │
 │      │       │           └────────────────────┘                 │        │       │
 │      │       │                                                  │        │       │
 │      │     register                                           register   │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
┌▼──────┴───────┴───┐                                           ┌──┴────────┴───────▼─┐
│                   │───────────────── req ────────────────────►│                     │
│       note        │                                           │        user         │
│                   │◄──────────────── resp ────────────────────│                     │
└───────────────────┘                                           └─────────────────────┘
      thrift                                                           protobuf
```

## Structure
| directory        | Description          |
|------------------|----------------------|
| /pkg/config/sql/ | SQL Script           |
| /pkg/constants   | Constants            |
| /pkg/errno/      | Custom error code    |
| /pkg/middleware/ | RPC middleware       | 
| /pkg/tracer/     | Jaeger               |
| /kitex_gen       | Kitex generated code |
| /idl/            | Protobuf & Thrift    |
| /cmd/api/        | API service          |
| /cmd/note/       | Note API service     | 
| /cmd/user/       | User API service     |

## Quick Start

### 1. Build
```shell
docker-compose up
```
### 2. RPC
```shell
cd cmd/note
sh build.sh
sh output/bootstrap.sh

cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 3. API Server
```shell
cd cmd/api
chmod +x run.sh
./run.sh
```

### 4. Jaeger

see `http://localhost:16686/`