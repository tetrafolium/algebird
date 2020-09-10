FROM golang:1.14-alpine AS hadolint-task

### Install git ...
RUN apk add --update --no-cache git && \
    echo "+++ $(git version)"

### Install hadolint ...
RUN apk add --update --no-cache hadolint && \
    echo "+++ $(hadolint --version)"

ENV GOBIN="$GOROOT/bin" \
    GOPATH="/.go" \
    PATH="${GOPATH}/bin:/usr/local/go/bin:$PATH"

ENV REPOPATH="github.com/tetrafolium/algebird" \
    TOOLPATH="github.com/tetrafolium/inspecode-tasks"
ENV REPODIR="${GOPATH}/src/${REPOPATH}" \
    TOOLDIR="${GOPATH}/src/${TOOLPATH}"

### Get inspecode-tasks tool ...
RUN go get -u "${TOOLPATH}" || true

ARG OUTDIR
ENV OUTDIR="${OUTDIR:-"/.reports"}"

RUN mkdir -p "${REPODIR}" "${OUTDIR}"
COPY . "${REPODIR}"
WORKDIR "${REPODIR}"

### Run hadolint ...
RUN ( find . -type f -name '*Dockerfile*' | \
      xargs hadolint --format json > "${OUTDIR}/hadolint.json" ) || true
RUN ls -la "${OUTDIR}"

### Convert hadolint JSON to SARIF ...
RUN go run "${TOOLDIR}/hadolint/cmd/main.go" \
        < "${OUTDIR}/hadolint.json" \
        > "${OUTDIR}/hadolint.sarif"
RUN ls -la "${OUTDIR}"
