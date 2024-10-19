# GOLANG

function golang_api_directories() {
  # DIRECTORIES

  directories=(
    "cmd/api"
    "infrastructure/database/queries"
    "infrastructure/http/handlers"
    "infrastructure/http/middlewares"
    "infrastructure/http/routes"
    "internal/domain/entities"
    "internal/domain/repositories"
    "internal/domain/usecases"
    "migrations"
    "pkg/config"
    "pkg/database"
    "pkg/utilities/auth"
  )

  for dir in "${directories[@]}"; do
    mkdir -p "$dir"
  done

  # FILES
  files=(
    "cmd/api/main.go"
    "infrastructure/http/router.go"
    "pkg/config/config.go"
    "pkg/config/server.go"
    "pkg/config/database.go"
    "pkg/config/jwt.go"
    "pkg/database/setup.go"
    "pkg/utilities/auth/jwt.go"
    "pkg/utilities/auth/password_encrypt.go"
    ".air.toml"
    ".dockerignore"
    ".env"
    ".env.example"
    ".gitignore"
    "docker-compose.yml"
    "Dockerfile"
    "Makefile"
    "README.md"
    )

  for file in "${files[@]}"; do
    touch "$file"
  done

  echo "CLEAN ARCHITECTURE DIRECTORY STRUCTURE AND FILES CREATED CORRECTLY."
}
