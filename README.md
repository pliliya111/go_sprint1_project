# Сервис подсчёта арифметических выражений

Веб-сервис на GO, который принимает арифметические выражения через HTTP-запросы и возвращает их результаты.

## Запуск проекта
```
go run ./cmd/main.go
```

## API Endpoint

### POST /api/v1/calculate
#### Примеры запросов
1) Успешный запрос
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
Ответ
```commandline
{
    "result": "6"
}
```
2) Ошибка валидации
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2a"
}'
```
Ответ
```
{
    "error": "Expression is not valid"
}
```
3) Внутренняя ошибка сервера
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2/0"
}'
```
Ответ
```
{
    "error": "Internal server error"
}
```