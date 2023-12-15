[Versão em Português](./README_PT.md)

# Nightscout Brazilian Food DB

Add to your Nightscout the content from the latest Brazilian Diabetes Society carbs counting manual.

## Requirements

* A nightscout deployment with the `food` plugin installed.
* A working installation of Go (https://go.dev/)

## Running

To add the food to your nightscout database run:

```bash
NS_BASE_URL="https://your-nightscout-host.com" \
NS_API_SECRET="0123456789yourAPItoken" \
go run main.go
```

This will make requests to your nightscout api sending the data from `data.csv`. Once the database is
populated, there's no need to use this system anymore.

Check out the [food plugin](https://nightscout.github.io/nightscout/setup_variables/#food-custom-foods) to get more instructions on how to proceed from here.

## Créditos

* Manual de Contagem de Carboidratos para pessoas com diabetes is written and distributed by the Brazilian Diabetes Society https://diabetes.org.br/e-book/manual-de-contagem-de-carboidratos/