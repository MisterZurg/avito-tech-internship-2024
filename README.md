# Сервис баннеров

[![Lint Status](https://img.shields.io/github/actions/workflow/status/MisterZurg/avito-tech-internship-2024/golangci-lint.yml?branch=master&style=for-the-badge)](https://github.com/MisterZurg/avito-tech-internship-2024/actions?workflow=golangci-lint)
[![Coverage Status](https://img.shields.io/codecov/c/gh/MisterZurg/avito-tech-internship-2024.svg?logo=codecov&style=for-the-badge)](https://codecov.io/gh/MisterZurg/avito-tech-internship-2024)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://pkg.go.dev/MisterZurg/avito-tech-internship-2024)


> [!IMPORTANT]
> В Авито есть большое количество неоднородного контента, для которого необходимо иметь единую систему управления.  В частности, необходимо показывать разный контент пользователям в зависимости от их принадлежности к какой-либо группе. Данный контент мы будем предоставлять с помощью баннеров.

> **Задача:** Необходимо реализовать сервис, который позволяет показывать пользователям баннеры, в зависимости от требуемой фичи и тега пользователя, а также управлять баннерами и связанными с ними тегами и фичами.

## Контракты можно тестировать в .http

## Используемые зависимости и тулы
- [echo](https://github.com/labstack/echo) high performance, minimalist Go web framework. Task included by default
- [oapi-codege](https://github.com/deepmap/oapi-codegen) Client and Server Code Generator from the oapi.json
- [sqlx](https://github.com/jmoiron/sqlx) extension on go's standard database/sql
- [pgx](https://github.com/jackc/pgx) pure Go driver and toolkit for PostgreSQL
- [goose](https://github.com/pressly/goose) database migration tool
- [env](https://github.com/caarlos0/env) simple and zero-dependencies library to parse environment variables into structs