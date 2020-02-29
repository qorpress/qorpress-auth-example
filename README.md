# QorPress Auth Example

This example a forked version of qor framework and implements the required patches to make work the qor auth example properly.

## Screenshots
![login](https://github.com/qorpress/qorpress-auth-example/raw/master/docs/screenshots/signin.png "login")
![register](https://github.com/qorpress/qorpress-auth-example/raw/master/docs/screenshots/register.png "register")
![recover password](https://github.com/qorpress/qorpress-auth-example/raw/master/docs/screenshots/recover.png "recover password")

## Pre-requesistes
- git
- docker
- docker compose
- golang 1.13/1.14

#### Fixed
- ~~Fixed：Unknown column 'basics.provider' in 'where clause')~~
  - https://github.com/qorpress/auth/pull/20/files
- ~~Fixed: Missing From attribute in mailer~~
  - https://github.com/qorpress/auth/blob/master/providers/password/confirm.go#L43

#### Warning
All packages included in this program are forks from the original qor framework.

### Install
```bash
cd $GOPATH/src/
mkdir -p qorpress
cd qorpress
git clone --depth=1 https://github.com/qorpress/qorpress-auth-example
cd qorpress-auth-example
mv .config/gopress.yml-example .config/gopress.yml
```

then change the values in the ```.config/gopress.yml``` file.

### Run
There are 2 ways to run this example, one is provided from a docker container (alpine based), the other is to run it locally with golang installed on your workstation.

#### Docker (quick)

This version allows to run the example with mysql container.

```bash
cd $GOPATH/src/github.com/qorpress/qorpress-auth-example
docker-compose up -d --build
```
open http://localhost:4000 in your browser

or
```bash
docker build -t qorpress/qorpress-auth-example --no-cache .
docker run -ti -p 4000:4000 -v ${PWD}/.config/gopress.yml:/go/bin/.config/gopress.yml qorpress/qorpress-auth-example
```
open http://localhost:4000 in your browser

#### Locally (for dev)

This version allows to run the example with an sqlite database.

```bash
cd $GOPATH/src/github.com/qorpress/qorpress-auth-example
go run main.go
```
open http://localhost:4000 in your browser

### Links
- Login - http://localhost:4000/auth/login
- Register - http://localhost:4000/auth/register
- New Password - http://localhost:4000/auth/password/new

### To do
* compile templates into an asset.go file
* multi-stage docker builder
* use the i18 package for loading login/register form translation

### Bugs
* cannot not use go.mod as it doesn't not copy template files (that's why we are using glide)
