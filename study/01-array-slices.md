### Теория
Знакомство с конвенцией по code style: Конвенция go команды

### Массивы и срезы

- https://go101.org/article/type-system-overview.html типы и алиасы в го 
- https://www.sohamkamani.com/blog/golang/arrays-vs-slices/
- https://programming.guide/go/nil-slice-vs-empty-slice.html
- https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94
- https://www.callicoder.com/golang-slices/
- https://github.com/golang/go/wiki/SliceTricks
- https://go101.org/article/memory-leaking.html 
  
В обязательном порядке изучаем исходники: 
- https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/slice.go
- https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/slice_test.go 
  
### Строки
- https://blog.golang.org/strings
- https://medium.com/rungo/string-data-type-in-go-8af2b639478
- https://medium.com/go-walkthrough/go-walkthrough-bytes-strings-packages-499be9f4b5bd
- https://habr.com/ru/post/307554/
- https://medium.com/go-walkthrough/go-walkthrough-io-package-8ac5e95a9fbd (оригинал io)
- https://habr.com/ru/post/306914/ (перевод io)
- https://github.com/golang/go/blob/master/src/io/io.go (исходники io)
- https://github.com/golang/go/blob/master/src/strings/strings.go (исходники строки)
  
### Практика
Реализовать паттерн фасад https://en.wikipedia.org/wiki/Facade_pattern в соответствии с конвенцией

Для практического задания необходимо создать публичный репозиторий на github'e в контрибьютеры добавить (ssmmxx).

Реализация паттерна должна быть следующая: 
- описать пакет facade в директории internal/facade
- в директории cmd/facade создать файл main.go, который фактически будет представлять собой интеграционное тестирование пакета
- необходимо изолировать каждую сущность, участвующую в паттерне в отдельный пакет. Для удобства предлагается в пакете internal/facade сделать подпапки для каждой сущности.