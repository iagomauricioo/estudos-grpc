# API de Produtos com gRPC

Este projeto implementa um serviço de produtos usando gRPC em Go.

## Pré-requisitos

- Go 1.24.5 ou superior
- Protocol Buffers (protoc) instalado
- Plugins do Go para protoc

## Instalação das Dependências

### 1. Instalar os plugins do Go para Protocol Buffers

```bash
# Instalar o plugin para gerar código Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Instalar o plugin para gerar código gRPC
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 2. Adicionar os binários ao PATH

```bash
# Adicionar o diretório bin do Go ao PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

## Estrutura do Projeto

```
api-products/
├── src/
│   ├── proto/
│   │   └── product-service.proto    # Definição do serviço
│   └── pb/
│       └── products/                # Código Go gerado
│           ├── product-service.pb.go
│           └── product-service_grpc.pb.go
├── main.go                          # Servidor gRPC
├── go.mod                           # Dependências do Go
└── README.md                        # Este arquivo
```

## Como Usar

### 1. Gerar código Go a partir do .proto

```bash
# Na pasta api-products_unary
protoc src/proto/*.proto --go_out=./ --go-grpc_out=./
```

### 2. Atualizar dependências

```bash
go mod tidy
```

### 3. Executar o servidor

```bash
go run main.go
```

O servidor estará rodando na porta 50051.

## Definição do Serviço

O serviço `ProductService` oferece duas operações:

- **Create**: Cria um novo produto
- **FindAll**: Lista todos os produtos

### Estrutura do Produto

```protobuf
message Product {
    int32 id = 1;
    string name = 2;
    string description = 3;
    double price = 4;
    int32 quantity = 5;
}
```

## Comandos Úteis

```bash
# Verificar se o protoc está instalado
which protoc

# Verificar se os plugins Go estão no PATH
which protoc-gen-go
which protoc-gen-go-grpc

# Verificar a versão do Go
go version

# Verificar o GOPATH
go env GOPATH
```

## Solução de Problemas

### Erro: "protoc-gen-go: program not found"

Este erro ocorre quando os plugins do Go não estão no PATH. Execute:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Erro: "protoc: command not found"

Instale o Protocol Buffers:

```bash
# Ubuntu/Debian
sudo apt-get install protobuf-compiler

# Ou baixe de https://github.com/protocolbuffers/protobuf/releases
```

## Próximos Passos

1. Implementar o servidor gRPC no `main.go`
2. Criar um cliente para testar o serviço
3. Adicionar validações e tratamento de erros
4. Implementar persistência de dados
5. Adicionar testes unitários 