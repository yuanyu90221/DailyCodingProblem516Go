# DailyCodingProblem516Go

## 問題描述

這個問題是由 Zillow 所問。

我們定義一個數字為 **sevenish** 數字，代表這個數字是 7 的次方數或是 7 的次方數之和。 **sevenish** 數列的前面幾個數字是 1, 7, 8, 49, 等等。實作一個演算法來找出第 n 個 **sevenish** 數字。

## 解題思路

已知每個 sevenish 數字都是 7 的次方數或是 7 的平方數之和。

所以每個 sevenish 數字化作以 7 為基數的字串都會符合
1[0,1]*$

而從例子中可以發現

1 -> 0b1 -> 1(以7為基數) -> 1 (7 的 0 次方)

2 -> 0b10 -> 10(以7為基數) -> 7 (7 的 1 次方)

3 -> 0b11 -> 11(以7為基數) -> 8 = 7 + 1 (7 的 1 次方 + 7 的 0 次方)

4 -> 0b100 -> 100(以7為基數) -> 49 = (7 的 2 次方)

n -> n 的 2 進位表示式 -> (n 的 2 進位表示式) 把基數改為 7 轉回 10 進位數值

## 程式碼

```golang
package sevenish

import "strconv"

func FindSevenishNumber(n int64) int64 {
	base2 := strconv.FormatInt(n, 2)
	result, _ := strconv.ParseInt(base2, 7, 64)
	return result
}

```