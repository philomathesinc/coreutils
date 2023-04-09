# GNU `cut`

## Scope

Iteratively working on the following.

### Inputs
- [ ] Implement fields flag with delimiter - `--fields=LIST --delimiter=DELIM`
- [ ] Implement characters flag - `--characters=LIST`
- [ ] Implement bytes flag - `--bytes=LIST`

### Outputs
- [ ] Implement output delimiter - `--output-delimiter=STRING`
- [ ] Implement output filtering - `--only-delimited`, `--complement`

## Development

If using VSCode, you can go to the Run and Debug screen to Debug the code in various ways.

1. Using the VSCode window to launch the program using different args - Check the [`.vscode/launch.json`](./.vscode/launch.json) to check the different configurations. For example - "Colon;Fields=1", "Tab;Fields=1", etc.

2. Using a separate terminal to launch the binary and debugging in the VSCode window. For this to work, you have to build the binary and then execute using these commands.
    ```
    go build -gcflags=all="-N -l" .

    # Assuming your binary is called `cut`
    # Pass any flags you want to your program after the `--`(double hyphen)
    dlv exec --api-version 2 --headless -l 127.0.0.1:2345 ./cut -- -f 1 testdata/lists.txt
    ```

## References
1. GNU `cut` Documentation
    - https://www.gnu.org/software/coreutils/manual/html_node/cut-invocation.html
    - https://www.gnu.org/software/coreutils/manual/html_node/The-cut-command.html
2. Usage info via `cut --help`
    ```
    Usage: cut OPTION... [FILE]...
    Print selected parts of lines from each FILE to standard output.

    With no FILE, or when FILE is -, read standard input.

    Mandatory arguments to long options are mandatory for short options too.
    -b, --bytes=LIST        select only these bytes
    -c, --characters=LIST   select only these characters
    -d, --delimiter=DELIM   use DELIM instead of TAB for field delimiter
    -f, --fields=LIST       select only these fields;  also print any line
                                that contains no delimiter character, unless
                                the -s option is specified
    -n                      (ignored)
        --complement        complement the set of selected bytes, characters
                                or fields
    -s, --only-delimited    do not print lines not containing delimiters
        --output-delimiter=STRING  use STRING as the output delimiter
                                the default is to use the input delimiter
    -z, --zero-terminated    line delimiter is NUL, not newline
        --help        display this help and exit
        --version     output version information and exit

    Use one, and only one of -b, -c or -f.  Each LIST is made up of one
    range, or many ranges separated by commas.  Selected input is written
    in the same order that it is read, and is written exactly once.
    Each range is one of:

    N     N'th byte, character or field, counted from 1
    N-    from N'th byte, character or field, to end of line
    N-M   from N'th to M'th (included) byte, character or field
    -M    from first to M'th (included) byte, character or field

    GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
    Full documentation <https://www.gnu.org/software/coreutils/cut>
    or available locally via: info '(coreutils) cut invocation'
    ```
3. Man Page of GNU `cut`
    ```
    CUT(1)                                                             User Commands                                                            CUT(1)

    NAME
        cut - remove sections from each line of files

    SYNOPSIS
        cut OPTION... [FILE]...

    DESCRIPTION
        Print selected parts of lines from each FILE to standard output.

        With no FILE, or when FILE is -, read standard input.

        Mandatory arguments to long options are mandatory for short options too.

        -b, --bytes=LIST
                select only these bytes

        -c, --characters=LIST
                select only these characters

        -d, --delimiter=DELIM
                use DELIM instead of TAB for field delimiter

        -f, --fields=LIST
                select only these fields;  also print any line that contains no delimiter character, unless the -s option is specified

        -n     (ignored)

        --complement
                complement the set of selected bytes, characters or fields

        -s, --only-delimited
                do not print lines not containing delimiters

        --output-delimiter=STRING
                use STRING as the output delimiter the default is to use the input delimiter

        -z, --zero-terminated
                line delimiter is NUL, not newline

        --help display this help and exit

        --version
                output version information and exit

        Use one, and only one of -b, -c or -f.  Each LIST is made up of one range, or many ranges separated by commas.  Selected input is written
        in the same order that it is read, and is written exactly once.  Each range is one of:

        N      N'th byte, character or field, counted from 1

        N-     from N'th byte, character or field, to end of line

        N-M    from N'th to M'th (included) byte, character or field

        -M     from first to M'th (included) byte, character or field

    AUTHOR
        Written by David M. Ihnat, David MacKenzie, and Jim Meyering.

    REPORTING BUGS
        GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
        Report any translation bugs to <https://translationproject.org/team/>

    COPYRIGHT
        Copyright © 2023 Free Software Foundation, Inc.  License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
        This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent permitted by law.

    SEE ALSO
        Full documentation <https://www.gnu.org/software/coreutils/cut>
        or available locally via: info '(coreutils) cut invocation'

    GNU coreutils 9.2                                                   March 2023                                                              CUT(1)
    ```