## Запуск
### 1. Скопируйте .env
```sh
cp .env.example .env
```

### 2. Запустите приложение с помощью docker-compose:
При запуске нужно определить тип хранилища "POSTGRES" или IN_MEMORY
```sh
STORE_TYPE="POSTGRES" docker compose up -d
```

## Пример использования REST API сервера

### Создание короткой ссылки
Запрос:
```sh
curl -X POST localhost:8000/links \
    --data '{
    "full_url": "https://google.com/"
    }'
```
Ответ:
```json
{
  "short_url":"https://shorty.io/aaaacJ1AyE"
}
```

### Получение полной ссылки по короткой
Запрос:
```sh
curl -X GET localhost:8000/links \
    --data '{
      "short_url":"https://shorty.io/aaaacJ1AyE"
    }'
```
Ответ:
```json
{
    "full_url": "https://google.com/"
}
```


### Создание короткой ссылки
Запрос:
```sh
curl -X POST localhost:8000/links \
    --data '{
    "full_url": "https://google.com/"
    }'
```
Ответ:
```json
{
  "short_url":"https://shorty.io/aaaacJ1AyE"
}
```

## Пример использования GRPC сервера
`!Запросы нужно выполнять в директории с приложением!`
### Создание короткой ссылки
Запрос:
```sh
grpcurl -plaintext \
        -import-path api \
        -proto grpc/v1/shorten.proto \
        -d '{"full_url":"https://examp.com"}' \
        localhost:8001 shorten.v1.ShortenService/Create
```
Ответ:
```json
{
  "shortUrl": "https://shorty.io/aaaacJ1AyG"
}
```

### Получение полной ссылки по короткой
Запрос:
```sh
grpcurl -plaintext \
        -import-path api \
        -proto grpc/v1/shorten.proto \
        -d '{"short_url":"https://shorty.io/aaaacJ1AyB"}' \
        localhost:8001 shorten.v1.ShortenService/Get
```
Ответ:
```json
{
  "fullUrl": "https://examp.com"
}
```
