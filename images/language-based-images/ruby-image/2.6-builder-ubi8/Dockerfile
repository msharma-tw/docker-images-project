FROM registry.access.redhat.com/ubi8:8.4-206.1626828523 as base-builder

# Install bare minimum builder utilities
RUN dnf install -y \
    bzip2 \
    zip \
    unzip \
    wget \
    gcc-c++ \
    make \
    automake \
    autoconf \
    libtool \
    git \
    && dnf clean all 

FROM base-builder as ruby

# This line is important to make sure rubygems puts bin scripts in $GEM_HOME/bin and not /usr/bin
# this helps when using docker multi-stage build: fetching and building gems in a builder image and copying
# $GEM_HOME to the ruby runtime image.
ENV GEM_HOME=/usr/local/share/gems
ENV PATH=$GEM_HOME/bin:$PATH

RUN dnf -y module reset ruby  \
    && dnf -y module enable ruby:2.6  \
    && dnf install -y  \
      ruby-2.6.7 \
      ruby-devel-2.6.7 \
      libffi-devel \
      rubygem-rake \
      rubygem-bundler \
      redhat-rpm-config \
    && gem sources --remove https://rubygems.org/ \
    && dnf clean all

COPY CHANGELOG /etc/CHANGELOG

FROM ruby as final
