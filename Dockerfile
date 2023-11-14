FROM golang:latest

ARG CURR_USER_UID=1000
ARG CURR_USER_GID=1000

RUN apt-get update -y --no-install-recommends && apt-get -y --no-install-recommends install git make cmake default-jre

RUN mkdir -p /home/gousr/project
RUN groupadd -f -g $CURR_USER_GID grp && useradd -M -N -u $CURR_USER_UID -G $CURR_USER_GID -d /home/gousr gousr
RUN chown -R gousr:grp /home/gousr

WORKDIR /home/gousr/project
COPY go.mod .
RUN chown gousr:grp go.mod
RUN go mod download
USER gousr
VOLUME /home/gousr/project

