FROM golang:alpine3.18 as build
WORKDIR /app/
COPY . .
RUN apk add --update make
RUN make build



FROM busybox:stable
WORKDIR /app/
RUN mkdir ./html/
COPY --from=build /app/bin/main ./
COPY --from=build /app/html/* ./html/
EXPOSE 8080
CMD [ "./main" ]
