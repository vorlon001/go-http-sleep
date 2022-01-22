# go-http-sleep

delayed response http server, useful for testing various timeout issue for application running behind proxy

### Usage

    $ go run ./main.go

    $ curl 127.0.0.1:8080/sleep?s=5s # sleep 5s and return
    delayed for 5s

alternatively, you may run it as minimal Docker image

    $ docker build -t go-http-sleep .

    $ docker image ls
    REPOSITORY      TAG       IMAGE ID       CREATED              SIZE
    go-http-sleep   latest    9afd5ccafcb9   About a minute ago   6.07MB

    $ docker run -d -p 8080:8080 go-http-sleep
    8dad85556621c1e2077c9ff921967bd75b11ee00123dee1a2e56fb9abaf6936b

    $ curl "127.0.0.1:8080/sleep?s=5"
    delayed for 5s


### LICENSE

[UNLICENSE](LICENSE)
