FROM ubuntu:24.04

WORKDIR /tmp

RUN rm -rf /var/lib/apt/lists/*
RUN apt update 
RUN apt-get update

RUN apt-get install -y wget
RUN apt-get install -y curl

# Go
WORKDIR /tmp
RUN rm -rf /usr/local/go
RUN wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz
RUN tar  -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin

# TypeScript
# esbuild を使います
WORKDIR /home/app
RUN apt install -y nodejs
RUN apt install -y npm
RUN npm i -D esbuild esbuild-register

# PHP
WORKDIR /tmp
RUN apt install -y php

# rust
WORKDIR /tmp
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs -o sh.rustup.rs
RUN sh sh.rustup.rs -y 

WORKDIR /home/app
