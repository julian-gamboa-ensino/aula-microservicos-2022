# Estágio de build
FROM alpine:latest AS build

# Instale o Go
RUN apk add --no-cache go

WORKDIR /build

# Copie as dependências
COPY go.mod go.sum ./
RUN go mod download && go mod tidy

# Baixe as dependências (opcional)
RUN go get github.com/aws/aws-lambda-go/events
RUN go get github.com/aws/aws-lambda-go/lambda
RUN go get github.com/gin-gonic/gin


# Copie o código-fonte
COPY *.go ./

# Compile o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o main
RUN go clean -modcache

# Estágio de execução
FROM alpine:latest

WORKDIR /var/task

# Copie o binário do estágio de build
COPY --from=build /build/main ./

# Comando para executar a função
CMD ["./main"]
