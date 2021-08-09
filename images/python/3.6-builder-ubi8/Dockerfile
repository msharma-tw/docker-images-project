FROM registry.access.redhat.com/ubi8:8.4-206.1626828523 as python3

COPY pip.conf /etc/

# Pex is really convenient for distributing Python apps
# Python linters and static code analysis useful for builds
RUN dnf install -y \
        python39-3.9.2 \
        python39-devel-3.9.2 \
    && pip3 install --no-cache-dir \
        pex==2.1.45 \
        pylint==2.9.6 \
        pyflakes==2.3.1 \
        pep8==1.7.1 \
    && dnf clean all

COPY CHANGELOG /etc/CHANGELOG

FROM python3 as final

