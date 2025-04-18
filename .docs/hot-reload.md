# Como Usar o Hot Reload com Air

## Visão Geral

O Go Starter Kit inclui suporte para hot reload usando a ferramenta [Air](https://github.com/cosmtrek/air). O hot reload permite que você desenvolva mais rapidamente, pois a aplicação é automaticamente reiniciada quando os arquivos são modificados, eliminando a necessidade de reiniciar manualmente o servidor a cada alteração.

## Instalação

O Air já está configurado como dependência do projeto, mas você também pode instalá-lo globalmente para facilitar o uso:

```bash
# Usando o comando do Makefile
make install-air

# Ou diretamente via Go
go install github.com/cosmtrek/air@latest
```

## Configuração

O projeto já inclui um arquivo de configuração `.air.toml` na raiz, que define como o Air deve monitorar e reiniciar a aplicação. As principais configurações incluem:

- Diretórios e arquivos a serem monitorados
- Comandos para compilar e executar a aplicação
- Extensões de arquivos a serem observadas
- Configurações de log e cores para a saída

## Como Usar

Para iniciar a aplicação com hot reload:

```bash
# Usando o comando do Makefile
make dev

# Ou diretamente via Air
air
```

A aplicação será iniciada e o Air começará a monitorar os arquivos do projeto. Quando você fizer alterações em qualquer arquivo Go, o Air automaticamente recompilará e reiniciará a aplicação.

## Personalização

Se você precisar personalizar o comportamento do Air, pode editar o arquivo `.air.toml` na raiz do projeto. Consulte a [documentação oficial do Air](https://github.com/cosmtrek/air) para mais informações sobre as opções de configuração disponíveis.

## Solução de Problemas

Se o hot reload não estiver funcionando corretamente:

1. Verifique se o diretório `tmp` existe na raiz do projeto
2. Certifique-se de que o Air está instalado corretamente
3. Verifique se o arquivo `.air.toml` está configurado corretamente
4. Verifique os logs do Air para identificar possíveis erros

## Observações

- O hot reload é recomendado apenas para ambiente de desenvolvimento
- Em produção, use o comando `make run` ou execute diretamente o binário compilado
