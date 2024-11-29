# social-network-algorithm
Sistema de sugestão de conexões em uma rede social utilizando algoritmos de busca em largura (BFS) diretamente integrado a um banco de dados.

## Estrutura Do Projeto
O projeto está estruturado em três camadas principais:

- Handler: Responsável por receber requisições HTTP, processá-las e enviar respostas.

- Service: Implementa a lógica de negócio, incluindo o algoritmo BFS e ordenação das sugestões.

- Repository: Contém funções para interação com o banco de dados, como consultas e manipulações de dados.

## Banco de Dados

O modelo utiliza um banco de dados SQL para armazenar:

- Usuários (tabela users)

- Conexões entre usuários (tabela user_connections)

## Algoritmo de Sugestão

Busca em Largura (BFS):

- O algoritmo parte do usuário atual e explora conexões em níveis.

- Evita sugerir conexões diretas ou o próprio usuário.

- Calcula a "distância" de cada usuário a partir do inicial.

## Ordenação:

Sugestões são ordenadas com base em:

- Menor caminho (menor distância no grafo).

- Maior número de conexões comuns.

- Fluxo de Execução

- O handler recebe a requisição HTTP do cliente (usuário solicitando sugestões).

- A camada de service executa o algoritmo BFS para buscar as conexões mais relevantes.

- A camada de repository consulta as conexões diretamente do banco de dados.

- As sugestões são ordenadas e retornadas ao cliente.


