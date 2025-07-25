## Normalização ##

### Bancos de Dados Relacionais ###
Pratica na Aula

#### Cenário ####
Após uma análise realizada em um sistema de faturamento, foi detectado um design ruim no banco de dados. Ele tem uma tabela de faturas que armazena diferentes tipos de dados.

Como você pode ver, a tabela tem dados que poderiam ser normalizados e separados em diferentes entidades.       
![table](/godatabase/dia02/tarde/exercicios-pratica/image/table.png)

#### Exercício ####
Solicitado para o cenário acima:
1. Aplicar regras de normalização e desenvolver um modelo DER que atinja a terceira forma normal (3FN).
![der](/godatabase/dia02/tarde/exercicios-pratica/image/der.png)

2. Descreva com suas próprias palavras cada etapa da decomposição e da aplicação das regras para visualizar a abordagem adotada:

- Normalização 1FN  
Para realizar a normalização 1FN separei algumas informações  clientes, artigos e faturas, criando tabelas independentes.

- Normalização 2FN      
Na etapa 2FN notei que cliente e artigo não dependiam da PK da fatura, então, transferi essa informação para uma nova tabela.

- Normalização 3FN   
Na etapa 3FN procurei me certificar de que cada tabela contém somente informações que dependem unicamente de sua PK, evitando redundância e garantindo a integridade e flexibilidade do banco.