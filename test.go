package test

//go:generate protoc --go_out=./build/         client.proto
//go:generate protoc --gotemplate_out=template_dir=templates:./build/ client.proto
