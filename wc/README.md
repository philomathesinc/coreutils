# `wc`

```sh
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
```