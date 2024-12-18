# Calcserv_Go
A web service for calculating arithmetic expressions

Для запуска программы введите в терминал команду: go run ./cmd/main.go

Эта программа вычесляет результаты арифметических выражений, которые пользователь отправляет в виде POST-запроса на url (по умолчанию: http://localhost:8080) с телом:

{
    "expression": "выражение, которое ввёл пользователь"
}

В выражении можно использовать такие арифметические операции как: сложение ("+"), вычитание ("-"), умножение ("*") и деление ("/").
Можно использовать открывающие и закрывающие круглые скобки ("(", ")").
В выражении можно ставить пробелы между цифрами и знаками, можно не ставить.

В случае успешного вычесления результата пользователь получит HTTP-ответ с телом:

{
    "result":"результат выражения"
}
и кодом 200
![image](https://github.com/user-attachments/assets/7a3ff61d-d235-465d-91dd-5651e01c5233)

В случае деления на ноль пользователь получит HTTP-ответ с телом:

{
    "error":"Division by zero"
}
и кодом 422
![image](https://github.com/user-attachments/assets/8d5f39ab-da70-4f59-928b-49ff9c81ba3e)

В случае ввода недопустимых символов (неразрешенные символы, неразрешенные операции, лишние скобки) пользователь получит HTTP-ответ с телом:

{
    "error":"Expression is not valid"
}
и кодом 422

![image](https://github.com/user-attachments/assets/06a22a5f-a889-4497-9769-68f6470f3f49)
![image](https://github.com/user-attachments/assets/022f0f61-c3b9-41e8-ad83-3f82055930e2)
![image](https://github.com/user-attachments/assets/fa961daa-87b2-4f47-b95c-80002cb755f3)

В случае других ошибок пользователь получит HTTP-ответ с телом:

{
    "error":"Internal server error"
}
и кодом 500

Также при некорректном тесте запроса пользователь увидит на экране ошибку: "invalid request body".
![image](https://github.com/user-attachments/assets/6c5c5531-f316-4cba-90aa-35e9ef72c426)
