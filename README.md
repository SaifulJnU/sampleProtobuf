# sampleProtobuf

steps:
first make the proto/service.proto file
second to compile it at first create a foldre for my casd i named it 'pb' then compile it by the following command from the root project directory:
``` protoc -I proto/ proto/service.proto --go_out=paths=source_relative:pb --go-grpc_out=paths=source_relative:pb```
this command produce auto generated code the new file you will find under the 'pb' folder.
third then access data from main.go

by the way care full abut package name and package importation 
