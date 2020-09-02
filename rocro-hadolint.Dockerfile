FROM hadolint/hadolint:latest AS hadolint-task

RUN echo "===> Install git ..." && \
    apk add --update --no-cache git && \
    echo "+++ $(git version)"

RUN echo "===> Install golang ..." && \
    apk add --update --no-cache go && \
    echo "+++ $(go version)"

ENV GOBIN="$GOROOT/bin" \
    GOPATH="/.go" \
    PATH="${GOPATH}/bin:/usr/local/go/bin:$PATH"

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

RUN echo "===> Run hadolint ..." && \
    ( find . -name '*Dockerfile*' | \
      xargs hadolint --format json > "${OUTDIR}/hadolint.json" ) || true

RUN ls -la "${OUTDIR}"
RUN echo '----------' && \
    cat "${OUTDIR}/hadolint.json" && \
    echo '----------'

RUN echo "===> Convert hadolint JSON to SARIF ..." && \
    go run "${TOOLDIR}/hadolint/cmd/main.go" \
        < "${OUTDIR}/hadolint.json" \
        > "${OUTDIR}/hadolint.sarif"

RUN ls -la "${OUTDIR}"
RUN echo '----------' && \
    cat -n "${OUTDIR}/hadolint.sarif" && \
    echo '----------'
