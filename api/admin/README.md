# Admin Service

This is the Admin service

Generated with

```
micro new infinite-window-micro/api/admin --namespace=com.infinite --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.infinite.web.admin
- Type: web
- Alias: admin

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./admin-web
```

Build a docker image
```
make docker
```