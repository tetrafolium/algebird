FROM python:3-alpine3.11 AS yamllint-task

RUN echo "===> Install golang ..." && \
    apk add --update --no-cache go && \
    echo "+++ $(go version)"

ENV GOBIN="$GOROOT/bin" \
    GOPATH="/.go" \
    PATH="${GOPATH}/bin:/usr/local/go/bin:$PATH"

RUN echo "===> Install the yamllint ..." && \
    pip3 install 'yamllint>=1.24.0,<1.25.0' && \
    echo "+++ $(yamllint --version)"

ENV REPOPATH="github.com/tetrafolium/algebird" \
    TOOLPATH="github.com/tetrafolium/inspecode-tasks"
ENV REPODIR="${GOPATH}/src/${REPOPATH}" \
    TOOLDIR="${GOPATH}/src/${TOOLPATH}"

RUN echo "===> Get tool ..." && \
    go get -u "${TOOLPATH}" || true

ARG OUTDIR
ENV OUTDIR="${OUTDIR:-"/.reports"}"

RUN mkdir -p "${REPODIR}" "${OUTDIR}"
COPY . "${REPODIR}"
WORKDIR "${REPODIR}"

RUN echo "===> Run yamllint ..." && \
    yamllint -f parsable . > "${OUTDIR}/yamllint.issues" || true

RUN ls -la "${OUTDIR}"
RUN echo '----------' && \
    cat -n "${OUTDIR}/yamllint.issues" && \
    echo '----------'

RUN echo "===> Convert yamllint issues to SARIF ..." && \
    go run "${TOOLDIR}/yamllint/cmd/main.go" \
        < "${OUTDIR}/yamllint.issues" \
        > "${OUTDIR}/yamllint.sarif"

RUN ls -la "${OUTDIR}"
RUN echo '----------' && \
    cat -n "${OUTDIR}/yamllint.sarif" && \
    echo '----------'
