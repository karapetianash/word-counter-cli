Introduction.
This CLI tool counts the number of words, lines or bytes provided as input using the standard input (STDIN) connection. By default, this tool will count the number of words, unless it receives the -l flag, in which case it’ll count the number of lines instead, or the -b (or both -l and -b flags) flag for counting bytes. This version reads data from STDIN and displays the number of words.

Usage.
To build the app, run the following command in the root folder:

> go build .

Above command will generate word-counter-cli file. This name is defined in the go.mod file and it will be the initialized module name.

After that you can run the file using the cmd and pass the text argument for our app to run:

> .\word-counter-cli.exe -b
...some text...