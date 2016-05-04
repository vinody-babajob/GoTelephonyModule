FROM ubuntu
RUN apt-get -y update
RUN apt-get install -y software-properties-common
RUN apt-get install -y rabbitmq-server

RUN rabbitmqctl add_user admin admin
RUN rabbitmqctl set_permissions admin ".*" ".*" ".*"

RUN rabbitmq-plugins enable rabbitmq_management

RUN service rabbitmq-server start

EXPOSE 15672
EXPOSE 5672
