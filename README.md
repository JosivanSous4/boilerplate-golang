# Boilerplate Golang

Esqueleto de um projeto em golang usando Clear Arch.

### Rodando local

Para rodar o projeto, rode os comandos abaixo

```bash
docker-compose up --build -d
docker exec localstack awslocal sqs create-queue --queue-name product_queue
```

## Implementações

- MongoDB e MySQL (Gorm)
- Fiber
- JWT
- Hash e Compare de senha
- SQS e RabbitMQ
- Middlware
- Tratamento de erros

### Documentação da API

#### Adiciona produto

```http
  POST /products
```

```bash
{
    "id": "10",
    "name": "Primeiro",
    "description": "Primeiro produto",
    "price": 50
}
```

| Parâmetro     | Tipo     | Descrição        |
| :------------ | :------- | :--------------- |
| `id`          | `string` | **Obrigatório**. |
| `name`        | `string` | **Obrigatório**. |
| `description` | `string` | **Obrigatório**. |
| `price`       | `number` | **Obrigatório**. |

#### Retorna um produto

```http
  GET /products/${id}
```

```bash
{
    "id": "10",
    "name": "Primeiro",
    "description": "Primeiro produto",
    "price": 50
}
```
