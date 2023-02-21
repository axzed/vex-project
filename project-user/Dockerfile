FROM alpine
WORKDIR /Initial
COPY ./target/project-user .
COPY ./config/app-docker.yaml .
RUN  mkdir config && mv app-docker.yaml config/app.yaml
EXPOSE 8080 8881
ENTRYPOINT ["./project-user"]