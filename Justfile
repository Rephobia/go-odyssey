# list of all justfile recipes
default:
    just --list --unsorted

fmt:
    go run ./01_fmt

array:
    go run ./02_array

defer:
    go run ./04_defer
