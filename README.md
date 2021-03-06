# day_256_http_get
## Микросервис для поиска 256 дня в году
Для запуска используйте команду  
```
gradle startService
```
или
```
go run main.go  
```
### Описание работы сервиса
Программа использует стандартную библиотеку "net/http" для прослушивания запросов на 80м порте. При каждом запросе в отдельном потоке (горутине) вызывается обработчик, вычисляющий дату и отправляющий ответ пользователю.  
Год мы получаем из GET-запроса в виде строки. Если этот аргумент не удаётся интерпретировать как число, ответ не возвращается.  
Для вычисления 256-го дня в году мы сначала создаём структуру Time, указывающую на начало этого года. 
Затем к этому Time прибавляется Duration (временной промежуток) равный 256 дням или 22118400 секундам. 
Для этого была использована стандартная библиотека "time".  
Полученное время преобразуется в строку формата DD/MM/YY.  
Для формирования ответа клиенту заполняются 2 поля в стркутуре. Одно всегда равно 200, а в другое мы записываем полученную строку с датой.  
После этого стркутура преобразуется в JSON и уже в виде текста отправляется клиенту. Для этого была использована стандартная библиотека "encoding/json".
### Примеры запросов
```
> curl http://localhost?year=2017
{"errorCode":200,"dataMessage":"13/09/17"}
> curl http://localhost?year=2018
{"errorCode":200,"dataMessage":"13/09/18"}
> curl http://localhost?year=2019
{"errorCode":200,"dataMessage":"13/09/19"}
> curl http://localhost?year=2020
{"errorCode":200,"dataMessage":"12/09/20"}
> curl http://localhost?year=2021
{"errorCode":200,"dataMessage":"13/09/21"}
> curl http://localhost?year=1940
{"errorCode":200,"dataMessage":"12/09/40"}
> curl http://localhost?year=12
{"errorCode":200,"dataMessage":"12/09/12"}
> curl http://localhost?year=0
{"errorCode":200,"dataMessage":"12/09/00"}
> curl http://localhost?year=-201
{"errorCode":200,"dataMessage":"13/09/01"}
> curl http://localhost?year=-204
{"errorCode":200,"dataMessage":"12/09/04"}
```
### Лицензия
This software is distributed under MIT license  
зис софваре ис дистрибутедет андер MIT лайценз  
