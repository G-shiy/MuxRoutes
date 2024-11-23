# MuxRoutes

## Visão Geral
Este projeto utiliza PostgreSQL como banco de dados, configurado com o Docker. A aplicação é desenvolvida com Go e usa o roteador mux.

Siga as instruções abaixo para configurar e executar o ambiente localmente.

---

## Pré-requisitos
1. **Docker e Docker Compose** instalados.
   - [Guia de instalação do Docker](https://docs.docker.com/get-docker/)
   - [Guia de instalação do Docker Compose](https://docs.docker.com/compose/install/)
2. **Go** instalado na sua máquina.
   - [Guia de instalação do Go](https://go.dev/doc/install)

---

## Configuração

### 1. Criar o arquivo `.env`
Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis de ambiente:

```env
DB_USER=seu_usuario        # Substitua por seu usuário
DB_PASSWORD=sua_senha      # Substitua por sua senha
DB_NAME=nome_do_banco      # Substitua pelo nome do banco
``` 

### 2. Iniciar o Banco de Dados
Use o comando abaixo para iniciar o banco de dados PostgreSQL com Docker:

```bash
docker-compose up -d
```
- Isso iniciará o banco de dados em segundo plano.
- O Docker Compose usará o arquivo docker-compose.yml fornecido no projeto para configurar o PostgreSQL.

### 3. Testar a Conexão ao Banco
```bash
docker exec -it nome_do_container psql -U seu_usuario -d nome_do_banco
```
*Substitua nome_do_container, seu_usuario e nome_do_banco pelos valores usados no .env.*

