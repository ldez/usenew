# usenew

Find calls to functions that can be replaced with the built-in `new` function.

## Usage

### Inside golangci-lint

```yaml
linters:
  enable:
    - usenew
```

### As CLI

```bash
go install github.com/ldez/usetesting/cmd/usenew@latest
```

## References

- https://go.dev/doc/go1.26#language
