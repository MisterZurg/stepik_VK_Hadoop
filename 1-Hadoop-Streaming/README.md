# 3.5 Hadoop Streaming

## Mapper
Напишите программу, которая реализует **mapper** для задачи **WordCount** в **Hadoop Streaming.**

**Sample Input:**
```
Vivere est cogitare
Vivere militate est
Scientia potentia est
```
**Sample Output:**
```
Vivere	1
est	1
cogitare	1
Vivere	1
militate	1
est	1
Scientia	1
potentia	1
est	1
```
[Решение](1-Hadoop-Streaming/mapper-wordcount.go)

## Reducer
Напишите программу, которая реализует reducer для задачи WordCount в Hadoop Streaming.

**Sample Input:**
```
cogitare	1
est	1
est	1
est	1
militate	1
potentia	1
Scientia	1
Vivere	1
Vivere	1
```
**Sample Output:**
```
cogitare	1
est	3
militate	1
potentia	1
Scientia	1
Vivere	2
```