FROM golang:latest AS build_base
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o free-tsa-bench .

FROM adoptopenjdk/openjdk8:alpine
COPY --from=build_base /app/free-tsa-bench /free-tsa-bench
COPY --from=build_base /app/resources/adioss.p12 /adioss.p12
COPY --from=build_base /app/resources/Generated_pattern.java /Generated_pattern.java
ENTRYPOINT ["./free-tsa-bench"]