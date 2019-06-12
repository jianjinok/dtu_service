
FROM centos:centos7

ADD dtu_service /home/dtu_service
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /home/

EXPOSE 5550
EXPOSE 5551

ENTRYPOINT ["/home/dtu_service"]
CMD ["-M=15"]

#docker run -tdi -p5550:5550 -p5551:5551  --name="dtu_servicev1.x"  dtu_service:v1.x
