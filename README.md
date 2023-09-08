# terraformCloudAPI
Codigo para interagir com o terraform Cloud por API

codigo inicial - cria um workspace no terrafomCloud, precisa passar como variavel de ambiente:
**TF_WORKSPACE_NAME** = nome que deseja criar de workspace
**TF_ORGANIZATION_NAME** = nome da organizacao 
**TF_API_TOKEN** = o token para acesso
**TF_PROJECT_NAME** = nome do projeto que deseja associar ao workspace

## Observacao

O workspace é criado para o tipo api-driven workflow 