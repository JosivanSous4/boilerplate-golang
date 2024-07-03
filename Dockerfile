# Usando uma imagem oficial do Golang como imagem base
FROM golang:1.21.4

# Setar o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos go.mod e go.sum e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# Copiar todos os arquivos do projeto para o diretório de trabalho
COPY . .

# Para desenvolvimento
RUN go install github.com/cosmtrek/air@v1.44.0

# Compilar o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/rest

ENV TZ=America/Sao_Paulo

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
