# gRPC com Go - Estudos e Exemplos

Este repositório contém estudos práticos sobre gRPC (Google Remote Procedure Call) implementados em Go, demonstrando diferentes padrões de comunicação e arquiteturas.

## O que é gRPC?

gRPC é um framework de comunicação RPC (Remote Procedure Call) de alto desempenho desenvolvido pelo Google. Ele permite que aplicações se comuniquem de forma eficiente através de diferentes linguagens e plataformas usando Protocol Buffers como linguagem de definição de interface.

### Características principais:

- **Protocol Buffers**: Usa `.proto` files para definir contratos de serviço
- **Multi-language**: Suporte para múltiplas linguagens de programação
- **HTTP/2**: Utiliza HTTP/2 como protocolo de transporte
- **Streaming**: Suporte para streaming unidirecional e bidirecional
- **Code Generation**: Gera código automaticamente a partir dos arquivos `.proto`
- **Type Safety**: Tipagem forte e validação automática

### Padrões de Comunicação gRPC:

1. **Unary RPC**: Cliente envia uma requisição e recebe uma resposta
2. **Server Streaming RPC**: Cliente envia uma requisição e recebe um stream de respostas
3. **Client Streaming RPC**: Cliente envia um stream de requisições e recebe uma resposta
4. **Bidirectional Streaming RPC**: Ambos enviam streams de mensagens

## Estrutura do Repositório

```
gRPC/
├── unary/                           # Exemplos de comunicação unária
│   ├── api-products_unary/          # Servidor de produtos (CRUD completo)
│   └── client_api-products_unary/   # Cliente para testar o servidor
├── client-streaming/                # Exemplos de client streaming
│   ├── api-calc/                   # Servidor de cálculos
│   └── client_api-calc/            # Cliente para testar cálculos
└── docker-compose.yml              # Configuração do banco PostgreSQL
```

## Módulos Detalhados

### 1. Unary RPC - API de Produtos

**Localização**: `unary/api-products_unary/`

Este módulo demonstra um CRUD completo de produtos usando comunicação unária gRPC com persistência em PostgreSQL.

#### Funcionalidades:
- ✅ **Create**: Criar novos produtos
- ✅ **Read**: Buscar produto por ID e listar todos
- ✅ **Update**: Atualizar produtos existentes
- ✅ **Delete**: Remover produtos
- ✅ **Persistência**: Banco PostgreSQL com GORM
- ✅ **Mapeamento**: Conversão entre protobuf e GORM

#### Arquitetura:
```
src/
├── proto/
│   └── product-service.proto       # Definição do serviço
├── pb/products/                    # Código gerado pelo protoc
├── model/
│   └── product.go                  # Modelo GORM
├── repository/
│   └── product-repository.go      # Camada de acesso a dados
├── mapper/
│   └── grpc-gorm-mapper.go        # Conversão protobuf ↔ GORM
└── config/
    └── database.go                 # Configuração do banco
```

#### Definição do Serviço:
```protobuf
service ProductService {
    rpc Create(Product) returns (Product);
    rpc FindAll(google.protobuf.Empty) returns (ProductList);
    rpc FindById(ProductId) returns (Product);
    rpc Update(Product) returns (Product);
    rpc Delete(ProductId) returns (google.protobuf.Empty);
}
```

#### Como executar:

1. **Iniciar o banco de dados**:
```bash
docker-compose up -d
```

2. **Gerar código protobuf**:
```bash
cd unary/api-products_unary
protoc src/proto/*.proto --go_out=./ --go-grpc_out=./
```

3. **Executar o servidor**:
```bash
go run main.go
```

4. **Testar com o cliente**:
```bash
cd ../client_api-products_unary
go run main.go
```

### 2. Client Streaming RPC - Calculadora

**Localização**: `client-streaming/`

Este módulo demonstra o padrão de client streaming, onde o cliente envia múltiplos valores e o servidor processa todos para retornar um resultado final.

#### Funcionalidades:
- ✅ **Client Streaming**: Cliente envia stream de números
- ✅ **Processamento**: Servidor calcula soma, média e quantidade
- ✅ **Resposta única**: Retorna resultado consolidado

#### Definição do Serviço:
```protobuf
service CalcService {
    rpc Calc(stream Input) returns (Output);
}
```

#### Como funciona:
1. Cliente envia números sequencialmente: `[1, 2, 3]`
2. Servidor processa cada número recebido
3. Quando o stream termina, retorna estatísticas:
   - Quantidade de números
   - Soma total
   - Média aritmética

#### Como executar:

1. **Executar o servidor**:
```bash
cd client-streaming/api-calc
go run main.go
```

2. **Executar o cliente**:
```bash
cd ../client_api-calc
go run main.go
```

## Pré-requisitos

### Ferramentas necessárias:

1. **Go** (versão 1.24.5 ou superior)
2. **Protocol Buffers** (protoc)
3. **Plugins do Go para protoc**

### Instalação:

```bash
# Instalar protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Instalar protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Adicionar ao PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

### Dependências do sistema:

```bash
# Ubuntu/Debian
sudo apt-get install protobuf-compiler

# macOS
brew install protobuf
```

## Tecnologias Utilizadas

- **gRPC**: Framework de comunicação RPC
- **Protocol Buffers**: Serialização de dados
- **Go**: Linguagem de programação
- **PostgreSQL**: Banco de dados relacional
- **GORM**: ORM para Go
- **Docker**: Containerização do banco de dados

## Conceitos Demonstrados

### 1. Protocol Buffers
- Definição de mensagens e serviços
- Geração automática de código
- Serialização eficiente

### 2. Padrões de Comunicação
- **Unary**: Requisição-resposta simples
- **Client Streaming**: Múltiplas requisições, uma resposta

### 3. Arquitetura de Software
- Separação de responsabilidades
- Repository pattern
- Mapeamento entre camadas
- Configuração de banco de dados

### 4. Integração com Banco de Dados
- ORM com GORM
- Migração automática
- Operações CRUD

## Comandos Úteis

### Gerar código protobuf:
```bash
protoc src/proto/*.proto --go_out=./ --go-grpc_out=./
```

### Atualizar dependências:
```bash
go mod tidy
```

### Verificar instalação:
```bash
# Verificar protoc
which protoc

# Verificar plugins Go
which protoc-gen-go
which protoc-gen-go-grpc

# Verificar versão do Go
go version
```

## Próximos Passos

1. **Server Streaming RPC**: Implementar streaming do servidor para o cliente
2. **Bidirectional Streaming**: Comunicação bidirecional em tempo real
3. **Interceptors**: Middleware para logging, autenticação, etc.
4. **Load Balancing**: Distribuição de carga entre múltiplos servidores
5. **Service Discovery**: Descoberta automática de serviços
6. **TLS/SSL**: Comunicação segura com certificados
7. **Testes**: Testes unitários e de integração
8. **Monitoramento**: Métricas e observabilidade
