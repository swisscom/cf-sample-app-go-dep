# cf-sample-app-go-dep
A sample [Go](https://golang.org/) application to deploy to Cloud Foundry using
[dep](https://github.com/golang/dep)

This basic example uses a custom router and serves 2 endpoints
`/health` and `/` it uses the environment var `VERSION` defined in the
[manifest.yml](manifest.yml) the one later is returned when querying the
`/health` endpoint.

The application has the following structure:

    |-- app
    |     `--hello
    |       `--hello.go
    |-- health.go
    |-- main.go
    `-- manifest.yml

# How to use it

Install and set your [GO environment](https://golang.org/doc/install)

For example using $HOME/go for your workspace

    $ export GOPATH=$HOME/go

Create the directory:

    $ mkdir -p $HOME/go/src/github.com/swisscom

Clone project into that directory:

    $ git clone https://github.com/swisscom/cf-sample-app-go-dep.git $HOME/go/src/github.com/swisscom/cf-sample-app-go-dep

> [git](https://git-scm.com/) is required to clone the repository.

To test the app locally before deploying it in to app-cloud, build it by just typing make:

    $ cd $HOME/go/src/github.com/swisscom/cf-sample-app-go-dep
    $ make

After make finishes you can run the app:

    $ ./cf-sample-app-go-dep
    2018/07/06 13:39:12 Adding path: /health [ALL]
    2018/07/06 13:39:12 Adding path: / [ALL]

> This steps requires [dep](https://github.com/golang/dep)

Visit [http://localhost:8080](http://localhost:8080)

## Deploying the app to app-cloud

If you tested locally, remove the `vendor` directory or simply use:

    $ make clean

Then just use:

    $ cf push
