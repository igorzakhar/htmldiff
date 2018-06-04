# HTML diff service. Реализация на Go.

Cервис для сравнения документов свёрстанных в html, показывающий разницу между двумя файлами. Выводит построчно изменения, сделанные в файле.

# Установка

В программе используются следующие сторонние пакеты:  
[github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)  
[github.com/aryann/difflib](https://github.com/aryann/difflib)  

Для установки пакетов воспользуйтесь следующими командами:
```
$ go get github.com/gin-gonic/gin
```
```
$ go get github.com/aryann/difflib
```


# Использование

Скопируйте данный репозиторий в каталог ```$GOPATH/src/```.

Запуск сервера:
```
$ go run server.go
```
После запуска сервис доступен по адресу [http://127.0.0.1:8080/](http://127.0.0.1:8080)

### Создание исполняемого файла

Перейдите в каталог ```$GOPATH/src/htmldiff``` и выполните команду:
```
$ go build
```

После этого исполняемый файл можно перенести в любой каталог вместе с статическими файлами (каталог ```static/```) и шаблоном (каталог ```templates/```).

Пример запуска исполняемого файла в ОС Linux:
```
$ ./htmldiff
```

# Цели проекта

Проект создан в учебных целях. 