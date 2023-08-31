# rinha de backend

Projeto em Go para participar da [Rinha de backend](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/) 

Foi feito de forma simples, com foco em Consistência (com o trade off de um volume razoável de RPS).

O que foi usado:
 - Go 1.21.0
 - GoChi
 - Postgres
 - nGinX
 - RabbitMQ para salvamento assíncrono dos recursos
 - Redis, como cache

 ## Decisões
  - para o create, a API consulta a existencia do apelido no cache, e caso não exista o dado é salvo em uma fila no RabbitMQ. Após salvar, gravo o mesmo dado no REDIS, como cache;
  - Um worker dentro da própria API consome a fila do RabbitMQ e realiza a inserção no banco de dados usando comando COPY para uso de escrita em lote;
  - para o get por ID, a API valida primeiro se encontra o registro no cache, caso contrário realiza uma consulta no banco;
  - a busca por termos é feita diretamente no banco.

 Não tem escopo de funcionamento em ambiente de produção. 

 ## Para executar

 ```
 docker compose -up
 ```

 ### Para testar
 
 O script usado para stress test na rinha também está no repositório no diretório `stress-test`. Após os serviços subirem executar:

 ```
 ./run-test.sh
 ```

 Os serviços mencionados [aqui](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/) estarão prontos para executar pela porta 9999.