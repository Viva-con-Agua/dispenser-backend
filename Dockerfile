FROM golang

WORKDIR /go/src/dispenser-backend

ADD . /go/src/dispenser-backend
RUN ./install_packages.sh


CMD ["go", "run", "server.go"]
