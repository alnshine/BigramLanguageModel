# Bigram Language Model

### Общие заметки 
* Весь проект написан на Golang
* Есть только файл с именами и `main.go`(на компе должен быть установлен golang, для запуска нужно прописать `go run .`)

#### Логика модели:
* Программа читает `name.txt`, создает мапу с количеством биграмм и их вероятностями 
* Рандомно выбирает первую букву и начинает строить имя, учитывая вероятности биграмм 
* Программа спросит хотите ли вы увидеть вероятности биграмм, если же нет - следующий пункт
* Создается отдельный файл `output.txt`, где так же будут указаны биграммы и их вероятности в % 


#### Запуск:
* После запуска проограмма выводит:
  ```
    gh
    griseoryaeza
    gnaryl
    gaya
    gh
    graia
    gttelemy
    Вы хотите увидеть вероятности всех биграм?
    Y-Да N-Нет
    Я на всякий случай все запишу в отдельный output.txt
