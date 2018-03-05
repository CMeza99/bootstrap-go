# bootstrap-go

A and easy way to create images for Debian-base distributions.

This is an early *work in progress* and is not functional yet.

## Background

This was inspired from [bootstrap-vz|https://github.com/andsens/bootstrap-vz]. It is specificly built for Debian. After some work in an effort to abstract the upstream code to make it usable for Debian-based distro, I realized a large refactor was needed. So I decided to start fresh and take the lessons learned from bootstrap-vz. I also thought this to be a good opportunity get to some experience with Golang.

## Development Environment

```sh
docker-compose --project-name bootstrap-go build --force-rm --pull
docker-compose --project-name bootstrap-go run --rm bootstrap-go bash
```
