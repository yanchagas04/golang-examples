# 🐹 Train Go — Repositório de Estudos

Repositório dedicado ao aprendizado prático da linguagem **Go (Golang)**, contendo anotações, exemplos básicos e boas práticas da linguagem.

---

## 🚀 Comandos Essenciais

### 1. Inicializar um Módulo
O arquivo `go.mod` define o caminho do módulo e a versão do Go utilizada pelo projeto:

```bash
# Inicializa um novo módulo Go
go mod init <nome-do-modulo>

# Exemplo:
go mod init train-go
```

### 2. Executar Código Diretamente
Útil para desenvolvimento e testes rápidos (compila e executa em memória):

```bash
# Executa um arquivo Go diretamente
go run standard/hello_world.go

# Ou acessando o diretório do módulo:
cd standard
go run hello_world.go
```

### 3. Compilar um Executável
Gera um arquivo binário compilado independente:

```bash
# Compilar a partir da pasta standard/
cd standard
go build hello_world.go

# Executar o binário gerado:
# Linux/macOS:
./hello_world

# Windows (PowerShell):
.\hello_world.exe
```

### 4. Manutenção e Qualidade de Código

```bash
# Formata automaticamente o código conforme os padrões do Go
go fmt ./...

# Analisa o código em busca de erros comuns e más práticas
go vet ./...

# Sincroniza e limpa dependências do go.mod / go.sum
go mod tidy

# Executa testes unitários (quando houver arquivos *_test.go)
go test ./...
```

---

## 📁 Estrutura do Projeto

```text
train-go/
├── standard/          # Módulo com os fundamentos da linguagem Go
│   ├── hello_world.go # Entrada básica do programa (package main)
│   └── go.mod         # Definição do módulo Go (module standard)
├── .gitignore         # Arquivos ignorados pelo Git
└── README.md          # Guia de comandos e referências de estudo
```