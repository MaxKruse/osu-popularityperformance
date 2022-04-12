# Backend (using PostgreSQL, Redis and OAuth)

This is the backend, being responsible for OAuth2 authentication to osu! and internal user sessions.


## Requirements

You will need the [Semver-CLI](https://github.com/maykonlf/semver-cli) for tagging of images inside the build script.

## Build

For build details, check out the [Dockerfile](Dockerfile)

To build the backend, run:

> ./release

which will add a patch number to the version number and build the backend, or

> ./debug

which just builds and tags the image with `latest`.