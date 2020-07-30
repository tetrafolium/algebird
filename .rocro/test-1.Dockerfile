FROM golang:1.11.0-alpine3.7 as test

COPY . /repo
WORKDIR /repo
RUN ls -la
#RUN mkdir -p .rocro/task-results
#RUN cp .rocro/test-1.sarif.json .rocro/task-results/test-1.sarif

CMD true
