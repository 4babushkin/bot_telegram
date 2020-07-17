# Телеграм бот для системы АСКУЭ

## Для запуска Телеграм бота 

переименовываем `.env_example`

```
cp .env_example .env
```
Изменяем переменные 
```env
MSSQL_ADDRESS=192.168.0.1
MSSQL_USER=user
MSSQL_PASS=pass
BOT_TOKEN=123456789:******_*-********-gIM2pHhb4k
```

Запускаем
```
docker-compose up -d
```


### Так собираем когда изменим код

docker build -t 4babushkin/askuebot .

### Так пушим в реджестри
docker push 4babushkin/askuebot:latest

