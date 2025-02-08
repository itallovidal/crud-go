# crud-go
## Sobre o desafio

Neste projeto, será desenvolvido uma API RESTful que irá realizar operações CRUD. Ele começará com um mapa representando o banco de dados, e ao longo dos estudos será implementado um DB real utilizando PostgreSQL. 

O objeto é fixar, aprender e entender conceitos gerais de GO, como nomenclatura, organização de projeto, funcionamento do pacote HTTP, uso de bibliotecas mais simples, integração com banco de dados e aplicabilidade de regras de clean code. 

## Instruções

Inicie um projeto (módulo) Go e implemente os seguintes endpoints:

### Endpoints

| Método | URL               | Descrição |
|--------|-------------------|-----------|
| **POST**   | `/api/users`     | Cria um usuário usando as informações enviadas no corpo da requisição. |
| **GET**    | `/api/users`     | Retorna um array de usuários. |
| **GET**    | `/api/users/:id` | Retorna o objeto do usuário com o `id` especificado. |
| **DELETE** | `/api/users/:id` | Remove o usuário com o `id` especificado e retorna o usuário deletado. |
| **PUT**    | `/api/users/:id` | Atualiza o usuário com o `id` especificado usando dados do corpo da requisição. Retorna o usuário modificado. |

## API Requests
Abaixo estão exemplos de requisições para interagir com a API.

### Cadastra um ou vários usuários
```sh
curl --request POST \
  --url http://localhost:8080/api/users \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/10.3.0' \
  --data '[
  {
    "first_name": "Alice Johnson",
    "last_name": "Johnson",
    "biography": "Apaixonada por aventuras e livros."
  },
  {
    "first_name": "Carlos Mendes",
    "last_name": "Mendes",
    "biography": "Gosta de tecnologia e música clássica."
  },
  {
    "first_name": "Beatriz Souza",
    "last_name": "Souza",
    "biography": "Exploradora urbana e fotógrafa amadora."
  },
  {
    "first_name": "Daniel Oliveira",
    "last_name": "Oliveira",
    "biography": "Entusiasta de esportes radicais e café."
  },
  {
    "first_name": "Fernanda Costa",
    "last_name": "Costa",
    "biography": "Cinéfila e fã de ficção científica."
  }
]
'
```
### Devolve todos os usuários
```sh
curl --request GET \
  --url http://localhost:8080/api/users \
  --header 'User-Agent: insomnia/10.3.0'
```
### Retorna usuário com ID específico
```sh
curl --request GET \
  --url http://localhost:8080/api/users/1703c5e5-d6e0-467f-840e-7b4b32ab3650 \
  --header 'User-Agent: insomnia/10.3.0'
```
### Deleta usuário com ID específico
```sh
curl --request DELETE \
  --url http://localhost:8080/api/users/1703c5e5-d6e0-467f-840e-7b4b32ab3650 \
  --header 'User-Agent: insomnia/10.3.0'
```
### Atualiza e retorna um usuário específico
```sh
curl --request PUT \
  --url http://localhost:8080/api/users/1703c5e5-d6e0-467f-840e-7b4b32ab3650 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/10.3.0' \
  --data '  {
    "first_name": "Vitoria",
    "last_name": "Vidal",
    "biography": "Mediciner muito chata"
  }'
```