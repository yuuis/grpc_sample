# gRPC Client

## call

```sh
bundle exec rails c
```

```rb
# request to bake a pancake  
Bakery.bake_pancake(Bakery::Menu::CLASSIC)

# request to send a report
Bakery.report
```

## update boiler plate codes

```sh
bundle exec grpc_tools_ruby_protoc -I ../proto/ --ruby_out=app/gen/api/pancake/maker --grpc_out=app/gen/api/pancake/maker ../proto/pancake.proto
```
