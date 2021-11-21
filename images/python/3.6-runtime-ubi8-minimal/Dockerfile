FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4-205.1626828526 as base-builder

COPY pip.conf /etc/

RUN microdnf install -y python36-3.6.8 && microdnf clean all

COPY CHANGELOG /etc/CHANGELOG

FROM base-builder as final