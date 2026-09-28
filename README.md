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
go mod init train-go-standard
```

### 2. Executar Código Diretamente
Útil para desenvolvimento e testes rápidos (compila e executa em memória):

```bash
# Executa a partir da raiz do projeto:
go run ./standard/main.go

# Ou acessando o diretório do módulo:
cd standard
go run main.go
```

### 3. Compilar um Executável
Gera um arquivo binário compilado independente:

```bash
# Compilar a partir da pasta standard/
cd standard
go build main.go

# Executar o binário gerado:
# Linux/macOS:
./main

# Windows (PowerShell):
.\main.exe
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
├── standard/                  # Módulo com fundamentos da linguagem Go
│   ├── main.go                # Ponto de entrada do programa (package main)
│   ├── types/                 # Modelagem e tipos do módulo standard
│   │   └── user.go            # Struct User, construtor com UUID, methods e Stringer
│   └── go.mod                 # Definição do módulo Go (module train-go-standard)
├── types/                     # Exemplos comparativos de funções e ponteiros
│   └── types.go               # Demonstração de funções com value vs pointer receiver
├── .gitignore                 # Arquivos ignorados pelo Git
└── README.md                  # Guia de comandos e referências de estudo
```

---

## 🧠 Guia de Conceitos: Tipos, Structs e Métodos

### 1. Structs e Visibilidade (`standard/types/user.go`)
Em Go, não existem classes; usamos **structs** para definir estruturas de dados:

- **Campos Públicos (Exportados)**: Primeira letra **maiúscula** (ex.: `Id`, `Name`, `Email`). Podem ser acessados por outros pacotes.
- **Campos Privados (Não-exportados)**: Primeira letra **minúscula** (ex.: `password`). Acessíveis apenas dentro do próprio pacote `types`.

```go
type User struct {
	Id       uuid.UUID // Campo público
	Name     string    // Campo público
	Email    string    // Campo público
	password string    // Campo privado
}
```

### 2. Funções vs Métodos com Receiver
Em Go, funções podem ser associadas diretamente a uma struct através de **receivers**, comportando-se como métodos:

- **Função comum (`types/types.go`)**:
  ```go
  func ChangeEmail(user *User, email string) {
      user.Email = email
  }
  ```
- **Método com Receiver (`standard/types/user.go`)**:
  ```go
  func (user *User) ChangeEmail(email string) {
      user.Email = email
  }
  ```

### 3. Passagem por Valor vs Ponteiro (`User` vs `*User`)
- **Value Receiver (`func (user User) ...` ou `func GetName(user User)`)**:
  - Cria uma cópia da struct na memória.
  - Modificações na cópia **não afetam** o dado original.
  - Menos eficiente para structs médias ou grandes.
- **Pointer Receiver (`func (user *User) ...`)**:
  - Recebe o ponteiro para o endereço de memória original.
  - Permite alterar os dados da struct (`ChangeName`, `ChangePassword`, `ChangeEmail`).
  - Mais eficiente, pois evita a cópia completa dos campos.

### 4. Padrão Construtor (`NewUser`)
Go não possui construtores automáticos. A convenção idiomática é criar funções com prefixo `New...` que retornam uma nova instância (geralmente um ponteiro para a struct):

```go
func NewUser(name string, email string, password string) *User {
	return &User{
		Id:       uuid.New(),
		Name:     name,
		Email:    email,
		password: password,
	}
}
```

### 5. Customização de Saída com `fmt.Stringer` (`String()`)
Ao implementar o método `String() string`, a struct atende à interface nativa `fmt.Stringer`. Dessa forma, funções como `fmt.Println(user)` exibem a formatação customizada:

```go
func (user *User) String() string {
	return fmt.Sprintf("{\n\tId: \t%s\n\tName: \t%s\n\tEmail: \t%s\n}", user.Id, user.Name, user.Email)
}
```