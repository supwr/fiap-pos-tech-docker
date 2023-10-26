## Programming languages list

### Proposta

> Você precisará criar uma aplicação web simples (linguagens e framework à sua escolha), que apenas acesse um banco de dados e retorne uma lista de linguagens de programação. Essas linguagens podem ser pré-carregadas no banco de dados.
> O banco de dados escolhido deve ser executado também em container e a conexão entre aplicação e o container deve acontecer utilizando Docker network. 
> A aplicação deverá ser executada em um container Docker, uma imagem deverá ser gerada utilizando o Dockerfile e publicada no Docker Hub, de forma pública.
> Com a imagem publicada no Docker Hub, deverá ser criado um arquivo Docker Compose para orquestrar toda a configuração dos serviços, redes e variáveis de ambiente necessárias.

### Entregáveis

> Link público do repositório com aplicação e seu Dockerfile;
> No README desse repositório deverá conter o passo a passo para execução da aplicação utilizando container, isto é, qualquer configuração ou variáveis necessárias deverão estar descritas;
> O arquivo docker-compose.yml com a orquestração necessária para executar a aplicação utilizando a imagem pública postada no docker hub.

### Executando aplicação

```
docker-compose up
```

### Consumindo API

```
curl --location 'localhost:8000'
```

### Docker Hub image
```
https://hub.docker.com/repository/docker/supwr/fiap-programming-languages/general
```