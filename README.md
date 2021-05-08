# Simple Golang REST API

## Install

1. Clone
1. Copy and rename `.env.sample` to `.env`. Set the database and port.
1. Import the `tables.sql` to database to create and insert data to table.
1. Run `go run main.go`

## Routes

| Endpoint | Method | Param | Desc |
| --- | --- | --- | --- |
| `/student` | `GET` | | Mengambil daftar siswa |
| `/student` |`POST`| `name`, `grade` | Menambahkan siswa |

## Test

To run unit test along with coverage, run `go test -v -cover ./...`
