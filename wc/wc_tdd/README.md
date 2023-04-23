# GNU `wc`

## Scope
- Iteratively work on the following, making sure that the wc package implements the functionality without tying it to a CLI.
- Mark each point done after adding test cases.

### Outputs
- [ ] Count of lines - `wc -l testdata/one.txt`
- [ ] Count of words - `wc -w testdata/one.txt`
- [ ] Count of characters - `wc -c testdata/one.txt`
- [ ] All flags - `wc -l -w -c testdata/one.txt`
- [ ] No flags - `wc testdata/one.txt`
- [ ] Combined flags - `wc -lcw testdata/one.txt`
- [ ] Any combination of flags - `wc -lcw testdata/one.txt`
- [ ] Any order for flags or args - `wc -l testdata/one.txt` == `wc testdata/one.txt -l`

### Inputs
- [ ] Single file
- [ ] Multiple file
- [ ] Standard input
- [ ] Standard input - file is passed as `-` character

## References
1. GNU `wc` Documentation
    - https://www.gnu.org/software/coreutils/manual/html_node/wc-invocation.html
2. Usage info via `wc --help`
    ```
    wc --help
    Usage: wc [OPTION]... [FILE]...
      or:  wc [OPTION]... --files0-from=F
    Print newline, word, and byte counts for each FILE, and a total line if
    more than one FILE is specified.  A word is a non-zero-length sequence of
    printable characters delimited by white space.

    With no FILE, or when FILE is -, read standard input.

    The options below may be used to select which counts are printed, always in
    the following order: newline, word, character, byte, maximum line length.
      -c, --bytes            print the byte counts
      -m, --chars            print the character counts
      -l, --lines            print the newline counts
          --files0-from=F    read input from the files specified by
                              NUL-terminated names in file F;
                              If F is - then read names from standard input
      -L, --max-line-length  print the maximum display width
      -w, --words            print the word counts
          --help        display this help and exit
          --version     output version information and exit

    GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
    Report any translation bugs to <https://translationproject.org/team/>
    Full documentation <https://www.gnu.org/software/coreutils/wc>
    or available locally via: info '(coreutils) wc invocation'
    ```
3. Man Page of GNU `wc`
    ```
    NAME
          wc - print newline, word, and byte counts for each file

    SYNOPSIS
          wc [OPTION]... [FILE]...
          wc [OPTION]... --files0-from=F

    DESCRIPTION
          Print newline, word, and byte counts for each FILE, and a total line if more than one FILE is specified.  A word is a non-zero-length sequence of printable characters delimited by white space.

          With no FILE, or when FILE is -, read standard input.

          The options below may be used to select which counts are printed, always in the following order: newline, word, character, byte, maximum line length.

          -c, --bytes
                  print the byte counts

          -m, --chars
                  print the character counts

          -l, --lines
                  print the newline counts

          --files0-from=F
                  read input from the files specified by NUL-terminated names in file F; If F is - then read names from standard input

          -L, --max-line-length
                  print the maximum display width

          -w, --words
                  print the word counts

          --help display this help and exit

          --version
                  output version information and exit

    AUTHOR
          Written by Paul Rubin and David MacKenzie.

    REPORTING BUGS
          GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
          Report any translation bugs to <https://translationproject.org/team/>

    COPYRIGHT
          Copyright © 2022 Free Software Foundation, Inc.  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
          This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent permitted by law.

    SEE ALSO
          Full documentation <https://www.gnu.org/software/coreutils/wc>
          or available locally via: info '(coreutils) wc invocation'

    GNU coreutils 9.1                                                                                                        April 2022                                                                                                                   WC(1)
    ```
