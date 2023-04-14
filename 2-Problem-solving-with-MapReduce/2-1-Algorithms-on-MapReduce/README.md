# 4.1 Алгоритмы на MapReduce

## In-mapper combining v.1
Напишите программу, которая реализует **In-mapper combining v.1** для задачи **WordCount** в **Hadoop Streaming**.

**Sample Input:**
```
aut Caesar aut nihil
aut aut
de mortuis aut bene aut nihil
```
**Sample Output:**
```
nihil	1
aut	2
Caesar	1
aut	2
nihil	1
aut	2
de	1
bene	1
mortuis	1
```

## In-mapper combining v.2
Напишите программу, которая реализует **In-mapper combining v.2** для задачи **WordCount** в **Hadoop Streaming**.

**Sample Input:**
```
aut Caesar aut nihil
aut aut
de mortuis aut bene aut nihil
```
**Sample Output:**
```
aut	6
mortuis	1
bene	1
Caesar	1
de	1
nihil	2
```

## Reducer Implementation
Реализуйте **reducer** в задаче подсчета среднего времени, проведенного пользователем на странице.

**Mapper** передает в **reducer** данные в виде **key / value**, где **key** - адрес страницы, **value** - число секунд, проведенных пользователем на данной странице.

Среднее время на выходе приведите к целому числу.

Sample Input:
```
www.facebook.com	100
www.google.com	10
www.google.com	5
www.google.com	15
www.stepic.org	60
www.stepic.org	100
```
Sample Output:
```
www.facebook.com	100
www.google.com	10
www.stepic.org	80
```

## Combiner Implementation
Реализуйте **Combiner** в задаче подсчета среднего времени, проведенного пользователем на странице.

Mapper пишет данные в виде **key / value**, где **key** - адрес страницы, **value** - пара чисел, разделенных ";". Первое - число секунд, проведенных пользователем на данной странице, второе равно 1.

**Sample Input:**
```
www.facebook.com	100;1
www.google.com	10;1
www.google.com	5;1
www.google.com	15;1
stepic.org	60;1
stepic.org	100;1
```
**Sample Output:**
```
www.facebook.com	100;1
www.google.com	30;3
stepic.org	160;2
```