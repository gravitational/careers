FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y --no-install-recommends cmake && \
    rm -rf /var/lib/apt/lists/*

RUN apt-get update -y --fix-missing && \
    apt-get -q -y upgrade && \
    apt-get install -q -y --no-install-recommends \
        texlive \
        texlive-base

WORKDIR /tmp
