FROM golang:1.10-stretch

ENV GOBIN /go/bin
ENV WKHTMLTOPDF_PATH /opt/wkhtmltox/bin

RUN mkdir -p /go/src/golang-questionnaire
RUN mkdir -p /opt/wkhtmltox

RUN apt-get update
RUN apt-get install -y xz-utils libxrender1 libfontconfig1 libxext6
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /opt

RUN wget -O wkhtmltox.tar.xz https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
RUN tar -xf wkhtmltox.tar.xz

WORKDIR /go/src/golang-questionnaire

COPY . .

RUN dep ensure

RUN go build -o questionnaireApp .
CMD ["./questionnaireApp"]