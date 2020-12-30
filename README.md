# gRPC sample

## Directory

```
.
├── api ... API Server (Golang)
├── client ... API Client (Rails)
├── proto ... proto files
```

## API Server

### How to update boiler plate codes for server

```shell script
cd grpc_sample
protoc -Iproto --go_out=plugins=grpc:api proto/*.proto
```

### Run server
```shell script
cd api
go run server/server.go
```

## Client (Rails)

### How to update boiler plate codes for client

```shell script
cd client
bundle install

bundle exec grpc_tools_ruby_protoc -I ../proto/ --ruby_out=app/gen/api/pancake/maker --grpc_out=app/gen/api/pancake/maker ../proto/pancake.proto
```

### Run client
```shell script
bundle exec rails c
```

```ruby
# request to bake a pancake  
Bakery.bake_pancake(Bakery::Menu::CLASSIC)

# request to send a report
Bakery.report
```

## Client (`grpc_cli`)

### Install `grpc_cli`

brew or build from repository > [grpc/command_line_tool.md at master · grpc/grpc · GitHub](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md)
```shell script
brew tap grpc/grpc
brew install grpc
```

### Use

```shell script
# ls
grpc_cli ls localhost:50051 pancake.maker.PancakeBakerService -l

# call
./grpc_cli call localhost:50051 pancake.maker.PancakeBakerService.Bake 'menu: 1'
./grpc_cli call localhost:50051 pancake.maker.PancakeBakerService.Report ''
```
