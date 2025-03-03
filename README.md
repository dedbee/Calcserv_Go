# Calcserv_Go - сервис подсчета арифметических выражений.

Этот проект представляет собой асинхронный сервер для вычисления арифметических выражений. Сервер состоит из двух основных компонентов:  
  -Оркестратор — принимает арифметические выражения от пользователей, разбивает их на задачи и управляет их выполнением.  
  -Агент — выполняет задачи, полученные от оркестратора, и возвращает результаты.  
Проект написан на языке Go и использует HTTP для взаимодействия между компонентами.  

Структура проекта  
CalcServ_Go/  
├── cmd/
│   └── main.go                # Точка входа в приложение (оркестратор)  
├── internal/  
│   ├── agent/                 # Логика агента (вычислителя)  
│   │   └── agent.go
│   ├── application/           # Логика приложения (оркестратора)  
│   │   └── application.go  
│   └── orchestrator/          # Логика оркестратора (управление задачами)  
│       └── orchestrator.go  
├── pkg/  
│   ├── calculation/           # Логика вычислений (арифметические операции)  
│   │   ├── calculation.go  
│   │   └── errors.go  
├── go.mod                     # Файл модуля Go (зависимости)  
└── README.md                  # Документация проекта  

Инструкция по запуску  
Клонируйте репозиторий:  
`git clone https://github.com/dedbee/Calcserv_Go.git`  
Запустите проект:  
`go run cmd/main.go`  
