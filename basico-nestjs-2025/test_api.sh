#!/usr/bin/env bash
# test_api.sh — laboratório didático de cURL para a API NestJS + Swagger (books)
# Requisitos: bash, curl. (Opcional: jq para "pretty print" do JSON)

set -euo pipefail

# =========================
# Configurações
# =========================
BASE_URL="${BASE_URL:-http://localhost:3000/v1}"
AUTH_TOKEN="${AUTH_TOKEN:-fake-jwt}"     # Bearer "falso" só para demo
USE_COLOR="${USE_COLOR:-1}"              # 1 = colorir output, 0 = sem cores
PAUSE_SECS="${PAUSE_SECS:-0}"            # segundos de pausa entre passos (0 = sem pausa)

# =========================
# Aparência / Utilidades
# =========================
if [[ "${USE_COLOR}" == "1" ]]; then
  BOLD="\033[1m"; DIM="\033[2m"; OK="\033[32m"; INFO="\033[36m"; WARN="\033[33m"; ERR="\033[31m"; RESET="\033[0m"
else
  BOLD=""; DIM=""; OK=""; INFO=""; WARN=""; ERR=""; RESET=""
fi

have_jq=0
if command -v jq >/dev/null 2>&1; then have_jq=1; fi

pause() { sleep "${PAUSE_SECS}"; }

title() {
  echo -e "\n${BOLD}==> $*${RESET}"
}

step() {
  echo -e "${INFO}→ $*${RESET}"
}

ok() {
  echo -e "${OK}✔ $*${RESET}"
}

warn() {
  echo -e "${WARN}⚠ $*${RESET}"
}

fail() {
  echo -e "${ERR}✖ $*${RESET}"
  exit 1
}

pp_json() {
  if [[ $have_jq -eq 1 ]]; then jq .; else cat; fi
}

curl_json() {
  # Uso: curl_json METHOD PATH [DATA_JSON]
  local method="$1"
  local path="$2"
  local data="${3:-}"

  if [[ -n "${data}" ]]; then
    response=$(curl -sS -X "$method" \
      -H "Content-Type: application/json" \
      -H "Authorization: Bearer ${AUTH_TOKEN}" \
      "${BASE_URL}${path}" \
      -d "${data}")
  else
    response=$(curl -sS -X "$method" \
      -H "Authorization: Bearer ${AUTH_TOKEN}" \
      "${BASE_URL}${path}")
  fi

  echo "${response}"
}

require_up() {
  title "Verificando se a API está no ar (${BASE_URL})"
  if curl -sS "${BASE_URL}/health" >/dev/null 2>&1; then
    ok "API respondeu em /health"
  else
    warn "Não encontrei /health. Tentando /books (a API pode não ter health controller)..."
    if curl -sS "${BASE_URL}/books" >/dev/null 2>&1; then
      ok "API respondeu em /books"
    else
      fail "A API não respondeu. Confirme que está rodando: npm run start:dev (porta 3000 por padrão)."
    fi
  fi
}

# =========================
# Execução do laboratório
# =========================

title "Início do laboratório com cURL"
step "Base URL: ${BASE_URL}"
step "Auth Token (didático): ${AUTH_TOKEN}"
[[ $have_jq -eq 1 ]] && ok "jq detectado — respostas serão 'pretty printed'." || warn "jq não encontrado — instalá-lo deixa as respostas mais legíveis."

require_up
pause

# 1) Criar 3 livros
title "1) Criar 3 livros (POST /books)"

BOOK1_JSON='{"title":"Clean Code","author":"Robert C. Martin","year":2008,"description":"Práticas de código limpo"}'
BOOK2_JSON='{"title":"Refactoring","author":"Martin Fowler","year":1999}'
BOOK3_JSON='{"title":"The Pragmatic Programmer","author":"Andrew Hunt & David Thomas","year":1999}'

resp1=$(curl_json POST "/books" "${BOOK1_JSON}")
echo "${resp1}" | pp_json
ID1=$(echo "${resp1}" | (jq -r '.id' 2>/dev/null || true))
[[ -n "${ID1}" && "${ID1}" != "null" ]] && ok "Livro 1 criado com id=${ID1}" || fail "Não consegui obter id do Livro 1"

pause

resp2=$(curl_json POST "/books" "${BOOK2_JSON}")
echo "${resp2}" | pp_json
ID2=$(echo "${resp2}" | (jq -r '.id' 2>/dev/null || true))
[[ -n "${ID2}" && "${ID2}" != "null" ]] && ok "Livro 2 criado com id=${ID2}" || fail "Não consegui obter id do Livro 2"

pause

resp3=$(curl_json POST "/books" "${BOOK3_JSON}")
echo "${resp3}" | pp_json
ID3=$(echo "${resp3}" | (jq -r '.id' 2>/dev/null || true))
[[ -n "${ID3}" && "${ID3}" != "null" ]] && ok "Livro 3 criado com id=${ID3}" || fail "Não consegui obter id do Livro 3"

pause

# 2) Listar todos
title "2) Listar todos os livros (GET /books)"
curl -sS -X GET "${BASE_URL}/books" | pp_json
pause

# 3) Buscar o segundo livro pelo id
title "3) Buscar o segundo livro (GET /books/${ID2})"
curl -sS -X GET "${BASE_URL}/books/${ID2}" | pp_json
pause

# 4) Atualizar o segundo livro (PATCH)
title "4) Atualizar descrição do segundo livro (PATCH /books/${ID2})"
UPDATE_JSON='{"description":"Edição comentada com exemplos modernos"}'
curl -sS -X PATCH \
  -H "Content-Type: application/json" \
  "${BASE_URL}/books/${ID2}" \
  -d "${UPDATE_JSON}" | pp_json
pause

# 5) Remover o primeiro livro
title "5) Remover o primeiro livro (DELETE /books/${ID1})"
curl -sS -X DELETE "${BASE_URL}/books/${ID1}" | pp_json
pause

# 6) Listar final
title "6) Conferir resultado final (GET /books)"
curl -sS -X GET "${BASE_URL}/books" | pp_json
pause

# 7) Demonstração de erros
title "7) Erros didáticos"
step "7.1) Buscar livro inexistente (GET /books/9999)"
STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X GET "${BASE_URL}/books/9999" || true)
echo -e "${DIM}HTTP ${STATUS}${RESET}"
curl -sS -X GET "${BASE_URL}/books/9999" | pp_json || true
pause

step "7.2) Tentar criar livro inválido (sem 'title') — deve dar 400"
INVALID_JSON='{"author":"Autor X","year":2020}'
STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${AUTH_TOKEN}" \
  -d "${INVALID_JSON}" \
  "${BASE_URL}/books" || true)
echo -e "${DIM}HTTP ${STATUS}${RESET}"
curl -sS -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${AUTH_TOKEN}" \
  -d "${INVALID_JSON}" \
  "${BASE_URL}/books" | pp_json || true

ok "Laboratório finalizado! Veja também a documentação em: http://localhost:3000/docs"
