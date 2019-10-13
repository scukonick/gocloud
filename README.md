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

### Восторженная реакция публики и критиков!!!
![image](https://user-images.githubusercontent.com/1472592/66718888-a86b0d00-edf1-11e9-84dd-05c4d9cb0790.png)

![image](https://user-images.githubusercontent.com/1472592/66718913-e36d4080-edf1-11e9-80c2-500c1bab2408.png)

![image](https://user-images.githubusercontent.com/1472592/66718924-013aa580-edf2-11e9-8082-dba8be13aab7.png)

![image](https://user-images.githubusercontent.com/1472592/66718935-19aac000-edf2-11e9-884f-e726fac7f03c.png)

![image](https://user-images.githubusercontent.com/1472592/66718948-48c13180-edf2-11e9-9ea2-5777f1f8e736.png)

