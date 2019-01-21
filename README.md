# smpl

smpl ("Simple") is a simple reverse proxy. The initial commit is meant purely for serving my dwn repository, but my next goal will be to make it multipurpose and to provide configurability via a graphical (probably web) UI.

## Inspiration

My goal is to eventually achieve what I thought caddy was supposed to be--a simple web server with a super simple ui for configuration. Caddy achieved some of the simplicity I expected through its fantastic defaults and flexible plugins. However, I felt that they missed the vision when they kept using the config format reminiscent of nginx. While it is fantastic for most developers, the less tech savvy designers (who might be using host-provided control panels currently) aren't going to necessarily use caddy to take the next step into controlling their technical destiny. I'd like to provide a solution somewhat in between--one with a shiny UI and lots of explanation and documentation available alongside simplified configurations. I won't necessarily ever support load-balancing or the most efficient static file serving, but it should be more than fast enough to support dozens to hundreds of small business websites on the same budget server.

## Usage

```
go build

smpl serve
smpl echo
smpl help
```

## Roadmap

Right now I'm hardcoding support for serving DWN over TLS. I will immediately work on configuring a more flexible reverse proxy with sane production defaults and serving TLS from certs stored via certmagic. I will also work on loading from and saving to a bbolt database. After that I will begin working on an admin UI via html templates.

## License

MIT license.