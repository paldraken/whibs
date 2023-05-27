# whibs

what happened in bx sql


Позволяет следить за текущими запросами к БД анализируя sql log файл битрикса.
Чтобы не запутаться в идущем спаме можно установить фильтры по разным параметрам. 

В битриксе лог включается в файле `bitrix/php_interface/dbconn.php`
```php
$DBDebug = true;
$DBDebugToFile = true;
```
Собрать
`go build -ldflags="-s -w" -o ./whibs .`

Справка
`./whibs --help`