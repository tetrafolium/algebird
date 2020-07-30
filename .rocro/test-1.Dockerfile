FROM alpine:3.7 as test

RUN mkdir -p /.rocro
COPY ./test-1.sarif.json /.rocro/inspecode.sarif
RUN ls -laR /.rocro

#CMD true
