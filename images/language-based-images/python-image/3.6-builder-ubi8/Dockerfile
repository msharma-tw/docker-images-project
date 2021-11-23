FROM registry.access.redhat.com/ubi8:8.4-206.1626828523 as base-builder

COPY pip.conf /etc/

# Pex is really convenient for distributing Python apps
# Python linters and static code analysis useful for builds
RUN dnf install -y \
    python36-3.6.8 \
    python36-devel-3.6.8 \
    && pip3 install --no-cache-dir \
    pex \
    pyflakes \
    pep8 \
    && dnf clean all

COPY CHANGELOG /etc/CHANGELOG

FROM base-builder as final

