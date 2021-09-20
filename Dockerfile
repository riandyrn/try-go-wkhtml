FROM golang:1.15.6-alpine3.12 as build
WORKDIR /go/src/github.com/riandyrn/try-go-wkhtml

COPY . .

WORKDIR /go/src/github.com/riandyrn/try-go-wkhtml
RUN go build -o app

FROM madnight/docker-alpine-wkhtmltopdf:0.12.5-alpine3.10 as wkbuild

FROM alpine:3.12
RUN apk add \
    libgcc libstdc++ libx11 glib libxrender libxext libintl \
    ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family \
    ca-certificates tzdata

COPY --from=build /go/src/github.com/riandyrn/try-go-wkhtml .
COPY --from=build /go/src/github.com/riandyrn/try-go-wkhtml/Lateef-Regular.ttf /tmp/Lateef-Regular.ttf
COPY --from=build /go/src/github.com/riandyrn/try-go-wkhtml/LinLibertine_R.ttf /tmp/LinLibertine_R.ttf
COPY --from=wkbuild ./bin/wkhtmltopdf .

ENTRYPOINT "./app"