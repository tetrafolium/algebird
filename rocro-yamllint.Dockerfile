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

ENV REPO=github.com/tetrafolium/algebird \
    TASKSTOOL=github.com/tetrafolium/inspecode-tasks \
    REPODIR=${GOPATH}/src/${REPO} \
    TOOLDIR=${GOPATH}/src/${TASKSTOOL} \
    OUTDIR=/.reports
RUN mkdir -p ${REPODIR} ${OUTDIR}
COPY . ${REPODIR}
WORKDIR ${REPODIR}

RUN echo "===> Get tool ..." && \
    go get ${TASKSTOOL} || true

RUN echo "===> Run yamllint ..." && \
    yamllint -f parsable . > ${OUTDIR}/yamllint.issues || true

RUN echo "===> Convert yamllint issues to SARIF ..." && \
    go run $TOOLDIR}/yamllint/cmd/main.go \
        < ${OUTDIR}/yamllint.issues \
        > ${OUTDIR}/yamllint.sarif

RUN ls -la ${OUTDIR}
RUN echo '----------' && \
    cat -n ${OUTDIR}/yamllint.issues && \
    echo '----------' && \
    cat -n ${OUTDIR}/yamllint.sarif && \
    echo '----------'
