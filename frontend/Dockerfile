# 1. Разработка
FROM node:8.9-alpine AS development
RUN apk --update add openssl git

# git-lfs
RUN apk add --update --virtual build-dependencies curl
RUN curl -sLO https://github.com/github/git-lfs/releases/download/v2.0.1/git-lfs-linux-amd64-2.0.1.tar.gz
RUN tar xf git-lfs-linux-amd64-2.0.1.tar.gz && mv git-lfs-2.0.1/git-lfs /usr/bin/ && rm -Rf git-lfs-2.0.1
RUN git lfs install

WORKDIR /opt/src

COPY package.json yarn.lock /opt/src/

RUN yarn
COPY . /opt/src

