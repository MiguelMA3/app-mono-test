# Desafio Técnico: Perfil de Usuário e Timeline (Nível Júnior)

Esta é a implementação do desafio técnico para a posição de desenvolvedor, correspondente ao nível Júnior. A aplicação consiste em uma API RESTful em Go e um frontend em React, que juntos formam uma pequena rede social com perfis de usuário, posts e curtidas.

## Sumário
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Como Configurar e Rodar a Aplicação](#como-configurar-e-rodar-a-aplicação)
  - [Pré-requisitos](#pré-requisitos)
  - [Backend (API em Go)](#1-backend-api-em-go)
  - [Frontend (Aplicação em React)](#2-frontend-aplicação-em-react)
- [Acessando a Documentação (Swagger)](#acessando-a-documentação-swagger)
- [Executando os Testes](#executando-os-testes)
  - [Testes do Backend](#testes-do-backend-api-em-go)
  - [Testes do Frontend](#testes-do-frontend-react)
- [Decisões Técnicas](#decisões-técnicas)

---

## Tecnologias Utilizadas

- **Backend:**
  - **Go:** Linguagem escolhida pela sua performance e simplicidade.
  - **Gin Gonic:** Framework web para a construção da API RESTful, conhecido pela sua velocidade.
  - **GORM:** ORM para a interação com o banco de dados, facilitando as queries.
  - **JWT (JSON Web Tokens):** Para a implementação de autenticação stateless.
  - **SQLite:** Banco de dados leve e baseado em arquivo, ideal para desenvolvimento e prototipagem.
  - **Swag:** Ferramenta para gerar a documentação da API no formato Swagger.

- **Frontend:**
  - **React:** Biblioteca para a construção da interface de usuário.
  - **React Router DOM:** Para gerenciamento de rotas e navegação na aplicação (Single Page Application).
  - **Axios:** Para realizar as requisições HTTP entre o frontend e a API.
  - **CSS:** em desenvolvimento

- **Estrutura:**
  - **Monorepo:** O projeto foi mantido em um monorepo para gerenciar o código do frontend e do backend.

## Como Configurar e Rodar a Aplicação

### Pré-requisitos
- Go (versão 1.24 ou superior)
- Node.js (versão 18 ou superior)
- Git

### 1. Backend (API em Go)

Abra um terminal e execute os seguintes comandos:

```bash
# Navegue até o diretório da API
cd packages/api-go

# Instale as dependências (só precisa na primeira vez)
go mod tidy

# Execute a aplicação
go run main.go
```

## Acessando a Documentação (Swagger)

Com o backend (API em Go) em execução, você pode acessar a documentação completa da API, gerada pelo Swagger, no seu navegador através do seguinte link:

[**http://localhost:8080/swagger/index.html**](http://localhost:8080/swagger/index.html)

> Lembre-se de regenerar a documentação com o comando `swag init` dentro da pasta `packages/api-go` caso faça alterações nos comentários da API.

## Executando os Testes

O projeto possui testes unitários tanto para o backend quanto para o frontend, conforme solicitado nos requisitos.

### Testes do Backend (API em Go)

Os testes validam os principais endpoints da API. Para executá-los:

```bash
# Navegue até o diretório da API
cd packages/api-go

# Execute todos os testes
go test ./...
```

### Testes do Frontend (React)

O frontend foi criado com `Create React App` e inclui uma estrutura de testes. Para executá-los:

```bash
# Navegue até o diretório do frontend
cd packages/frontend-react

# Execute os testes em modo interativo
npm test
```

## Decisões Técnicas

- **Backend em Go:** A escolha pelo Go e o framework Gin foi motivada pela alta performance e baixo consumo de recursos, características importantes para APIs escaláveis. O uso do GORM com `Preload` otimizou as consultas ao banco, evitando o problema de N+1 queries ao buscar posts e seus autores.

- **Autenticação JWT:** A autenticação stateless com JWT foi implementada para garantir que a API seja escalável e não precise armazenar o estado da sessão. A segurança foi reforçada ao extrair o `username` do token no backend para a criação de posts, prevenindo que um usuário poste em nome de outro.

- **Seeding do Banco de Dados:** Foi criada uma rotina de "seeding" para popular o banco de dados na primeira inicialização. Isso garante que a aplicação sempre tenha dados de exemplo para facilitar testes e demonstrações, melhorando a experiência do desenvolvedor e do avaliador.

- **Frontend em React:** O React foi escolhido por ser uma base sólida e popular para o desenvolvimento de SPAs (Single Page Applications). A estrutura de componentes foi separada em `pages` e `services` para uma melhor organização. O estado de autenticação é gerenciado no componente `App.js` para garantir que a interface reaja corretamente ao login e logout