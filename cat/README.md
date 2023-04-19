# GNU `cat`

## Scope

- Iteratively work on the following.
- Mark each point done after adding test cases.

### Inputs
- [ ] Implement basic working of cat with 1 file (no flags)
- [ ] Implement basic working of cat with multiple files (no flags)
- [ ] Implement basic working of cat with stdin as input (no flags)

### Outputs
- [ ] Concatenate FILE(s) to standard output.

## Development

If using VSCode, you can go to the Run and Debug screen to Debug the code in various ways.

1. Using the VSCode window to launch the program using different args - Check the [`.vscode/launch.json`](./.vscode/launch.json) to check the different configurations.

2. Using a separate terminal to launch the binary and debugging in the VSCode window. For this to work, you have to build the binary and then execute using these commands.
    ```
    go build -gcflags=all="-N -l" .

    # Assuming your binary is called `cat`
    # Pass any flags you want to your program after the `--`(double hyphen)
    dlv exec --api-version 2 --headless -l 127.0.0.1:2345 ./cat -- -f 1 testdata/file1.txt
    ```

## References
1. GNU `cat` Documentation
    - 
2. Usage info via `cat --help`
    ```
    Usage: cat [OPTION]... [FILE]...
    Concatenate FILE(s) to standard output.

    With no FILE, or when FILE is -, read standard input.

    -A, --show-all           equivalent to -vET
    -b, --number-nonblank    number nonempty output lines, overrides -n
    -e                       equivalent to -vE
    -E, --show-ends          display $ at end of each line
    -n, --number             number all output lines
    -s, --squeeze-blank      suppress repeated empty output lines
    -t                       equivalent to -vT
    -T, --show-tabs          display TAB characters as ^I
    -u                       (ignored)
    -v, --show-nonprinting   use ^ and M- notation, except for LFD and TAB
        --help        display this help and exit
        --version     output version information and exit

    Examples:
    gcat f - g  Output f's contents, then standard input, then g's contents.
    gcat        Copy standard input to standard output.

    GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
    Report any translation bugs to <https://translationproject.org/team/>
    Full documentation <https://www.gnu.org/software/coreutils/cat>
    or available locally via: info '(coreutils) cat invocation'
    
    ```
