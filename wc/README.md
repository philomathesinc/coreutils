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