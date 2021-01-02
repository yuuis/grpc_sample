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
install protobuf and plugins
```shell script
brew install protobuf
go get -u -v github.com/golang/protobuf/protoc-gen-go
go get -u -v google.golang.org/grpc
```

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
bundle exec grpc_tools_ruby_protoc -I ../proto/ --ruby_out=app/gen/api/image/uploader --grpc_out=app/gen/api/image/uploader ../proto/image_uploader.proto
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

```ruby
# upload an image
ImageUploader.chunked_upload('file_path')
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

## Reference

[スターティングgRPC](https://www.amazon.co.jp/%E3%82%B9%E3%82%BF%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0gRPC-%E6%8A%80%E8%A1%93%E3%81%AE%E6%B3%89%E3%82%B7%E3%83%AA%E3%83%BC%E3%82%BA%EF%BC%88NextPublishing%EF%BC%89-%E6%AD%A6%E4%B8%8A-%E5%B0%86%E6%A8%B9-ebook/dp/B087R87L6Z)
