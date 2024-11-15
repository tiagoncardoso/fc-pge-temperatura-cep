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

### 📝 Usando a API:

- **Buscar temperatura baseada no CEP informado:**
```bash
$ curl --location 'http://localhost:8000/temperature/{zipCode}' \
```
#### Exemplo de resposta de sucesso (status code 200):
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.6
}
```

#### Exemplo de resposta de falha - CEP inválido (status code 422):
```json
{
  "error": "invalid zipcode"
}
```

#### Exemplo de resposta de falha - CEP não encontrado (status code 404):
```json
{
  "error": "can not find zipcode"
}
```