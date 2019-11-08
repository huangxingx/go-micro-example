# AdminUser Service

This is the AdminUser service

Generated with

```
micro new infinite-window-micro/srv/adminUser --namespace=com.example --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.example.srv.adminUser
- Type: srv
- Alias: adminUser

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
./adminUser-srv
```

Build a docker image
```
make docker
```