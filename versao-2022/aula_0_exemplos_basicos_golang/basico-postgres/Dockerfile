# Imagem base oficial do Go
FROM golang:1.22.2 as builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o restante dos arquivos do código-fonte para o diretório de trabalho
COPY . .

# Compila o aplicativo para um binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Utiliza uma imagem Docker leve para o ambiente de execução
FROM alpine:latest

# Define o diretório de trabalho
WORKDIR /root/

# Copia o binário do primeiro estágio (builder)
COPY --from=builder /app/main .


COPY --from=builder /app/templates ./templates

# Expõe a porta que o aplicativo utilizará
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./main"]
