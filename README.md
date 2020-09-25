# value-vs-ptr
Some benchmarks comparing performace for passing to functions by value and be pointer

```
vf10 109.715966ms
pf10 76.374738ms
vf9 56.761252ms
pf9 66.089294ms
vf7 42.96892ms
pf7 47.186989ms
vf4 23.044532ms
pf4 24.659477ms
vf3 min 17.787846ms
pf3 23.202628ms
vf2 11.741511ms
pf2 13.046917ms
vf1 6.285633ms
pf1 7.534411ms
```

```
vf(n) = pass n field struct by value
```
```
pf(n) = pass n field struct by pointer 
```

On my machine passing by value seems to be faster up the a struct of 9 fields. But for 10 fields passing by pointer is about 50% faster.
