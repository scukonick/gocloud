# GoCloud
## Бесплатные облачные вычисления для Go
### Никогда такого не было!
Все вычисления выполняются совершенно бесплатно в гугловском облаке!

Все бесплатно, не берутся деньги ни за кол-во запросов, ни за вычислительные ресурсы!

#### Пример:
```bash
go get github.com/scukonick/gocloud
echo 'package main

import "fmt"

func main() {
        fmt.Println("I was computed in the cloud")
}
' | ./gocloud

I was computed in the cloud
```

### Варианты использования:
#### Бинарник
Отправляешь бинарнику в **stdin** свой код, в **stdout** получаешь результат, в **stderr** - ошибки

#### Библиотека:
```go
package main

import (
	"github.com/scukonick/gocloud/lib"
)
...
...
comp := lib.NewComputer()
result, err := comp.Run(code) // code - это твой код на Golang
```


