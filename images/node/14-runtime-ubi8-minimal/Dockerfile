FROM registry.access.redhat.com/ubi8:8.4-206.1626828523 as node14

RUN dnf -y module reset nodejs  \
    && dnf -y module enable nodejs:14  \
    && dnf install -y nodejs-14.17.5 \
    && npm install -g yarn@1 \
    && dnf clean all

COPY CHANGELOG /etc/CHANGELOG

FROM node14 as final
