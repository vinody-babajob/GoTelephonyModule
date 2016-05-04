FROM ubuntu
ARG ENV
RUN apt-get -y update
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:ubuntu-lxc/lxd-stable
RUN apt-get install -y golang
RUN apt-get install -y git

RUN cd /home && mkdir golangsrc
ENV  GOPATH "/home/golangsrc"

EXPOSE 8080

ENV ENVIRONMENT ${ENV}
