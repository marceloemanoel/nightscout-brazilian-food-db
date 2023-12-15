# Nightscout Food DB Feeder

Adiciona à sua instalação do Nightscout o conteúdo mais recente das tabelas de contagem de carboidrato
da Sociedade Brasileira de Diabetes.

## Requisitos

* Um servidor do Nightscout com o plugin `food` instalado.
* Instalação funcional de Go (https://go.dev/)

## Execução

Para adicionar os alimentos ao banco de dados do Nightscout execute o seguinte comando:

```bash
NS_BASE_URL="https://your-nightscout-host.com" NS_API_SECRET="0123456789yourAPItoken" go run main.go
```

Com isso, o programa fará requisições http ao servidor do nightscout enviando os dados contidos no arquivo `data.csv`.
Uma vez com o banco de dados populado, não há necessidade de executar o programa novamente.

Verifique a documentação do (plugin `food`)[https://nightscout.github.io/nightscout/setup_variables/#food-custom-foods] 
para maiores informações em como prosseguir.
