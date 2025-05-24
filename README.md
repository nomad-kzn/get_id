# get_id

`get_id` is a lightweight command-line utility written in Go that generates random UUID from CLI.


## Installation

If you have Go installed (version 1.16 or higher), you can install `get_id` globally using:

```bash
go install github.com/nomad-kzn/get_id@latest
```

This will place the `get_id` binary in your `$GOPATH/bin` or `$HOME/go/bin`, which you can add to your `PATH` for global usage:

```bash
export PATH=$PATH:$HOME/go/bin
```

## Usage

Simply run:

```bash
get_id gen
```

Example output:

```bash
a9810d73-4297-401c-a218-0f5f4102b79a
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
