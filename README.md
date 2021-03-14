# go-wavy
[![Go Reference](https://pkg.go.dev/badge/github.com/ogkevin/go-wavy.svg)](https://pkg.go.dev/github.com/ogkevin/go-wavy)

Golang client for interacting with [wavy.fm](https://wavy.fm).

## Installation

```bash
go get github.com/OGKevin/go-wavy/wavy
```

## Usage

```go
import "github.com/OGKevin/go-wavy/wavy"

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

c := wavy.NewClient(ctx, hclog.NewNullLogger(), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET")
profile, err := c.UserService().GetProfile(ctx, wavy.UserURI{Username: "OGKevin"})
if err != nil {
    panic(err)
}
```

## License

[MIT](https://choosealicense.com/licenses/mit)
