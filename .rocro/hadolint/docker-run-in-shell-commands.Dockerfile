FROM hadolint/hadolint:v2.4.0 AS linter-stage

### Install tools ...
RUN apk add --update --no-cache curl git

### Install hadolint ...
ENV HADOLINT_ARCHDIR="https://github.com/hadolint/hadolint/releases/download" \
    HADOLINT_VERSION="v1.18.0"
RUN echo "+++ ${HADOLINT_ARCHDIR}/${HADOLINT_VERSION}/hadolint-$(uname -s)-$(uname -m)"
RUN curl -sL -o /usr/bin/hadolint \
         "${HADOLINT_ARCHDIR}/${HADOLINT_VERSION}/hadolint-$(uname -s)-$(uname -m)" \
 && chmod 755 /usr/bin/hadolint

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
SHELL ["/bin/ash", "-eo", "pipefail", "-c"]
RUN ( find . -type f -name '*Dockerfile*' -print0 | xargs -0 hadolint --format json ) \
        > "${OUTDIR}/hadolint.json" || true
RUN ls -la "${OUTDIR}"

### Convert hadolint JSON to SARIF ...
RUN go run "${TOOLDIR}/hadolint/cmd/main.go" "${REPOPATH}" \
        < "${OUTDIR}/hadolint.json" \
        > "${OUTDIR}/hadolint.sarif"
RUN ls -la "${OUTDIR}"