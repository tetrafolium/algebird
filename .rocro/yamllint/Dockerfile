FROM python:alpine AS linter-stage

### Install golang ...
RUN apk add --update --no-cache go && \
    echo "+++ $(go version)"

ENV GOBIN="$GOROOT/bin" GOPATH="/.go"
ENV PATH="${GOPATH}/bin:/usr/local/go/bin:$PATH"

### Install yamllint tool ...
RUN pip3 install 'yamllint>=1.26.0,<1.27.0' && \
    echo "+++ $(yamllint --version)"

ENV REPOPATH="bitbucket.org/tetrafolium/algebird" \
    TOOLPATH="github.com/tetrafolium/inspecode-tasks"
ENV REPODIR="${GOPATH}/src/${REPOPATH}" \
    TOOLDIR="${GOPATH}/src/${TOOLPATH}"

ARG OUTDIR
ENV OUTDIR="${OUTDIR:-"/.reports"}"

RUN mkdir -p "${REPODIR}" "${OUTDIR}"
COPY . "${REPODIR}"
WORKDIR "${REPODIR}"

### Put symlink refers submodule-path at origin-path
RUN mkdir -p "$(dirname "${TOOLDIR}")" && \
    ln -s "${REPODIR}/$(basename "${TOOLPATH}")" "${TOOLDIR}"

### Run yamllint ...
RUN ( yamllint -f parsable . > "${OUTDIR}/yamllint.issues" || true ) && \
    ls -la "${OUTDIR}"

### Convert yamllint issues to SARIF ...
RUN GO111MODULE="off" \
    go run "${TOOLDIR}/yamllint/cmd/main.go" "${REPOPATH}" \
        < "${OUTDIR}/yamllint.issues" \
        > "${OUTDIR}/yamllint.sarif" && \
    ls -la "${OUTDIR}"
