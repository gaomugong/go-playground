# consul

# install
- https://developer.hashicorp.com/consul/install
- brew tap hashicorp/tap
- mac/intel: brew install hashicorp/tap/consul
- mac/m1: arch -arm64 brew install hashicorp/tap/consul

# start & test
- https://developer.hashicorp.com/consul/commands
- To start hashicorp/tap/consul now and restart at login:
  brew services start hashicorp/tap/consul
  
- Or, if you don't want/need a background service you can just run:
  `/opt/homebrew/opt/consul/bin/consul agent -dev -bind 127.0.0.1`

- https://developer.hashicorp.com/consul/docs/agent
- consul agent -dev -node dev-consul
- consul agent -dev -data-dir=/tmp/consul
- http://localhost:8500

# demo 
1. start 3 'demo-server' instance with different id

- ./main -port 8000
- ./main -port 8001
- ./main -port 8002

2. run client several times

- ./main
- ./main

2023/12/20 00:12:46 Health endpoints:
127.0.0.1:8000
127.0.0.1:8001
127.0.0.1:8002
2023/12/20 00:12:46 target-grpc -> 127.0.0.1:8000
2023/12/20 00:12:46 response message: hello, pitou

------

2023/12/20 00:13:01 Health endpoints:
127.0.0.1:8000
127.0.0.1:8001
127.0.0.1:8002
2023/12/20 00:13:01 target-grpc -> 127.0.0.1:8001
2023/12/20 00:13:01 response message: hello, pitou

