# Neurox

**Neurox** — веб-платформа для генерации изображений с помощью искусственного интеллекта.

Пользователь описывает желаемое изображение обычным текстом, после чего Neurox обрабатывает запрос и создаёт изображение с использованием AI.

Проект построен на Go и использует PostgreSQL для хранения данных, Redis для кэширования и вспомогательных операций, RabbitMQ для асинхронной обработки задач, S3-совместимое хранилище для изображений и Gemini API для работы с искусственным интеллектом.

---

## Возможности

* Регистрация и авторизация пользователей
* JWT-аутентификация
* Генерация изображений по текстовому описанию
* Асинхронная обработка задач через RabbitMQ
* Хранение изображений в S3-совместимом хранилище
* PostgreSQL для хранения данных приложения
* Redis для кэширования и временных данных
* Отправка электронной почты через SMTP
* Интеграция с Gemini API
* Валидация входных данных
* Swagger/OpenAPI документация
* Структурированное логирование
* Graceful shutdown HTTP-сервера

---

## Технологический стек

### Backend

* Go 1.25+
* JWT
* PostgreSQL
* Redis
* RabbitMQ
* S3 / MinIO
* Gemini API
* SMTP

### Основные библиотеки

| Библиотека                    | Назначение      |
| ----------------------------- | --------------- |
| `aws-sdk-go-v2`               | Работа с AWS/S3 |
| `pgx/v5`                      | PostgreSQL      |
| `go-redis/v9`                 | Redis           |
| `amqp091-go`                  | RabbitMQ        |
| `golang-jwt/jwt/v5`           | JWT             |
| `go-playground/validator/v10` | Валидация       |
| `google.golang.org/genai`     | Gemini API      |
| `go.uber.org/zap`             | Логирование     |
| `swaggo/swag`                 | Swagger/OpenAPI |
| `go-mail`                     | Отправка email  |
| `google/uuid`                 | UUID            |

---

## Архитектура

Neurox использует асинхронную архитектуру для обработки запросов на генерацию изображений.

Общий сценарий:

```text
                    ┌──────────────┐
                    │   Browser    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │  HTTP API    │
                    │    Go        │
                    └──────┬───────┘
                           │
             ┌─────────────┼─────────────┐
             │             │             │
             ▼             ▼             ▼
       ┌──────────┐   ┌──────────┐   ┌──────────┐
       │PostgreSQL│   │  Redis   │   │RabbitMQ  │
       └──────────┘   └──────────┘   └────┬─────┘
                                          │
                                          ▼
                                   ┌─────────────┐
                                   │ AI Worker   │
                                   │   Gemini    │
                                   └──────┬──────┘
                                          │
                                          ▼
                                   ┌─────────────┐
                                   │  S3/MinIO   │
                                   │   Storage   │
                                   └─────────────┘
```

### Генерация изображения

1. Пользователь вводит prompt.
2. Frontend сохраняет prompt и отправляет пользователя на страницу создания изображения.
3. Backend проверяет JWT.
4. Создаётся запрос на генерацию.
5. Задача передаётся в RabbitMQ.
6. Worker получает задачу.
7. Worker обращается к AI API.
8. Полученный результат сохраняется в S3.
9. Информация о результате сохраняется в PostgreSQL.
10. Клиент получает информацию о готовом изображении.

---

## Структура проекта

Примерная структура проекта:

```text
neurox/
├── cmd/
│   └── ...
│
├── internal/
│   ├── config/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── middleware/
│   ├── worker/
│   └── ...
│
├── migrations/
│   └── ...
│
├── docs/
│   └── ...
│
├── public/
│   └── index.html
│
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

Фактическая структура каталогов может отличаться в зависимости от текущей реализации проекта.

---

# Требования

Для локального запуска понадобятся:

* Go 1.25+
* PostgreSQL
* Redis
* RabbitMQ
* S3-совместимое хранилище (например, MinIO)
* Gemini API key
* SMTP-сервер — если используется функциональность отправки email

---

# Конфигурация

Все настройки приложения задаются через environment variables.

Создайте `.env` на основе `.env.example`:

```bash
cp .env.example .env
```

## Переменные окружения

### HTTP

```env
HTTP_ADDR=:5050
HTTP_SHUTDOWN_TIMEOUT=30s
```

| Переменная              | Описание                  |
| ----------------------- | ------------------------- |
| `HTTP_ADDR`             | Адрес и порт HTTP-сервера |
| `HTTP_SHUTDOWN_TIMEOUT` | Таймаут graceful shutdown |

---

### PostgreSQL

```env
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
POSTGRES_TIMEOUT=10s
```

| Переменная          | Описание                |
| ------------------- | ----------------------- |
| `POSTGRES_USER`     | Пользователь PostgreSQL |
| `POSTGRES_PASSWORD` | Пароль PostgreSQL       |
| `POSTGRES_DB`       | Имя базы данных         |
| `POSTGRES_TIMEOUT`  | Таймаут подключения     |

Если приложение использует отдельный host/port PostgreSQL, их необходимо добавить в конфигурацию согласно реализации проекта.

---

### Logging

```env
LOGGER_LEVEL=DEBUG
LOGGER_FOLDER=./
```

| Переменная      | Описание             |
| --------------- | -------------------- |
| `LOGGER_LEVEL`  | Уровень логирования  |
| `LOGGER_FOLDER` | Директория для логов |

Доступные уровни зависят от реализации логгера, например:

```text
DEBUG
INFO
WARN
ERROR
```

---

### JWT

```env
JWT_SECRET=
JWT_TTL=15m
```

| Переменная   | Описание                 |
| ------------ | ------------------------ |
| `JWT_SECRET` | Секрет для подписи JWT   |
| `JWT_TTL`    | Время жизни access token |

Для production необходимо использовать длинный случайно сгенерированный секрет.

---

### S3 / MinIO

```env
S3_ENDPOINT=http://localhost:9000
S3_PUBLIC_URL=http://localhost:9000
S3_REGION=us-east-1
S3_ACCESS_KEY=
S3_SECRET_KEY=
S3_BUCKET=neurox
```

| Переменная      | Описание             |
| --------------- | -------------------- |
| `S3_ENDPOINT`   | Endpoint S3/MinIO    |
| `S3_PUBLIC_URL` | Публичный URL файлов |
| `S3_REGION`     | Регион S3            |
| `S3_ACCESS_KEY` | Access Key           |
| `S3_SECRET_KEY` | Secret Key           |
| `S3_BUCKET`     | Bucket для файлов    |

Для локальной разработки удобно использовать MinIO.

---

### RabbitMQ

```env
PUBLISHER_USER=
PUBLISHER_PASSWORD=
PUBLISHER_HOST=localhost
PUBLISHER_PORT=5672
```

| Переменная           | Описание              |
| -------------------- | --------------------- |
| `PUBLISHER_USER`     | Пользователь RabbitMQ |
| `PUBLISHER_PASSWORD` | Пароль RabbitMQ       |
| `PUBLISHER_HOST`     | Host RabbitMQ         |
| `PUBLISHER_PORT`     | AMQP-порт             |

RabbitMQ используется для передачи задач генерации между API и worker.

---

### SMTP

```env
SMTP_HOST=
SMTP_PORT=
SMTP_USER=
SMTP_PASSWORD=
SMTP_FROM=
```

| Переменная      | Описание          |
| --------------- | ----------------- |
| `SMTP_HOST`     | SMTP-сервер       |
| `SMTP_PORT`     | SMTP-порт         |
| `SMTP_USER`     | Пользователь SMTP |
| `SMTP_PASSWORD` | Пароль SMTP       |
| `SMTP_FROM`     | Адрес отправителя |

---

### Redis

```env
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_TIMEOUT=10s
```

Redis используется приложением для кэширования и других быстрых операций, требующих in-memory storage.

---

### Gemini

```env
GEMINI_API_KEY=
GEMINI_TIMEOUT=10s
```

| Переменная       | Описание                  |
| ---------------- | ------------------------- |
| `GEMINI_API_KEY` | API key для Gemini        |
| `GEMINI_TIMEOUT` | Таймаут запросов к Gemini |

API key нельзя добавлять непосредственно в исходный код или коммитить в Git.

---

# Пример `.env`

Минимальный пример локальной конфигурации:

```env
HTTP_ADDR=:5050
HTTP_SHUTDOWN_TIMEOUT=30s

POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=neurox
POSTGRES_TIMEOUT=10s

LOGGER_LEVEL=DEBUG
LOGGER_FOLDER=./

JWT_SECRET=change-me-in-development
JWT_TTL=15m

S3_ENDPOINT=http://localhost:9000
S3_PUBLIC_URL=http://localhost:9000
S3_REGION=us-east-1
S3_ACCESS_KEY=minioadmin
S3_SECRET_KEY=minioadmin
S3_BUCKET=neurox

PUBLISHER_USER=guest
PUBLISHER_PASSWORD=guest
PUBLISHER_HOST=localhost
PUBLISHER_PORT=5672

SMTP_HOST=
SMTP_PORT=
SMTP_USER=
SMTP_PASSWORD=
SMTP_FROM=

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_TIMEOUT=10s

GEMINI_API_KEY=
GEMINI_TIMEOUT=10s
```

> Значения выше предназначены только для локальной разработки. Не используйте стандартные credentials в production.

---

# Установка

Клонируйте репозиторий:

```bash
git clone <repository-url>
cd neurox
```

Установите зависимости:

```bash
go mod download
```

Проверьте проект:

```bash
go test ./...
```

Соберите приложение:

```bash
go build ./...
```

---

# Локальная инфраструктура

Для полноценной работы Neurox необходимо запустить инфраструктурные сервисы:

```text
PostgreSQL
Redis
RabbitMQ
MinIO / S3
```

После запуска сервисов необходимо:

1. Создать базу данных Neurox.
2. Настроить необходимые таблицы/миграции.
3. Создать S3 bucket `neurox`.
4. Настроить credentials.
5. Заполнить `.env`.
6. Запустить backend.
7. Запустить worker, если он является отдельным процессом.

---

# Запуск

После настройки `.env` приложение можно запустить:

```bash
go run .
```

или, если entrypoint находится в `cmd`:

```bash
go run ./cmd/...
```

Для production-сборки:

```bash
go build -o neurox .
```

Запуск:

```bash
./neurox
```

По умолчанию HTTP-сервер запускается на:

```text
http://localhost:5050
```

---

# Frontend

Главная страница Neurox содержит:

* презентационный блок;
* описание возможностей;
* инструкцию по использованию;
* примеры prompt;
* форму генерации изображения;
* регистрацию и авторизацию;
* переход в профиль;
* переход к созданию изображения;
* logout.

Frontend взаимодействует с JWT, сохранённым в `localStorage`.

Используемый ключ:

```text
neurox_token
```

Email пользователя хранится в:

```text
neurox_email
```

Prompt для передачи на страницу создания изображения временно сохраняется в:

```text
neurox_prompt
```

через `sessionStorage`.

---

# Аутентификация

После успешной авторизации frontend сохраняет JWT:

```javascript
localStorage.setItem("neurox_token", token);
```

При наличии токена интерфейс переключается с гостевого режима на авторизованный.

### Гость

Пользователю доступны:

* Войти
* Регистрация
* Начать бесплатно

### Авторизованный пользователь

Пользователю доступны:

* Профиль
* Создать фото
* Выйти

При logout JWT удаляется:

```javascript
localStorage.removeItem("neurox_token");
localStorage.removeItem("neurox_email");
```

После этого страница перезагружается.

---

# Генерация изображений

На главной странице пользователь может ввести prompt длиной до 1000 символов.

Пример:

```text
Современный дом в горах, панорамные окна,
деревянная терраса, вечерний свет,
туман между горами
```

После нажатия **«Сгенерировать»**:

1. Проверяется наличие prompt.
2. Prompt сохраняется в `sessionStorage`.
3. Показывается состояние загрузки.
4. Авторизованный пользователь перенаправляется на:

```text
/requests/create
```

5. Неавторизованный пользователь перенаправляется на:

```text
/login
```

---

# API Documentation

Проект использует Swagger/OpenAPI через `swaggo`.

После запуска приложения документация может быть доступна по адресу:

```text
/swagger/index.html
```

если соответствующий Swagger route подключён в HTTP router.

Для генерации документации используется:

```bash
swag init
```

---

# Безопасность

Необходимо соблюдать следующие правила:

* не коммитить `.env`;
* не хранить `JWT_SECRET` в Git;
* не хранить `GEMINI_API_KEY` в исходном коде;
* не публиковать AWS/S3 credentials;
* использовать HTTPS в production;
* использовать отдельные credentials для production;
* ограничить права S3 bucket;
* использовать сильный JWT secret;
* валидировать пользовательский ввод;
* устанавливать разумные timeout для внешних сервисов.

Пример `.gitignore`:

```gitignore
.env
.env.*
!.env.example

*.log

bin/
tmp/

.idea/
.vscode/

.DS_Store
```

---

# Production

Перед deployment необходимо заменить development-конфигурацию:

```env
LOGGER_LEVEL=DEBUG
```

на подходящий production уровень логирования.

Также необходимо:

* использовать production PostgreSQL;
* использовать защищённый Redis;
* настроить RabbitMQ credentials;
* использовать production S3;
* включить HTTPS;
* использовать секреты через Secret Manager / environment;
* ограничить доступ к инфраструктурным сервисам;
* настроить мониторинг;
* настроить резервное копирование PostgreSQL;
* настроить lifecycle/storage policies для изображений.

---

# Graceful Shutdown

Приложение поддерживает graceful shutdown.

Параметр:

```env
HTTP_SHUTDOWN_TIMEOUT=30s
```

определяет максимальное время, которое серверу предоставляется для корректного завершения активных операций.

---

# Лицензия

Лицензия проекта пока не указана.

Если проект будет распространяться публично, рекомендуется добавить отдельный файл `LICENSE` с выбранной лицензией.

---

# Статус проекта

Проект находится в разработке.

Основные направления:

* [ ] Завершение API
* [ ] Генерация изображений
* [ ] Worker для обработки очереди
* [ ] Полная интеграция S3
* [ ] Полная интеграция Gemini
* [ ] История генераций
* [ ] Улучшение профиля пользователя
* [ ] Расширение Swagger-документации
* [ ] Тестовое покрытие
* [ ] Docker Compose для локального запуска
* [ ] Production deployment

---

## Автор

**Neurox**

Платформа для генерации изображений с использованием искусственного интеллекта.
