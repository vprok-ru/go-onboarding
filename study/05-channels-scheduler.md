## Теория
### Каналы и планировщик
- https://github.com/quii/learn-go-with-tests/blob/master/concurrency.md
- https://github.com/quii/learn-go-with-tests/tree/master/concurrency
- https://docs.google.com/document/d/1-n2O-c3C6kz-4vNKJdtR9bjh-kXqJe7RydSOTBxvB3g/edit?usp=sharing
- https://habr.com/ru/post/308070/
- https://github.com/quii/learn-go-with-tests/tree/master/select
- https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
- https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/chan_test.go
- https://blog.gopheracademy.com/advent-2019/directional-channels/ 
  
### Исходники
- https://github.com/golang/go/blob/master/src/runtime/proc.go

Если будет интересно, что за зверь такой sudog в гошном рантайме: https://groups.google.com/forum/#!topic/golang-codereviews/rC9BLPFvkW8

go 1.14 https://golang.org/doc/go1.14 

### Планировщик и sync pool
- https://docs.google.com/presentation/d/1UCOEHw-0oSUiOA94aTX79Ki7VbsmPtlYlv1EfyI1ylE/edit#slide=id.p
- https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html
- https://habr.com/ru/post/333654/
- https://habr.com/ru/post/277137/
- https://dev-gang.ru/article/go-ponjat-dizain-syncpool-cpvecztx8e/
- https://golang.org/src/sync/po
## Практика
Реализовать паттерн цепочка вызовов https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern в соответствии с конвенцией
