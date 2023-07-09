# 4.1 Алгоритмы на MapReduce

## In-mapper combining v.1
Напишите программу, которая реализует **In-mapper combining v.1** для задачи **WordCount** в **Hadoop Streaming**.

### Псевдокод
```go
// In-mapper combining v.1 для задачи WordCount
class Mapper
    method Map (docid id, doc d)
        H = new AssociativeArray
        for all term t in doc d do
            H{t} = H{t} + 1
        for all term t in H do
            Emit(term t, count H{t}) 
```
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

### Псевдокод

```go
// In-mapper combining v.2 для задачи WordCount
class Mapper
    method Initialize
        H = new AssociativeArray

    method Map(docid id, doc d)
        for all termt in doc d do 
            H{t} = H{t} + 1
    
    method Close 
        for all term t in H do
            Emit(term t, count H{t})
```
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

### Псевдокод
```go
// Среднее значение v1
class Mapper
    method Map(string t, integer r)
        Emit(string t, integer r)


class Reducer
    method Reduce(string t, integers [r1, r2, ...])
        sum = 0
        cnt = 0
        for all integers r in [r1, r2, ...] do
            sum = sum + r
            cnt = cnt + 1
        avg = sum / cnt
        Emit(string t, integer avg)
```
**Sample Input:**
```
www.facebook.com	100
www.google.com	10
www.google.com	5
www.google.com	15
www.stepic.org	60
www.stepic.org	100
```
**Sample Output:**
```
www.facebook.com	100
www.google.com	10
www.stepic.org	80
```

## Combiner Implementation
Реализуйте **Combiner** в задаче подсчета среднего времени, проведенного пользователем на странице.

Mapper пишет данные в виде **key / value**, где **key** - адрес страницы, **value** - пара чисел, разделенных ";". Первое - число секунд, проведенных пользователем на данной странице, второе равно 1.

### Псевдокод
```go
// Среднее значение v2
class Mapper 
    method Map (string t, integer r) 
        Emit(string t, integer r)


class Combiner 
    method Combine(string t, integers [rl, r2, ...]
        sum = cnt = 0

        for all integers r in [rl, r2, ..] do
            sum = sum + r
            cnt = cnt + 1
        
        Emit(string t, pair(sum, cnt))


class Reducer
        method Reduce(string t, pairs[(sl, c1), (s2, c2), ...])
        sum = cnt = 0

        for all pairs p in [(sl, c1), (s2, c2) ...] do
            sum = sum + p.s
            cnt = cnt + p.c 
        
        avg = sum / cnt
        Emit(string t, integer avg)
```

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


## Distinct Values v1 Mapper ph1
Реализуйте mapper из фазы 1 задачи Distinct Values v1.

Mapper принимает на вход строку, содержащую значение и через табуляцию список групп, разделенных запятой.

### Псевдокод
```go
// Фаза 1 - Distinct Values v1. 
class Mapper
    method Map(null, record [value f, categories [g1, g2,...]])
        for all category g in [gl, g2,...]
            Emit(record [g, f], count 1)
```

**Sample Input:**
```
1	a,b
2	a,d,e
1	b
3	a,b
```
**Sample Output:**
```
1,a	1
1,b	1
2,a	1
2,d	1
2,e	1
1,b	1
3,a	1
3,b	1
```


## Distinct Values v1 Reducer ph1
Реализуйте reducer из фазы 1 задачи Distinct Values v1.

Reducer принимает на вход данные, созданные mapper из предыдущей шага.

### Псевдокод
```go
// Фаза 1 - Distinct Values v1. 
class Reducer 
    method Reduce(record [g, f], counts [n1, n2, ...])
        Emit(record [g,f], null) 
```

**Sample Input:**
```
1,a	1
1,b	1
1,b	1
2,a	1
2,d	1
2,e	1
3,a	1
3,b	1
```
**Sample Output:**
```
1,a
1,b
2,a
2,d
2,e
3,a
3,b
```

## Distinct Values v1 Mapper ph2
Реализуйте mapper из фазы 2 задачи Distinct Values v1.

Mapper принимает на вход строку, содержащую значение и группу, разделенные запятой.

### Псевдокод
```go
// Фаза 2 - Distinct Values v1.
class Mapper
    method Map(record [f, g], null) 
        Emit(value g, count 1)


class Reducer
    method Reduce(value g, counts [n1, n2, ...])
        Emit(value g, sum([n1, n2, ...]) 
```

**Sample Input:**
```
1,a
2,a
3,a
1,b
3,b
2,d
2,e
```
**Sample Output:**
```
a	1
a	1
a	1
b	1
b	1
d	1
e	1
```

## Distinct Values v2 Reducer
Реализуйте reducer из задачи Distinct Values v2.

Reducer принимает на вход строки, каждая из которых состоит из разделенных табуляцией значения и названия группы.

### Псевдокод
```go
// Distinct Values v2.
class Mapper 
    method Map(null, record [value f, categories [gl, g2, ...])
        for all category g in [gl, g2, ...]
            Emit(value f, category g)


class Reducer
    method Initialize
        H = new AssociativeArray : category -> count

    method Reduce(value f, categories [gl, g2, ...])
        [gl', g2', ...] = ExcludeDuplicates( [gl, g2, ...])
        for all category g in [gl', g2', ...]
            H{g) = H(g) + 1

    method Close
        for all category g in H do
            Emit(category g, count H{g)) 
```

**Sample Input:**
```
1	a
1	b
1	b
2	a
2	d
2	e
3	a
3	b
```
**Sample Output:**
```
a	3
d	1
b	2
e	1
```

## Cross-Correlation : Pairs
Реализуйте mapper для задачи Cross-Correlation, который для каждого кортежа создает все пары элементов, входящих в него.

Mapper принимает на вход кортежи - строки, состоящие из объектов, разделенных пробелом.

Mapper пишет данные в виде key / value, где key - пара объектов, разделенных запятой, value - единица.

### Псевдокод
```go
// Cross-Corelation : Pairs
class Mapper 
    method Map(null, items [it, i2, ...] )
        for all item i in [i1, i2, ...]
            for all item j in [i1, i2, ...]
                Emit(pair [i j], count 1) 


class Reducer
    method Reduce(pair [i j], counts [c1, c2, ...])
        s = sum([c1, c2, ...])
        Emit(pair[i j], count s) 
```


**Sample Input:**
```
a b
a b a c
```

**Sample Output:**
```
a,b	1
b,a	1
a,b	1
a,c	1
b,a	1
b,a	1
b,c	1
a,b	1
a,c	1
c,a	1
c,b	1
c,a	1
```

## Cross-Correlation : Stripe

Реализуйте mapper для задачи Cross-Correlation, который для каждого объекта из кортежа создает stripe.

Mapper принимает на вход кортежи - строки, состоящие из объектов, разделенных пробелом.

Mapper пишет данные в виде key / value, где key - объект, value - соответствующий stripe (пример: a:1,b:2,c:3)

### Псевдокод
```go
// Cross-Corelation : Stripes
class Mapper
    method Map(null, items [i1, i2, ...])
        for all item i in [i1, i2, ...]
            H = new AssociativeArray : item -> counter
            for all item j in [i1, i2, ...]
                H{j} = H{j} + 1
            Emit(item i, stripe H) 


class Reducer
    method Reduce(item i, stripes [H1, H2, ...])
        H = new AssociativeArray : item -> counter
        H = merge-sum([HI, H2, ...])
        for all item j in H.keys()
            Emit(pair [i j], H{j}) 
```

**Sample Input:**
```
a b
a b a c
```
**Sample Output:**
```
a	b:1
b	a:1
a	b:1,c:1
b	a:2,c:1
a	b:1,c:1
c	b:1,a:2
```