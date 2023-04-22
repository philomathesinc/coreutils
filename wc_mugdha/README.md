## File reading operations
Went over [this webpage](https://www.golinuxcloud.com/golang-read-a-file-methods/) to understand different options for reading files.

### Possible options -
1. `os.Open()` and `file.Read()`
    - We can control the amount of bytes we want to load in the memory.
2. `ioutil.Read()`
    - This library is deprecated, hence didn't explore much on it.
3. `bufio.NewScanner()`
    - This library reads one line at a time, which can make things easier to process.
    - We have no control over what the length of a line is.
    - The webpage says it's performance is good even with large files.

#### Issues faced with `os.Open()` option -
- While trying to calculate word count for every x bytes, the last word gets cut and counted as two different words.
