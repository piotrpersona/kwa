# kwa

Dead simple awk

## Install

```sh
go install github.com/piotrpersona/kwa
```

## Usage

```sh
$ cat file | kwa 0
$ # displays first column of the file
```

## features and limitations

- works only with stdin
- accepts one argument, the column number indexed from 0
- the idea is to simply call `kwa 0`, `kwa 5` to display 1st and 6th column without need to remember complex awk syntax
