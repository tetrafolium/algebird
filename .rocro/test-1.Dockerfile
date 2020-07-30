FROM alpine:3.7 as test

COPY . /.repo
RUN ls -laR /.repo
RUN mkdir -p /.rocro
RUN cp /.repo/test-1.sarif.json /.rocro/inspecode.sarif
RUN ls -laR /.rocro

#CMD true
