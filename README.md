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
│   ├── main.go                # Ponto de entrada, instâncias de tipos e fluxo de erros
│   ├── types/                 # Modelagem e tipos do módulo standard
│   │   └── user.go            # Struct User, métodos (Set*), construtor, Stringer e validações
│   └── go.mod                 # Definição do módulo Go (module train-go-standard)
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

- **Função comum**:
  ```go
  func SetEmail(user *User, email string) {
      user.Email = email
  }
  ```
- **Método com Receiver (`standard/types/user.go`)**:
  ```go
  func (user *User) SetEmail(email string) {
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
  - Permite alterar os dados da struct (`SetName`, `SetPassword`, `SetEmail`).
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

---

## ⚠️ Guia de Conceitos: Tratamento de Erros

Diferente de outras linguagens (como Java, Python ou C#), **Go não possui exceções (`try/catch/throw`)**. Em vez disso, erros são tratados explicitamente como **valores de retorno** através da interface nativa `error`.

### 1. Retorno Múltiplo e a Interface `error`
Funções ou métodos propensos a falhas retornam o resultado esperado juntamente com um valor do tipo `error` (por convenção, como último valor retornado na assinatura):

```go
// Assinatura: (resultado, error)
func (user *User) Greet(other_person string) (string, error)
```

- **Em caso de Sucesso**: Retorna o dado esperado e `nil` no erro.
- **Em caso de Falha**: Retorna o *zero value* do dado (ex.: `""`, `0`, `nil`) e a instância do erro.

### 2. Criando Erros com `errors.New` (`standard/types/user.go`)
Para instanciar erros simples com mensagens descritivas, utiliza-se a função `errors.New` do pacote nativo `errors`:

```go
// Greet recebe o nome de quem será saudado e retorna uma mensagem ou erro caso o nome esteja vazio
func (user *User) Greet(other_person string) (string, error) {
	if other_person == "" {
		return "", errors.New("The other person's name cannot be empty!")
	}
	return fmt.Sprintf("Hello, %s, from %s!", other_person, user.Name), nil
}
```

### 3. Verificação Idiomática (`if err != nil`)
Em Go, o fluxo de controle de erros é explícito: deve-se verificar o erro imediatamente após a chamada da função:

```go
msg, err := user.Greet("")
if err != nil {
	// Tratamento do erro (ex.: logar, tentar novamente, encerrar, etc.)
}
```

> [!NOTE]
> Por convenção idiomática na comunidade Go, costuma-se nomear a variável de erro como `err` para evitar sombreamento (*shadowing*) do tipo primitivo `error`.

### 4. Logging e Interrupção com `log.Fatal` (`standard/main.go`)
No ponto de entrada (`main.go`), o pacote `log` é utilizado para configurar mensagens padronizadas e interromper a execução diante de um erro impeditivo:

```go
// Configuração do Logger:
log.SetPrefix("train-go: ") // Adiciona um prefixo customizado nas mensagens de log
log.SetFlags(0)             // Remove carimbos de data/hora (flags padrão)

// Execução e tratamento:
msg, error := user.Greet("")
if error != nil {
	log.Fatal(error) // Exibe a mensagem de erro formatada no stderr e encerra o programa (os.Exit(1))
}
fmt.Println(msg)
```

#### Saída exibida no terminal em caso de erro:
```text
train-go: The other person's name cannot be empty!
exit status 1
```