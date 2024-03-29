FROM golang:1.19-alpine

#RUN mkdir /Agenda
WORKDIR /Agenda

COPY . /Agenda


RUN go mod init Agenda
RUN go mod tidy
RUN go build cmd/Agenda/main.go

EXPOSE 8080

CMD [ "./main" ]

