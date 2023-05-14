## Фреймворк: 
Gin - это быстрый и легковесный фреймворк для разработки веб-приложений на Golang. Он имеет хорошую производительность и поддерживает маршрутизацию, обработку HTTP-запросов, валидацию данных и другие функции, необходимые для создания API.

## База данных: 
В качестве базы данных я использовал MongoDB, так как она обеспечивает быстрый доступ к данным и горизонтальное масштабирование.
MongoDB - это гибкая и масштабируемая NoSQL база данных, которая хорошо подходит для обработки больших объемов данных.

![Я поднял БД в докер контейнере](/images/1.png)

## Аутентификация: 
Для аутентификации использовал JSON Web Token (JWT), так как он легко реализуется и поддерживает защиту от атак межсервисной подделки запросов (CSRF). JWT (JSON Web Token) - это стандарт авторизации, который позволяет передавать токены между сторонами в формате JSON. Можно использовать библиотеку github.com/dgrijalva/jwt-go для создания и проверки JWT-токенов в Golang.

![JWT](/images/2.png)

## Кэширование
Для кэширования я использовал Redis, так как он предоставляет быстрый доступ к данным и поддерживает хранение данных в памяти и на диске.
Я использовал кэширование для ускорения повторных запросов на чтение информации о студентах, что может уменьшить нагрузку на базу данных и улучшить производительность приложения.

## Мониторинг
Для мониторинга буду использовать Prometheus и Grafana, так как они обеспечивают мощный инструментарий для мониторинга и визуализации метрик.Но в нашем проекте я его не использовал ,но добавил счетчик в main файл.Я хотел реализовать счетчик метрик типа CounterVec, который считает общее количество запросов.
```
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "requests_total",
		Help: "Total number of requests processed by the server, partitioned by status code and HTTP method.",
	}, []string{"code", "method"})

	prometheus.MustRegister(counter)
```

## Тестирование
Я написал юнит-тесты для тестирования функциональности обработчика запросов на аутентификацию пользователя.(Sign-In, Sign-Up)

## Контейнеризация
Для контейнеризации микросервисов я бы использовал Dockerfile и Docker Compose.

## Логирование
Для логирования я использовал Logrus, так как он обеспечивает широкие возможности для настройки формата вывода логов и поддерживает многоуровневую настройку уровней логирования.

## Документация 
Для документирования кода я использовал Swagger.

![](/images/3.png)

## Запуск проекта
- docker compose up
У каждого микросервиса есть свой compose

## Сервисы 
Первый микросервис https://github.com/nurmeden/students-service
Второй микросервис https://github.com/nurmeden/courses-service

## Микросервисы взаимодействуют друг с другом посредством API эндпоинтов для получения информации о студентах и курсах.Реализовал с помощью HTTP GET запроса, отправляется запрос на эндпоинт и ответ на запрос декодируется в модель и возвращается в формате JSON
