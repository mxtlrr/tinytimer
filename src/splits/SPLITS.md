# Splits in TinyTimer
TinyTimer (TT) supports the use of splits, which helps you time specific
steps in a solve, e.g F2B, CMLL and LSE for roux. TT stores each split's
data in an array of a data type known as `split_t`. It is defined as

<!-- TODO: Implement minutes -->
```golang
    TIME_SECONDS  int
    TIME_MILLISEC int
    NAME string
```

## Setting up Splits
TT looks for a `splits` file in the current working directory.
Alternatively, you can pass `--file [splits]` into TT, and it will look
there instead. Luckily, the format for setting up splits is very easy.

For example, if you were using 3x3 CFOP, your `splits` file would look like
this:
```
cross,f2l,oll,pll
```

## Using splits
Pressing any key will trigger a split change. Once you reach the end of the split count, pressing a key will just exit out the program, showing
your results.