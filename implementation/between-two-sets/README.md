https://www.hackerrank.com/challenges/between-two-sets

Consider two sets of positive integers, `A = {a0, a1, ..., a(n-1)}` and `B = {b0, b1, ..., b(m-1)}`. We say that a positive integer, `x`, is between sets `A` and `B` if the following conditions are satisfied:

1. All elements in `A` are factors of `x`.
2. `x` is a factor of all elements in `B`.
Given `A` and `B`, find and print the number of integers (i.e., possible `x`'s) that are between the two sets.

### Input Format

The first line contains two space-separated integers describing the respective values of `n` (the number of elements in set `A`) and `m` (the number of elements in set `B`). 
The second line contains `n` distinct space-separated integers describing `a0, a1, ..., a(n-1)`. 
The third line contains  distinct space-separated integers describing `b0, b1, ..., b(m-1)`.

### Constraints
`1 <= n, m <= 10`

`1 <= a(i) <= 100`

`1 <= b(i) <= 100`

### Output Format

Print the number of integers that are considered to be between `A` and `B`.

### Sample Input
```
2 3
2 4
16 32 96
```
### Sample Output
```
3
```