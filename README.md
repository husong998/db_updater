A simple CLI application that upserts DB records using CSV file

## Prerequisite

- MySQL server

## Usage

- Configure mysql connection in `config.toml`

```bash
mysql -uroot < init.sql
./gen.py 100000
go run . --csv=test/test.csv
```