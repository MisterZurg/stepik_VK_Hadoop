# 4.2 Реляционные функции

## 
Дан файл с логами переходов пользователей. Каждая строка состоит из 3 полей: время перехода (unix timestamp), ID пользователя, URL, на который перешел пользователь.

Напишите mapper с помощью Hadoop Streaming, печатающий только те строки из файла, которые соответствуют пользователю user10.

Sample Input:
```
1448713968	user2	https://ru.wikipedia.org/
1448764519	user10	https://stepic.org/
1448713968	user5	http://google.com/
1448773411	user10	https://stepic.org/explore/courses
1448709864	user3	http://vk.com/
```
Sample Output:
```
1448764519	user10	https://stepic.org/
1448773411	user10	https://stepic.org/explore/courses
```


Дан файл с логами переходов пользователей. Каждая строка состоит из 3 полей: время перехода (unix timestamp), ID пользователя, URL, на который перешел пользователь.

Напишите mapper с помощью Hadoop Streaming, печатающий URL из каждой строки.

Sample Input:

1448713968	user2	https://ru.wikipedia.org/
1448764519	user10	https://stepic.org/
1448713968	user5	http://google.com/
1448773411	user10	https://stepic.org/explore/courses
1448709864	user3	http://vk.com/
Sample Output:

https://ru.wikipedia.org/
https://stepic.org/
http://google.com/
https://stepic.org/explore/courses
http://vk.com/


Напишите reducer, который объединяет элементы из множества A и B. На вход в reducer приходят пары key / value, где key - элемент множества, value - маркер множества (A или B)

Sample Input:

1	A
2	A
2	B
3	B
Sample Output:

1
2
3



Напишите reducer, который делает пересечение элементов из множества A и B. На вход в reducer приходят пары key / value, где key - элемент множества, value - маркер множества (A или B)
Sample Input:

1	A
2	A
2	B
3	B
Sample Output:

2



Напишите reducer, который делает вычитание элементов множества B из множества A. На вход в reducer приходят пары key / value, где key - элемент множества, value - маркер множества (A или B)
Sample Input:

1	A
2	A
2	B
3	B
Sample Output:

1





Напишите reducer, который реализует симметричную разность множеств A и B (т.е. оставляет только те элементы, которые есть только в одном из множеств).
На вход в reducer приходят пары key / value, где key - элемент множества, value - маркер множества (A или B)
Sample Input:

1	A
2	A
2	B
3	B
Sample Output:

1
3




Напишите reducer, реализующий объединение двух файлов (Join) по id пользователя. Первый файл содержит 2 поля через табуляцию: id пользователя и запрос в поисковой системе. Второй файл содержит id пользователя и URL, на который перешел пользователь в поисковой системе.

Mapper передает данные в reducer в виде key / value, где key - id пользователя, value состоит из 2 частей: маркер файла-источника (query или url) и запрос или URL.

Каждая строка на выходе reducer должна содержать 3 поля, разделенных табуляцией: id пользователя, запрос, URL.

Sample Input:

user1	query:гугл
user1	url:google.ru
user2	query:стэпик
user2	query:стэпик курсы
user2	url:stepic.org
user2	url:stepic.org/explore/courses
user3	query:вконтакте
Sample Output:

user1	гугл	google.ru
user2	стэпик	stepic.org
user2	стэпик	stepic.org/explore/courses
user2	стэпик курсы	stepic.org
user2	стэпик курсы	stepic.org/explore/courses