FROM golang:1.11.0-alpine3.7 as test

VOLUME /repo
RUN cp /repo/.rocro/test-1.sarif.json /repo/.rocro/task-results.test-1.sarif

CMD true
