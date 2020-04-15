FROM ubuntu:18.04
RUN apt-get update -y && apt-get install -y openjdk-11-jre-headless curl debconf-utils
RUN curl --location --output /tmp/jmxtrans.deb https://repo1.maven.org/maven2/org/jmxtrans/jmxtrans/271/jmxtrans-271.deb
RUN echo "jmxtrans jmxtrans/jvm_heap_size string 512" | debconf-set-selections
RUN apt-get install -y /tmp/jmxtrans.deb
RUN mkdir -p /etc/jmxtrans
ADD jmxtrans.json /etc/jmxtrans/jmxtrans.json
CMD /usr/share/jmxtrans/bin/jmxtrans.sh /etc/jmxtrans/jmxtrans.json
