# Cardápio FURG Bot

Cardápio FURG Bot é um bot desenvolvido em GO durante a aula de Linguagens de Programação no 4o periodo do curso de Sistemas de Informação no C3 da FURG.

Esse bot tem como objetivo facilitar a visualização dos cardápios dos restaurantes universitários da FURG, fazendo um scraping da página no website oficial e postando em uma conta no twitter.
O bot foi desenvolvido utilizando a linguagem GO junto a biblioteca Seleniou para o scraping das informações e a API oficial do twitter para a postagem.

Perfil oficial do bot: https://twitter.com/cardapio_furg

## Fazer setup do projeto

Para fazer setup do projeto primeiro faça `git clone` do repositório e instale as dependências GO.

```sh
go install
```

## Preencher as variaveis de ambiente

Para configurar as chaves APIs necessárias, você deve copiar o arquivo `.env.example` e coloca-lo com o nome de `.env`. Após isso você deve preencher com as credencias de desenvolvedor do Twitter.

## Executar & Fazer postagem

Para executar o bot e fazer a postagem na página do twitter, você deve executar o seguinte comando:

```sh
go run main.go parser.go twitter.go
``` 

## Licença

Cardápio FURG Bot é um software de código aberto licenciado sob a licença [MIT license](LICENSE).
