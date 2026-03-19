# DevOps_Go

![Docker](https://img.shields.io/badge/Docker-available-2496ED)
![Go](https://img.shields.io/badge/Go-1.24-00ADD8)
![Nginx](https://img.shields.io/badge/Nginx-alpine-009639)
![Docker%20Compose](https://img.shields.io/badge/Docker%20Compose-available-2496ED)

Минимальная инфраструктура: Nginx reverse proxy + Go backend в Docker Compose

## Архитектура

```text
Client
  |
  v
localhost:80
  |
  v
nginx (published)
  |
  v
backend:8080 (internal only)
```

Nginx принимает HTTP-трафик на порту `80` и проксирует запросы во внутренний сервис через Docker service discovery (`backend:8080`). Порт бэкенда не публикуется на хост.

## Структура проекта

```text
.
|-- docker-compose.yml
|-- .env
|-- .env.example
|-- .gitignore
|-- backend
|   |-- Dockerfile
|   |-- .dockerignore
|   |-- go.mod
|   `-- main.go
`-- nginx
    `-- nginx.conf
```

## Сервисы

### backend

- Контекст сборки: `./backend`
- Внутренний порт: `8080`
- Публикация наружу: отсутствует
- Имя контейнера: `${BACKEND_NAME}` из `.env` (по умолчанию `app_backend`)
- Запуск: автоматически через `CMD ["./backend-service"]` в Dockerfile
- Ответ на `/`: `Hello from Effective Mobile!`
- Health endpoint: `/health`

### nginx

- Образ: `nginx:alpine`
- Внутренний порт: `80`
- Публикация наружу: `80:80`
- Имя контейнера: `${NGINX_NAME}` из `.env` (по умолчанию `app_nginx`)
- Источник конфигурации: `./nginx/nginx.conf`, смонтирован в `/etc/nginx/conf.d/default.conf:ro`

## Контейнеры во время выполнения

- `backend` собирается из отдельного multi-stage Dockerfile.
- `nginx` использует официальный образ `nginx:alpine`.
- Оба контейнера подключены к пользовательской bridge-сети `app_network`.

## .env.example

Файл .env.example содержит пример переменных окружения для проекта.

Перед запуском можно создать .env на его основе:

```bash
cp .env.example .env
```

## Запуск

```bash
docker compose up --build
```

## Проверка

Проверить основной ответ:

```bash
curl http://localhost/
```

Ожидаемый вывод:

```text
Hello from Effective Mobile!
```

Проверить заголовки ответа:

```bash
curl -i http://localhost/
```

Для подробной трассировки запроса:

```bash
curl -v http://localhost/
```

Убедиться, что наружу опубликован только Nginx:

```bash
docker compose ps
```

## Health Check

Бэкенд предоставляет endpoint:

```text
GET /health
```

Для сервиса настроен Docker healthcheck, состояние контейнера можно проверить командой:

```bash
docker compose ps
```

В колонке STATUS будет отображаться состояние healthy.

## Остановка

```bash
docker compose down
```

## Примечания по реализации

- Образ бэкенда использует multi-stage build.
- Runtime-образ: `alpine:3.22`.
- Контейнер бэкенда запускается не от root, а от пользователя `newuser`.
- `WORKDIR` явно задан в обоих этапах Dockerfile.
- Nginx проксирует трафик на `backend:8080` по имени сервиса Docker.
- В Nginx настроены proxy-заголовки: `Host`, `X-Real-IP`, `X-Forwarded-For`.
- `X-Forwarded-Proto` в текущей конфигурации не включен.
- `.env` и `.env.example` используются для переменных имен контейнеров в Compose.

## Сеть

- Сеть Compose: `app_network`
- Драйвер: `bridge`
- Внешняя точка входа: `nginx`
- Внутренний сервис без публикации порта: `backend`
- Жестко заданные IP-адреса не используются

## Стек

- Go 1.24
- Nginx (`nginx:alpine`)
- Docker
- Docker Compose
