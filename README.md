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
├── standard/                  # Módulo com os fundamentos da linguagem Go
│   ├── hello_world.go         # Entrada básica do programa (package main)
│   ├── types/                 # Tipos, structs e modelagem
│   │   ├── simple_type.go     # Structs, visibilidade, mutabilidade e construtores
│   │   └── complex_types.go   # Composição de structs (embedding/agregação)
│   └── go.mod                 # Definição do módulo Go (module standard)
├── .gitignore                 # Arquivos ignorados pelo Git
└── README.md                  # Guia de comandos e referências de estudo
```

---

## 🧠 Guia de Conceitos: Tipos e Structs (`types/`)

### 1. Structs e Visibilidade (`simple_type.go`)
Em Go, não existem classes; usamos **structs** para definir estruturas de dados:

- **Campos Públicos (Exportados)**: Primeira letra **maiúscula** (ex.: `ID`, `Name`, `Email`). Podem ser acessados por outros pacotes.
- **Campos Privados (Não-exportados)**: Primeira letra **minúscula** (ex.: `password`). Acessíveis apenas dentro do próprio pacote `types`.

### 2. Passagem por Valor vs Ponteiro (`User` vs `*User`)
- **Passagem por Valor (`user User`)**:
  - Cria uma cópia da struct na memória.
  - Modificações na cópia **não afetam** o dado original.
  - Menos eficiente em memória para structs de tamanho médio/grande.
- **Passagem por Ponteiro (`user *User`)**:
  - Passa o endereço de memória (`*`).
  - Permite alterar os dados da struct original (`ChangeName`, `ChangePassword`, `ChangeEmail`).
  - Mais eficiente, pois evita copiar os dados.

### 3. Padrão Construtor (`NewUser`, `NewAddress`, `NewPerson`)
Go não possui palavras-chave como `class` ou construtores automáticos. A convenção idiomática é criar funções com prefixo `New...`:

```go
// Retorna um ponteiro para a struct recém-criada
func NewUser(id int, name string, email string, password string) *User {
    return &User{
        ID:       id,
        Name:     name,
        Email:    email,
        password: password,
    }
}
```

### 4. Composição de Structs (`complex_types.go`)
Ao invés de herança tradicional orientada a objetos, Go adota **composição**:
- Uma struct pode conter outra struct como campo (ex.: `Person` contém um campo `Address Address`).
- Funções construtoras podem instanciar e encadear estruturas compostas:

```go
func NewPerson(name, email, street, city, state, zip string) Person {
    return Person{
        Name:    name,
        Email:   email,
        Address: NewAddress(street, city, state, zip),
    }
}
```