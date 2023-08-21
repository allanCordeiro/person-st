# rinha de backend

Projeto em Go para participar da [Rinha de backend](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/) 

Foi feito de forma simples, com foco em Consistência (com o trade off de um volume razoável de RPS).

O que foi usado:
 - Go 1.21.0
 - GoChi
 - Postgres
 - nGinX
 - Redis, como cache

 ## Decisões
  - para o create, a API vai direto no banco de dados. Após salvar, gravo o mesmo dado no REDIS, como cache;
  - para o get por ID, a API valida primeiro se encontra o registro no cache, caso contrário realiza uma consulta no banco;
  - a busca por termos é feita também diretamente no banco.

 Não tem escopo de funcionamento em ambiente de produção. 

 ## Para executar

 ```
 docker compose -up
 ```

 Os serviços mencionados [aqui](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/) estarão prontos para executar pela porta 9999.