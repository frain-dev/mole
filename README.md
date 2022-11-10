# Mole

Mole is an HTTP connect tunnel powered by smokescreen

## Install

Build the docker image and tag it appropriately

```bash
$ docker build -t frain-dev/mole:latest .
```

## Run

```bash
$ docker run -p 4750:4750 -e PROXY_PASSWORD=$PROXY_PASSWORD frain-dev/mole:latest
```

## Configuration

Configure the server that wants to use the proxy by setting the HTTP_PROXY environment variable

```bash
export HTTP_PROXY="http://IP_ADDRESS:PORT"
```

## Usage

For all outgoing requests the `Proxy-Authorization` header must be set to the value in the PROXY_PASSWORD environment variable.
