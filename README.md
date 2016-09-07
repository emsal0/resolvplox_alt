# resolvplox

resolvplox is a small DNS proxy (still in proof of concept) that allows one to define the routes by which certain domains are resolved. For example, if one's normal DNS server is not resolving `twitter.com`, they can configure resolvplox to resolve twitter.com, and only twitter.com, with an alternative server.

## Plox configs

Configuration is defined in what are called "plox configs", which are files that follow a particular basic format:
* Each line contains a whitespace-separated pair of a particular domain and the domain server by which it is to be resolved (the server must be an IPv4 address as of now).

    Example: `twitter.com 8.8.8.8`
* the default DNS server for all domains not specified in the current plox config must be defined, with the domain name `*`.

    Example: ```* 8.8.8.4```

An example complete valid plox config, called `simpleconfig.plox`:

```
twitter.com 8.8.8.8
* 8.8.8.4
```

## Installation and running

Simply use the `go install` tool.

```
go install github.com/emsal1863/resolvplox_alt
```

Then run with:

```
resolvplox_alt {PLOX CONFIG FILE}
```

This starts up a UDP server on localhost port 20841. Set your DNS settings in your OS to send DNS queries to localhost:20841. Your DNS settings should be working!


## Future

Some things need to be done:

* rename to resolvplox instead of resolvplox_alt (the original resolvplox was an Erlang project that isn't finished and probably won't be in the forseeable future)
* Support DHCP DNS servers

