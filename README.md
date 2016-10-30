# Nhite application

Here are the sources of the Nhite application.

The runtime is pure javascript, but the developement is made with GO and translated to javascript with `gopherjs`

## Cordova / Javascript

Nhite is a hybrid application iOS / Android / Browser.
The core application is based on `cordova`.

Therefore you need a working cordova installation to make it run.

## golang / gopherjs

### Requirements

You need [gopherjs](https://github.com/gopherjs/gopherjs) and [humble](https://github.com/go-humble/humble) to work on the sources.

`go get -v -u github.com/nhite/frontend`

should do most of the work.

### Temple

The templates are generated with the help of the [temple](https://github.com/go-humble/temple) tool
 
`go get -u github.com/go-humble/temple`

# Make it work

to generate the javascript file from the go sources:

`go generate ./...`

to test in: 
 
`cd nhiteApp && cordova run browser`

# Backend

The authentication flow requires a backend (see [nhite engine](https://github.com/nhite/engine) to exchange the oauth2 code for a JWT

TODO: document this part

