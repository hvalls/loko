# What is loko

loko is a very simple CLI tool that helps you creating docker-compose.yml file and including services you usually need to it.

# Install 

1. Clone this repo:

```
$ git clone https://github.com/hvalls/loko
```

2. Install:

```
$ make install
```

3. Build executable:

```
$ make build
```

# How to use

## Initialize

This creates an empty docker-compose.yml file in the current directory:
```
$ loko init
```

## Add service

This will add a new service in the docker-compose.yml using default configuration:

```bash
$ loko add <service>
# e.g: loko add redis
```

## Services

loko supports services included in `data/services` directory. They are copied to your `~/.loko/data/services` directory when it is installed. You can add your own services to that directory. 


