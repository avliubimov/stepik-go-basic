# [2.5.2 Число десятков числа](https://stepik.org/lesson/917015/step/6?unit=922794)

Дано неотрицательное целое число. Найдите число десятков(то есть вторую справа цифру).

|Формат входных данных|Формат выходных данных|
|---|---|
|На вход дается натуральное число, не превосходящее `10000`.|Выведите одно целое число — число десятков.|

<details>
<summary>Подсказка</summary>

Сначала получите число без последней цифры, а затем возьмите у получившегося числа последнюю цифру. Например, из `1024` получите число `102`, а затем возьмите последнюю цифру.

</details>

___
|№|Sample Input|Sample Output|
|---|---:|---:|
|1|`1024`|`2`|
|2|`1`|`0`|
|3|`10`|`1`|
|4|`99`|`9`|
