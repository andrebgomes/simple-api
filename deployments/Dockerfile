# use golang official image and build the project into a binary called 'app'
FROM golang:1.18 AS simple-app
# create new dir called app
WORKDIR /app
# copy needed files from project into the new dir
COPY go.mod .
COPY main.go .
RUN go mod tidy
# build the project
RUN go build -o app

# in order to have a much simpler and lighter image, only use ubuntu image as final
# use the ubuntu official image and copy the binary project from the previous image into this one
FROM ubuntu
WORKDIR /app
COPY --from=simple-app /app/app app
COPY config.yaml .
EXPOSE 8080

CMD ["./app", "-conf", "config.yaml"]