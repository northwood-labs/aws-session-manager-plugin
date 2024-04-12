FROM golang:1.22

RUN apt-get -y update && \
    apt-get -y upgrade && \
    apt-get -y install rpm tar gzip wget zip && \
    apt-get clean all

RUN mkdir /session-manager-plugin
WORKDIR /session-manager-plugin
