#!/bin/bash

# Defina os nomes antigo e novo do módulo
OLD_MODULE="boilerPlate"
NEW_MODULE="sixTask"

# Detecta o sistema operacional para usar o sed corretamente
if [[ "$OSTYPE" == "darwin"* ]]; then
  SED_CMD="sed -i ''"
else
  SED_CMD="sed -i"
fi

# Diretório base do projeto
BASE_DIR=$(pwd)

echo "📦 Renomeando módulo Go:"
echo "🔁 De: $OLD_MODULE"
echo "➡️  Para: $NEW_MODULE"
echo ""

# Atualiza go.mod
echo "📝 Atualizando go.mod..."
$SED_CMD "s|$OLD_MODULE|$NEW_MODULE|g" "$BASE_DIR/go.mod"

# Atualiza todos os arquivos .go
echo "🔍 Atualizando arquivos .go..."
find "$BASE_DIR" -type f -name "*.go" -exec $SED_CMD "s|$OLD_MODULE|$NEW_MODULE|g" {} +

# Atualiza arquivos auxiliares (opcional)
echo "📂 Atualizando arquivos auxiliares (.md, .yaml, etc)..."
find "$BASE_DIR" -type f \( -name "*.md" -o -name "*.yaml" -o -name "*.yml" -o -name "*.txt" \) \
  -exec $SED_CMD "s|$OLD_MODULE|$NEW_MODULE|g" {} +

# Atualiza dependências
echo ""
echo "🔧 Executando go mod tidy..."
go mod tidy

echo ""
echo "✅ Módulo renomeado com sucesso!"
