# s21adapter 🐺☝️

Адаптер/прокси для внутреннего GQL API платформы edu.21-school.ru, предоставляет спецификацию Swagger 2.

На данный момент адаптер не имеет механизма авторизации и выполняет все действия от имени одного пользователя, данные для входа которого необходимо передать в переменных окружения.

Пример настройки окружения:

```sh
S21_USERNAME=example@student.21-school.ru
S21_PASSWORD=example
```

## Использование контейнера

Готовые сборки адаптера публикуются в ghcr, см. [Packages](https://github.com/s21toolkit/s21adapter/pkgs/container/s21adapter).

## Готовая спецификация

Для каждой версии адаптера публикуется релиз с готовой спецификацией Swagger 2, см. [Releases](https://github.com/s21toolkit/s21adapter/releases).

## Сборка и запуск вручную

Сборка:

```sh
make # Сборка адаптера
make spec # Сборка генератора спецификации
```

Запуск адаптера:

```sh
./s21adapter
```

Генерация спецификации:

```sh
./s21adapter_spec > swagger.json
```

## Генерация методов

Методы адаптера генерируются на основе логов запросов платформы либо списка методов, поддерживаемых клиентом.

Пример генерации из логов с помощью [s21auto](https://github.com/s21toolkit/s21auto):

```sh
s21auto adapter generate --har log.har -o s21adapter/internal/controller
```

> Если какие-то методы не нужны, из папки internal/controller можно убрать всё кроме [`controller.go`](/internal/controller/controller.go)
