FROM golang

WORKDIR $GOPATH/src/github.com/dvbnrg/todoAPI

COPY . .

RUN go get -d -v ./...
# RUN export GO111MODULE=on  
# RUN go mod download
# RUN go mod verify

RUN go install -v ./...
RUN go build -v ./...

# RUN ls -laR

CMD [ "go", "run", "main.go" ]

EXPOSE 8080