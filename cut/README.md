# Re-implementing GNU `cut` with Go

## Testing

Run the tests with
```
go test -v -coverprofile=coverage.out
```

Check coverage
```
go tool cover -html coverage.out
```