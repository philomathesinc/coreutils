#### Supported functionaly
* Using `bufio.NewScanner()` for reading file contents.
* Multiple file inputs.
* All args supported in any order and any location (before or after filenames).
* Counting lines, words and characters for each file and printing them.
* Printing total of lines, words and character count.

#### Unsupported functionaly
* Flags cannot be aggregated like `-lwc`.

#### How to run
```
go run main.go -l ../data/first.txt ../data/fourth.txt
```
