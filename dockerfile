
FROM centos:centos7

ADD dtu_service /home/dtu_service

WORKDIR /home/

EXPOSE 3333
EXPOSE 3334


CMD "-M=15"

ENTRYPOINT /home/dtu_service

#docker run -tdi -p3333:3333 -p3334:3334  --name="dtu_servicev1.0"  dtu_service:v1.0
