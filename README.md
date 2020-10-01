# Bin2Dec

Binary-to-Decimal number converter

## Installation

This project uses Go Modules, so you could install running the following:

```bash
go get github.com/mirusky-dev/bin2dec
go install github.com/mirusky-dev/bin2dec
```

## Usage

```bash
# With flag
bin2dec -i 10101010
170

# With env var
export BIN2DEC_INPUT=10101010
bin2dec
170

# With args
bin2dec 10101010
170
```

## License
[MIT](https://choosealicense.com/licenses/mit/)