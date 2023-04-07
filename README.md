# Openfeature example

## Run flagd provider

```sh
make flagd
```

## Run code

```sh
go run main.go
```

- open <http://localhost:8080/hello>

## Change flags via flags.flagd.json file

```json
{
  "flags": {
    "welcome-message": {
      "variants": {
        "on": true,
        "off": false
      },
      "state": "ENABLED",
      "defaultVariant": "on/off" // as you need it
    }
  }
}
```

## Current problems

- <https://github.com/open-feature/flagd>
- after change json I don't see change without restart flagd container
