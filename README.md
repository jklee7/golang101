# Golang 101
[![N|Solid](https://blog.golang.org/gopher/gopher.png)](https://golang.org/)

A collection of simple Go programs to demonstrate the fundamentals of Go.
For detailed instructions on installing Go, refer to the [Go](https://golang.org/doc/install) website.

## Installing Go
```sh
$ curl -O https://storage.googleapis.com/golang/go1.7.5.linux-amd64.tar.gz ~/
$ tar -C /usr/local -xzf ~/golang/go1.7.5.linux-amd64.tar.gz
```

## Environment Setup
```sh
$ echo $'\n'"export PATH=\$PATH:/usr/local/go/bin" | sudo tee --append /etc/profile
$ echo "export GOPATH=\$HOME/git/golang101/workspace/src" | sudo tee --append ~/.bashrc
$ echo "export GOBIN=\$HOME/git/golang101/workspace/bin" | sudo tee --append ~/.bashrc
```

## Workspace setup
```sh
$ mkdir  -p ~/git/golang101/workspace/src
$ mkdir  -p ~/git/golang101/workspace/bin
$ mkdir  -p ~/git/golang101/workspace/pkg
```

Workspace folder structure should look like this:
```
.
├── README.md
└── workspace
    ├── bin
    ├── pkg
    └── src
```

