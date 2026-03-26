
# 🔗 Encurtador de Links em Go (Backend)

Este projeto é uma API REST de encurtamento de links desenvolvida em Go, projetada para ser resiliente e de alta performance, utilizando uma camada de cache com Redis e persistência de dados com PostgreSQL.

Toda a infraestrutura é orquestrada via Docker Compose, garantindo um ambiente isolado e de fácil replicação.
# 🚀 Tecnologias Utilizadas

Linguagem: Go (Golang) 🐹

Banco de Dados: PostgreSQL (Persistência) 🐘

Cache: Redis (Performance de leitura) ⚡

Infraestrutura: Docker & Docker Compose 🐳

Ambiente de Dev: Neovim + LazyVim (Linux/Pop!_OS) 💻

# 🏗️ Arquitetura e Diferenciais

O projeto foi construído pensando em conceitos de Engenharia de Software e resiliência de sistemas:

Estratégia de Cache: Antes de consultar o banco de dados, a API verifica o Redis. Isso reduz a latência e a carga no PostgreSQL.

Containerização Profissional: Utilização de Healthchecks no Docker Compose para garantir que a API só inicie após o banco de dados estar totalmente pronto para conexões.

Migrations Automáticas: O sistema verifica e cria as tabelas necessárias (IF NOT EXISTS) durante o startup, facilitando o deploy.

Configuração via Ambiente: Uso de variáveis de ambiente para facilitar a troca entre contextos de desenvolvimento e produção.

# 🛠️ Como Executar

Certifique-se de ter o Docker e o Docker Compose instalados.

  Clone o repositório:
    Bash

    git clone https://github.com/alvarolucio2007/encurtador-links-go.git
    cd encurtador-links-go

  Suba a infraestrutura completa:
    Bash

    docker compose up --build

A API estará disponível em http://localhost:8080.

#📡 Endpoints Principais

Método	Endpoint	Descrição
| Método | Endpoint | Descrição |
| :---: | :---: | :---: |
|POST|/shorten|	Envia um JSON {"url": "..."} e recebe o link encurtado.|
|GET|/:id|Redireciona para a URL original (com cache no Redis).
