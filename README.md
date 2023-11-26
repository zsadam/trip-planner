Build the Docker image
```shell
docker build -t trip-planner .
```

Run the tests
```shell
docker run -v .:/app trip-planner go test ./...
```

Run the linter
```shell
docker run -v .:/app trip-planner golangci-lint run
```

