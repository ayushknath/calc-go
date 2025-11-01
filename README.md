# Calculator

A simple command line calculator written in Go

## Build instructions

Follow the instructions and copy the commands below to build the app

1. Clone this repo

```bash
git clone https://github.com/ayushknath/calc-go.git
```

2. Move to the `src` directory and build the app

```bash
cd ./calc-go/src
go build -o ../build/calc
```

3. The executable is under the `build` directory

```bash
cd ../build
./calc add 2 3
```

## Features

The app has two modes
1. **Normal mode**: Here you can perform simple calculations by specifying the operation and the arguments (numbers) to operate upon
2. **Interactive mode**: Spins up a REPL where you can write "not-so-complex" expressions and get their result

### Normal mode

Simple calculations can be done with built-in commands, each of which accept two arguments like:
1. `./calc add 2 3` for addition
2. `./calc sub 2 3` for subtraction
3. `./calc mul 2 3` for multiplication
4. `./calc div 2 3` for division
5. `./calc exp 2 3` for exponentiation

### Interactive mode

To enter the REPL, type `./calc -i`. To quit the REPL, type `quit`

```bash
$ ./calc -i
Interactive mode

>> 2 + 3 * (8 - 4)
14.00
>> quit
$
```
