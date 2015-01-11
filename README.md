Gonalytics (backend)
=============

Gonalytic is a open source web analytics application. It consists of:
- backend (this repository)
- tracking script: https://github.com/piotrkowalczuk/gonalytics-tracking-script
- dashboard: https://github.com/piotrkowalczuk/gonalytics-dashboard

Installation
------------
1. Set you GOPATH properly (http://golang.org/doc/code.html#GOPATH)
2. `go get github.com/piotrkowalczuk/gonalytics-backend`
3. `go get` if some dependencies are missing
4. Create `conf/**/{env}.xml` based on `conf/**/{env}.xml.dist`

Commands
--------

#### Build

    go build

#### Starting

    ./gonalytics-backend tracker - starts tracking API HTTP server (visits, redirects etc)
    ./gonalytics-backend actions-worker - starts worker process that process incoming actions
    ./gonalytics-backend api - starts CRUD API HTTP server
    ./gonalytics-backend help [command] - display help message about available commands

Web API
--------
Public API documentation is available at http://docs.gonalyticsbackend.apiary.io. The documentation is integrated into the repository. The current documentation is a reflection of the `apiary.apib` file.

Dependencies
------------
- Cassandra
- Kafka
- ZooKeeper (Kafka deps)
