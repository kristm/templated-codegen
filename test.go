package test

//go:generate protoc --go_out=./build/         client.proto
//go:generate protoc --gotemplate_out=debug=true,template_dir=templates:./build/ client_nooptional.proto
