FROM golang:1.12-alpine

ENV HELM_VERSION "v2.13.1"
ENV HELM_HOME "/.helm"
ENV K8S_VERSION "v1.14.0"

RUN apk add --update ca-certificates \
    git \
    openssl \
    openssh \
    py-pip \
    findutils \
    bash \
    curl \
    go \
    gcc \
    musl-dev \
    make \
    docker

RUN curl -q https://raw.githubusercontent.com/helm/helm/master/scripts/get | bash -s -- --version $HELM_VERSION

RUN mkdir -p $HELM_HOME/plugins

RUN helm plugin install https://github.com/chartmuseum/helm-push

RUN helm init -c

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
RUN chmod u+x kubectl && mv kubectl /bin/kubectl

RUN adduser -D -u 1000 jenkins
USER 1000:1000
