
# 🧩 WebAssembly Templ UI – Go + Templ + Alpine.js (WASM Frontend)

Este projeto é uma aplicação experimental que utiliza **Go** compilado para **WebAssembly (WASM)**, com **Templ** para renderização de componentes HTML e **Alpine.js** para interatividade leve no navegador — sem uso de frameworks JS pesados como React, Vue ou Angular. Todo o conteúdo da interface é gerado dinamicamente no **navegador**, via `.wasm`, e injetado diretamente no DOM.

## 🚀 Tecnologias Utilizadas

- **Go** – linguagem principal do projeto, usada tanto no backend quanto no frontend via WASM.
- **Templ** – engine de templates moderna, tipada e integrada com o compilador Go.
- **WebAssembly (GOOS=js GOARCH=wasm)** – transforma o frontend em um binário `.wasm` que roda no navegador.
- **Alpine.js** – adiciona reatividade leve e controle de interface.
- **TailwindCSS** – para estilização rápida e moderna baseada em utilitários.

## 📁 Estrutura do Projeto

A estrutura do projeto é dividida entre arquivos de interface (UI), servidor HTTP, recursos estáticos e o módulo WebAssembly:

- `assets/`: arquivos estáticos como CSS, wasm_exec.js e o main.wasm gerado.
- `ui/layouts/`: contém o layout base da aplicação (HTML shell).
- `ui/modules/`: componentes reutilizáveis como Navbar, Footer, etc.
- `ui/pages/`: páginas individuais como Landing.
- `wasm/`: código Go compilado em WebAssembly.
- `main.go`: servidor HTTP responsável por servir o HTML e os assets.
- `Makefile`: automatização das etapas de build.
- `go.mod`: gerenciamento de dependências.

## ⚙️ Como Rodar Localmente

1. Instale as dependências:
   - Go 1.21+ instalado
   - CLI do Templ: `go install github.com/a-h/templ/cmd/templ@latest`

2. Gere os arquivos Templ:
   ```
   templ generate
   ```

3. Compile o WebAssembly:
   ```
   GOOS=js GOARCH=wasm go build -o ./assets/main.wasm ./wasm/wasm_main.go
   ```

4. Rode o servidor:
   ```
   go run main.go
   ```

Abra o navegador em [http://localhost:8090](http://localhost:8090) para visualizar a aplicação.

## 🧠 Como Funciona

- O servidor Go responde com um HTML base (`BaseLayout`) que inclui um `#app`.
- O navegador carrega `wasm_exec.js` e executa `main.wasm`.
- O WebAssembly (Go) renderiza o conteúdo HTML via Templ e injeta diretamente no DOM.
- A navegação e comunicação entre componentes pode ser feita com Alpine.js ou chamadas diretas via Go (exportando funções para o `window`).

## 🔧 Exemplo de Injeção via WASM

```go
var buf bytes.Buffer
_ = pages.Landing().Render(context.Background(), &buf)
js.Global().Get("document").Call("getElementById", "app").Set("innerHTML", buf.String())
```

## ✅ Benefícios

- Aplicação 100% dinâmica no frontend com performance nativa.
- Nenhum framework JS pesado.
- Componentes reativos com Alpine.js.
- Total controle do DOM com segurança de tipos em Go.
- Preparado para evoluir para SPA completa com rotas, formulários, plugins e muito mais.

## 📄 Licença

MIT © Jeferson Rafael Marques
