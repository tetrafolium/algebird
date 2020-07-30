FROM golang:1.11.0-alpine3.7 as test

COPY . /repo
WORKDIR /repo
RUN ls -laR
RUN mkdir -p task-results
RUN cp test-1.sarif.json task-results/test-1.sarif
RUN ls -laR

CMD true
