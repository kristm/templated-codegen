package test

//go:generate protoc --go_out=./build/         test.proto
//go:generate protoc --gotemplate_out=template_dir=templates:./build/ test.proto
