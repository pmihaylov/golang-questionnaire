FROM golang:1.10-stretch AS build

ENV GOBIN /go/bin
ENV WKHTMLTOPDF_PATH /opt/wkhtmltox/bin

ARG wkhtmltox_version=0.12.4
ARG wkhtmltox_dir=/opt/wkhtmltox
ARG project_dir=/go/src/golang-questionnaire

### Create directories
RUN mkdir -p ${wkhtmltox_dir} ${project_dir}

### Stuff for wkhtml and unpacking .tar.xz
WORKDIR ${wkhtmltox_dir}

RUN apt-get update && \
    apt-get install -y \
    apt-transport-https \
    curl \
    libfontconfig1 \
    libxext6 \
    libxrender1 \
    xz-utils

### Get wkhtltopdf, golang-migrate/migrate for postgres and golang/dep
RUN curl -LsS https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/${wkhtmltox_version}/wkhtmltox-${wkhtmltox_version}_linux-generic-amd64.tar.xz \
    | tar xJ --strip-components=1 && \
    go get -u -d github.com/golang-migrate/migrate/cli github.com/lib/pq && \
    go build -tags 'postgres' -o /usr/local/bin/migrate github.com/golang-migrate/migrate/cli && \
    curl -sS https://raw.githubusercontent.com/golang/dep/master/install.sh | sh


### Install dependencies
COPY Gopkg.lock Gopkg.toml ${project_dir}/
WORKDIR ${project_dir}
RUN dep ensure -vendor-only

### Copy and build app itself
COPY . ${project_dir}
RUN go build -o questionnaireApp .
CMD ["./questionnaireApp"]