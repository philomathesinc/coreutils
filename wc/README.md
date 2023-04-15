# `wc`

Follow along repo for the exercise at https://playbook.one2n.in/go-bootcamp/go-projects/word-count-exercise.

Copy of the problem statement exists here - [Problem Statement](./ProblemStatement.md)

## Random helpful commands

```sh
# The GNU wc installed via `coreutils` package(https://formulae.brew.sh/formula/coreutils) is invoked as gwc
$ echo -n "hello world" | gwc -m
11
$ echo "hello world" | gwc -m
12
$ printf "one.txt\0two.txt\0" > filelist.txt
$ cat filelist.txt
$ wc --files0-from=filelist.txt
  2  12  71 one.txt
  5  24 143 two.txt
  7  36 214 total
$ touch testdata/no-read.txt
$ chmod 333 testdata/no-read.txt
```

## Benchmark

Used [Hyperfine](https://github.com/sharkdp/hyperfine) against GNU `wc`.
```
➜  wc git:(wc_combined) ✗ hyperfine 'wc single.txt' './wc single.txt'                                 
Benchmark 1: wc single.txt
  Time (mean ± σ):      32.4 ms ±   3.5 ms    [User: 29.4 ms, System: 2.1 ms]
  Range (min … max):    25.3 ms …  43.3 ms    111 runs
 
Benchmark 2: ./wc single.txt
  Time (mean ± σ):      94.1 ms ±   8.6 ms    [User: 82.8 ms, System: 36.8 ms]
  Range (min … max):    84.5 ms … 130.2 ms    35 runs
 
Summary
  'wc single.txt' ran
    2.90 ± 0.41 times faster than './wc single.txt'
```
