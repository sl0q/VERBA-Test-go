# Тестовое задание для VERBA-group

Данный проект является тестовым заданием для VERBA-group.

В проекте реализованы REST API методы для управления задачами из TO-DO листа. Список доступных методов:
- POST /tasks - Создание задачи
- GET /tasks - Запрос списка всех задач
- GET /tasks/:id - Запрос задачи по ID
- PUT /tasks/:id - Обновление задачи по ID
- DELETE /tasks/:id - Удаление задачи по ID

Для хранения данных был использован PostgreSQL

## Зависимости
- Gorm
- go-gorm/postgres
- Fiber

## Установка

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get github.com/gofiber/fiber/v2
```
