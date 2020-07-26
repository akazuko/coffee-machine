# Coffee-machine

This is a go package which allows you to create a machine that can
serve beverages to N clients at a time based on availability of the
ingredients

### To run the tests
```
go test ./... -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html
```