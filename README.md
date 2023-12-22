# whibs

what happened in bitrix sql

Позволяет следить за текущими запросами к БД анализируя sql log файл битрикса.


Чтобы не запутаться в идущем спаме можно установить фильтры по разным параметрам. 
Позволяет отслеживать html клиент, он находится в файле `client.html`

В битриксе лог включается в файле `bitrix/php_interface/dbconn.php`
```php
$DBDebug = true;
$DBDebugToFile = true;
```
В корне установки должен появится файл `mysql_debug.sql` даже если используется postgres 

Собрать
`go build -ldflags="-s -w" -o ./whibs .`

```txt
Параметры командной строки
-p пусть е файлу mysql_debug.sql
-P порт для сервера
-t автоматически обрезать лог по достижении размера указанного размера, например 2MB 
```

Справка
`./whibs --help`

Пример запуска 
`./whibs -p /home/anton/bitrix24/bitrix/pg24.local/mysql_debug.sql`