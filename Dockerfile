FROM golang:1.21.6

WORKDIR /app

COPY /bin/templ-starter /templ-starter

EXPOSE 8080

CMD ["/templ-starter"]