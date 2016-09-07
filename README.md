# resolvplox

resolvplox is a small DNS proxy (still in proof of concept) that allows one to define the routes by which certain domains are resolved. For example, if one's normal DNS server is not resolving `twitter.com`, they can configure resolvplox to resolve twitter.com, and only twitter.com, with an alternative server.

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
