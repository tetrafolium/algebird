FROM ubuntu:18.04

# Install dependencies
RUN apt-get update && \
 apt-get -y install apache2

# Install apache and write yamllint issues
RUN echo './rocro.yaml:18:81: [error] line too long (92 > 80 characters) (line-length)' > /var/www/html/yamllint-issues.txt && \
    echo './.travis.yml:1:1: [warning] missing document start "---" (document-start)'  >> /var/www/html/yamllint-issues.txt

RUN echo './algebird-util/src/main/scala/com/twitter/algebird/util/summer/HeavyHittersCachingSummer.scala:30:73: "freqeuncy" is a misspelling of "frequency"' > /var/www/html/misspell-issues.txt && \
    echo './algebird-util/src/main/scala/com/twitter/algebird/util/summer/AsyncListMMapSum.scala:29:20: "asyncronous" is a misspelling of "asynchronous"'    >> /var/www/html/misspell-issues.txt && \
    echo './algebird-core/src/main/scala/com/twitter/algebird/Fold.scala:251:28: "inferrence" is a misspelling of "inference"'                               >> /var/www/html/misspell-issues.txt && \
    echo './algebird-core/src/main/scala/com/twitter/algebird/Eventually.scala:117:5: "Overriden" is a misspelling of "Overridden"'                          >> /var/www/html/misspell-issues.txt && \
    echo './algebird-core/src/main/scala/com/twitter/algebird/SummingIterator.scala:32:54: "individiual" is a misspelling of "individual"'                   >> /var/www/html/misspell-issues.txt


# Configure apache
RUN echo '. /etc/apache2/envvars'           > /root/run_apache.sh && \
    echo 'mkdir -p /var/run/apache2'       >> /root/run_apache.sh && \
    echo 'mkdir -p /var/lock/apache2'      >> /root/run_apache.sh && \ 
    echo '/usr/sbin/apache2 -D FOREGROUND' >> /root/run_apache.sh && \ 
    chmod 755 /root/run_apache.sh

EXPOSE 80

CMD /root/run_apache.sh
