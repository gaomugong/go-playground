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

