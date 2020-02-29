# QorPress Auth Example

## Description
This example a forked version of qor framework and implements the required patches to make work the qor auth example properly.

### Fixed
* Fixed：Unknown column 'basics.provider' in 'where clause')
  - https://github.com/qorpress/auth/pull/20/files
* Fixed: Missing From attribute in mailer
  - 

### Warning
All packages included in this program are forks from the original qor framework.
https://github.com/qorpress?utf8=%E2%9C%93&q=&type=fork&language=

### Pre-requesistes
- git
- docker
- docker compose
- golang 1.13/1.14

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
docker-compose up --build
open http://localhost:9000
```

#### Locally (for dev)

This version allows to run the example with an sqlite database.

```bash
cd $GOPATH/src/github.com/qorpress/qorpress-auth-example
go run main.go
open http://localhost:9000
```

### To do
* compile templates into an asset.go file
* multi-stage docker builder
* use the i18 package for loading login/register form translation

### Bugs
* cannot not use go.mod as it doesn't not copy template files (that's why we are using glide)
