# 1. Сборка -------------------------
FROM ruby:2.5-alpine3.7 AS builder
RUN apk update && apk --update add \
    libc6-compat \
    libstdc++ \
    libssl1.0 \
    openssh-client \
    postgresql-client \
    tzdata && \
  apk --update add --virtual build-dependencies \
    automake \
    autoconf \
    build-base \
    git \
    gmp-dev \
    ruby-dev \
    libc-dev \
    libtool \
    libxml2-dev \
    libxslt-dev \
    linux-headers \
    postgresql-dev

WORKDIR /opt/src

# 2. Разработка -------------------------
FROM builder AS development
ENV BUNDLE_PATH=/bundle \
    BUNDLE_BIN=/bundle/bin \
    GEM_HOME=/bundle
ENV PATH="${BUNDLE_BIN}:${PATH}"

COPY Gemfile*  /opt/src/
RUN bundle install

COPY . /opt/src

# 3. Релиз -------------------------
# FROM builder AS release
# RUN apk del build-dependencies

# ENV RAILS_ENV=production

# EXPOSE 8080
# CMD [ "bundle", "exec", "puma", "-C", "config/puma.rb" ]
