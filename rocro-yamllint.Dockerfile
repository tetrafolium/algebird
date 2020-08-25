FROM python:3-alpine3.11 AS yamllint-task

RUN echo "===> Install golang ..." && \
    apk add --update --no-cache go && \
    echo -n "+++ " ; go version

ENV GOBIN=$GOROOT/bin \
    GOPATH=/.go \
    PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH

RUN echo "===> Install the yamllint ..." && \
    pip3 install 'yamllint>=1.24.0,<1.25.0' && \
    echo -n "+++ " ; yamllint --version

ENV REPO=${GOPATH}/src/github.com/tetrafolium/algebird \
    OUTDIR=/.reports
RUN mkdir -p ${REPO} ${OUTDIR}
COPY . ${REPO}
WORKDIR ${REPO}

RUN echo "===> Run yamllint ..." && \
    yamllint -f parsable . > ${OUTDIR}/yamllint.out || true

RUN echo "===> Convert reports to SARIF ..." && \
    go run .rocro/yamllint/converter/cmd/main.go \
        < ${OUTDIR}/yamllint.out \
        > ${OUTDIR}/yamllint.sarif

RUN ls -la ${OUTDIR}
RUN echo '----------' && \
    cat -n ${OUTDIR}/yamllint.out && \
    echo '----------' && \
    cat -n ${OUTDIR}/yamllint.sarif && \
    echo '----------'
