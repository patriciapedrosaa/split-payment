
# Split Payment

A aplicação **Split Payment** consiste em uma aplicação para calcular e dividir um pagamento entre uma dada quantidade de pessoas. Ela recebe uma lista de compras e uma lista de e-mails, calcula o total e divide a conta igualmente para cada e-mail.

## Sobre

O **Split Payment** recebe duas entradas: uma lista de compras e uma lista de e-mails no formato JSON. 

A lista de compras contém um conjunto de produtos compostos por:
-  **item**: corresponde à descrição do produto;
- **quantity**: corresponde a quantidade do produto, podendo ser (unidade, peso, pacote);
- **price**: preço referente à unidade/peso/pacote do item;

Exemplo lista de compras:
```json
[
    {
        "item":"Biscoito",
        "quantity": 1,
        "price": 2.80
    },
    {
        "item":"Carne",
        "quantity": 1.5,
        "price": 60.0
    }
]
```

A lista de e-mails contém o cunjunto de e-mails:
Exemplo lista de e-mails:
```json
{
    "emails": [
        "pessoa1@email.com",
        "pessoa2@email.com",
        "pessoa3@email.com",
        "pessoa4@email.com",
        "pessoa5@email.com",
        "pessoa6@email.com",
        "pessoa7@email.com"

    ]
}
```


Por se tratar de divisão de valores, as listas de e-mails e de compras não devem estar vazias como também não devem conter quantidades ou preços iguais a zero ou negativos.

O sistema calcula o total desde que os preços estejam correspondentes às unidades de cada produto.  

Por exemplo, se a lista contém 1 pacote de biscoito, o preço deve ser referente ao pacote e não à uma unidade de biscoito. 
|Item|  Quant.|Preço
|--|--|--|
|Biscoito| 1| 2,80

A aplicação também possui testes unitários que validam o comportamento das funções implementadas.


## Tecnologias

- [Golang](https://golang.org/) *v1.16*
- [Testify](github.com/stretchr/testify)  *v1.7.0*

## Setup
Execute o comando na raíz do projeto: 
 `$ go run main.go`
 
Na pasta "resources" existem dois arquivos *JSON* que podem ser atualizados para testes locais:
 - /resources/emails.json: lista de e-mails
 - /resources/shoppingList.json: lista de compras

**Teste unitários:**
Execute o comando na raíz do projeto:
`$ go test -v ./service`