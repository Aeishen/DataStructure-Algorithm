## 递归

#### 定义：**递归算法是一种直接或者间接调用自身函数或者方法的算法。**

通俗来说，递归算法的实质是把问题分解成规模缩小的同类问题的子问题，然后递归调用方法来表示问题的解。它有如下特点：

- 1. 一个问题的解可以分解为几个子问题的解
- 2. 这个问题与分解之后的子问题，除了数据规模不同，求解思路完全一样
- 3. 存在递归终止条件，即必须有一个明确的递归结束条件，称之为递归出口

#### 经典案例：

1.数组求和：Sum(arr[0...n-1]) = arr[0] + Sum(arr[1...n-1])

2.汉诺塔问题：

3.爬台阶问题/斐波那契数列：