FROM python:3-alpine3.11

RUN echo "===> Install golang ..." && \
    apk add --update --no-cache go && \
    go version

ENV GOBIN=$GOROOT/bin \
    GOPATH=/go \
    PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH

RUN echo "===> Install the yamllint ..." && \
    pip3 install 'yamllint>=1.24.0,<1.25.0' && \
    yamllint --version

ENV REPO=${GOPATH}/src/github.com/tetrafolium/algebird
RUN mkdir -p ${REPO}
COPY . ${REPO}
WORKDIR ${REPO}

RUN yamllint -f parsable . > .rocro/yamllint.out || true
RUN go run .rocro/yamllint/converter/cmd/main.go < .rocro/yamllint.out > .rocro/yamllint.sarif
RUN ls -la .rocro
