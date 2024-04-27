# TinyTimer
TinyTimer (TT) is a little terminal-based timer, written in Go, with no
additional dependencies.


## Compilation

## Usage
*A man page,* `man 6 ./tinytimer` *is available if you use \*NIX.*


```
tinytimer [--file file] [--log logfile]
```

The `--file` switch is for the splits. By default, TT looks for a `splits.txt`.
If you're curious on how it works, read [SPLITS.md](./src/splits/SPLITS.md).

Next, the `--log` switch is for if you want to log your solve to a log file.
By default, this is turned off.