## Go Template

O sistema deve receber um CEP, identificar a cidade e retornar o clima atual (temperatura em graus celsius, fahrenheit e kelvin).

#### 🖥️ Detalhes Gerais:

Especificações e detalhes gerais do projeto. 
- O sistema deve receber um CEP válido de 8 digitos
- Será utilizada a API viaCEP para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
- Será utilizada a API WeatherAPI para consultar a temperatura da cidade: https://www.weatherapi.com/
- O sistema deve responder adequadamente nos seguintes cenários:
  - Em caso de sucesso:
    - Código HTTP: **200**
    - Response Body: **{ "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }**
  - Em caso de falha, caso o CEP seja inválido (com formato correto):
    - Código HTTP: **422**
    - Mensagem: **invalid zipcode**
  - Em caso de falha, caso o CEP não seja encontrado:
    -  Código HTTP: **404**
    - Mensagem: **can not find zipcode**
- Desenvolver testes automatizados para garantir a qualidade do código
- Utilizar docker-compose para subir a aplicação
- O sistema deverá ser publicado no Google Cloud Run

> 💡 Dica:<br/>
> - A conversão de Celsius para Fahrenheit é: **F = C * 9/5 + 32**
> - A conversão de Celsius para Kelvin é: **K = C + 273.15**

#### 🗂️ Estrutura do Projeto
    .
    ├── cmd                # Entrypoints da aplicação
    │    └── ordersystem   
    │           ├── main.go       ### Entrypoint principal
    │           ├── wire.go       ### Injeção de dependências
    │           └── .env          ### Arquivo de parametrizações globais
    ├── configs            # helpers para configuração da aplicação (viper)
    ├── internal
    │    ├── domain        # Core da aplicação
    │    │      ├── repository    ### Interfaces de repositório
    │    │      └── entity        ### Entidades de domínio
    │    ├── application   # Implementações de casos de uso e utilitários
    │    │      └── usecase       ### Casos de uso da aplicação
    │    └── infra         # Implementações de repositórios e conexões com serviços externos
    │           ├── repository    ### Implementações de repositório
    │           └── web           ### Implementações e códigos gerados para a API Rest
    ├── pkg                # Pacotes reutilizáveis utilizados na aplicação
    ├── .env               # Arquivo de parametrizações globais
    └── README.md

#### 🧭 Parametrização
A aplicação servidor possui um arquivo de configuração `cmd/ordersystem/.env` onde é possível definir os parâmetros de timeout e URL's das API's para busca das informações do endereço.

```
DB_DRIVER=mysql                 # Database driver
```

#### 🚀 Execução:
Para executar a aplicação, basta utilizar o docker-compose disponível na raiz do projeto. Para isso, execute o comando abaixo:
```bash
$ docker-compose up
```

> 💡 O comando acima poderá falhar caso alguma das portas utilizadas estejam em uso. Caso isso ocorra, será necessário alterar as portas no arquivo `.env` ou encerrar os processos que estão utilizando as portas (8000, 8080, 50051, 3306, 5672 e 15672).

### 📝 Usando as API's:

#### 1. REST API:

- **Criar um pedido:**
```bash
$ curl --location 'http://localhost:8000/order' \
--header 'Content-Type: application/json' \
--data '{
    "id": "aff0-2223-8842-fe215",
    "price": 66.5,
    "tax": 1.1
}'
```

- **Listar todos os pedidos (exemplo):**
```bash
$ curl --location 'http://localhost:8000/order'
```

- **Consultar um pedido (exemplo):**
```bash
$ curl --location 'http://localhost:8000/order/<<OrderId>>'
```

#### 2. GraphQL API:

> Para utilizar a API GraphQL, é necessário acessar o playground disponível em `http://localhost:8080/`.

- **Criar um pedido (exemplo):**
```graphql
mutation createOrder {
    createOrder(input:{id: "aff0-2223-8842-fe214",Price:854.1, Tax: 0.8}){
        id
    }
}
```

- **Listar todos os pedidos (exemplo):**
```graphql
query listOrders {
    listOrders {
        id
        Price
        Tax
        FinalPrice
    }
}
```

- **Consultar um pedido (exemplo):**
```graphql
query findOrder {
    listOrder(id:"aff0-2223-8842-fe215"){
        id
        Price
        Tax
        FinalPrice
    }
}
```

#### 3. gRPC API:

> Para a utilização da API gRPC, foi utilizado o Evans gRCP client. Para instalar, siga as instruções disponíveis em: [evans - install](https://github.com/ktr0731/evans?tab=readme-ov-file#installation)


- **Iniciando Evans:**
```bash
$ evans -r repl --host localhost --port 50051
 
localhost:50051>  package pb
pb@localhost:50051>  service OrderService
```

- **Criar um pedido (exemplo):**
```bash
pb.OrderService@localhost:50051> call CreateOrder
id (TYPE_STRING) => 1
price (TYPE_FLOAT) => 100
tax (TYPE_FLOAT) => 50
{
  "finalPrice": 150,
  "id": "1",
  "price": 100
}
```

- **Listar todos os pedidos (exemplo):**
```bash
pb.OrderService@localhost:50051> call ListOrders
{
  "orders": [
    {
      "finalPrice": 150,
      "id": "1",
      "price": 100,
      "tax": 50
    }
  ]
}
```

- **Consultar um pedido (exemplo):**
```bash
pb.OrderService@localhost:50051> call ListOrderById
id (TYPE_STRING) => aff0-2223-8842-fe214
{
  "finalPrice": 854.9,
  "id": "aff0-2223-8842-fe214",
  "price": 854.1,
  "tax": 0.8
}
```