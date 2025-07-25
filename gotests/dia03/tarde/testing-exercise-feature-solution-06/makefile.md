## Makefiles
Makefiles are a fundamental tool for organizing and automating tasks, such as compilation, building, testing, and more, by defining rules and dependencies in a structured manner.

---

### Targeting specific files for compilation, building.

Let's say you have a source text file named `input.txt`, and you want to use a Makefile to create an uppercase version of it named `output.txt`

```makefile
output.txt: input.txt
    tr '[:lower:]' '[:upper:]' < input.txt > output.txt
```

In this example:
- `target`: output.txt (file to be generated)
- `dependency`: input.txt (source file)
- `command`: tr '[:lower:]' '[:upper:]' < input.txt > output.txt (command to be executed converting the content of input.txt to uppercase and writing it to output.txt)

**Note:**
The command will be executed:
- `target` file `does not exist`
- `target` file is `not up to date` than the `dependency` file

### Targeting specific commands for execution (not files)
```makefile
.PHONY: clean
clean:
    rm -rm dir
```

In this example:
- `.PHONY` is a special keyword that tells make that `clean` is not a file (always execute the command `not up to date`).
- `clean` is the target command you want to execute.
- `rm -rm dir` is the command that removes the `dir` directory.