# gomock test

Run
```shell
go mod tidy
rm -f zz_generated_mock.go
go generate ./...
```

mockgen will print the generated code to stdout and create a empty `zz_generated_mock.go` file.

Or set a breakpoint in github.com/golang/mock/mockgen/model.go#L396 in function `func parameterFromType(t reflect.Type) (*Parameter, error)`
and debug the prog generated from mockgen (commited in this repo for testing `main/gen.go`) 

tt.Type will be `TestGenericType[github.com/daolis/mockgentest/test.TestType]`
