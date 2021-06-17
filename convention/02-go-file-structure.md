- Используем файлы для создания специфических сущностей
- Имена интерфейсов пишем в соответствии с конвенцией https://golang.org/doc/effective_go.html#interface-names
- receiver метода сокращается до одного символа https://gobyexample.com/methods
- Все ПУБЛИЧНЫЕ интерфейсы, структуры, методы публичных структур, фабрики, константы, переменные должны иметь комментарий https://www.digitalocean.com/community/tutorials/how-to-write-comments-in-go
- Пустая строка, как разделитель, в функциях и методах допускается, если у нас есть логическое разделение блоков. Возможно, в этом слуае, напрашивается еще и комментарий.
- receiver методов делаем только по указателю https://dave.cheney.net/tag/receivers
- Скрываем реализацию пакетов (приватные структуры)
- Используем специфические интерфейсы для описания поведения пакетов
- Для простых пакетов используем одинаковое название для публичного интерфейса и приватной структуры
- Для порождения скрытой структуры используем функцию-фабрику (NewИмяСтруктуры)
- Публичные константы и переменные (вне структуры или функции) не используем нигде, кроме main.go и тестов
- Используем только приватные константы и переменные (если есть необходимость, например, повысить читаемость кода или переиспользование памяти)
- Глобальные (внеконтекстные) структуры не используем
- Взаимодействие между пакетами исключительно через приватные интерфейсы
- Не порождаем зависимости от другого кода - весь необходимый инструментарий и знания структура должна получать при инициализации или в рантайме

Структура файла
- импорты
- константы
- переменные
- приватные интерфейсы
- публичный интерфейс
- приватная структура
- публичные методы структуры
- приватные методы структуры
- фабрика

Каждый файл должен заканчиваться одной пустой строкой
Goland: File | Settings | Editor -> Ensure line feed at file end on Save

Пример:
```go

// One file = one struct
package some_package

import (
    "standart/golang/libraries"
    "external/golang/libraries"
    "internal/golang/libraries"
)

const (
    weUseOnlyPrivateConstantsIfNecessary
    commonUseOfPrivateConstantsIsForTestsAndMainGoOnly
    WeCanNOTUsePublicConstants
)

var (
    sameAsConstants
)

type somePrivate interface {
    Do() (err error)
}

// SomePublic interface describes some public behavior
type SomePublic interface {
    Do() (err error)
}

type privateImplementationOfSomePublicInterface struct {
    someVar somePrivate
}

func (p *privateImplementationOfSomePublicInterface) Do() (err error) {
    return
}

func (p *privateImplementationOfSomePublicInterface) some() {
    // todo some private
}

// NewPrivateImplementationOfSomePublicInterface creates private implementation of SomePublic interface
func NewPrivateImplementationOfSomePublicInterface(someVar somePrivate) SomePublic {
    return &privateImplementationOfSomePublicInterface{
        someVar: someVar,
    }
}

// <<<---- В конце файла обязательно пустая строка
```
