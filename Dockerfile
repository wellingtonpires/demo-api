#BASE IMAGE
FROM golang:alpine

#INSTALL NEOFETCH
RUN apk add neofetch

#COPY FILES
COPY go.mod /usr/src/app/
COPY demoapi.go /usr/src/app/

#CONTAINER PORT
EXPOSE 8080

#RUN APPLICATION
CMD ["go", "run", "/usr/src/app/demoapi.go"]

