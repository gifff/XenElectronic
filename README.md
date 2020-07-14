# XenElectronic

## Problem Statements

Build a responsive web application using React and have the following features for the MVP:

1. Customers should be able to view the list of the products based on the product categories
2. Customers should be able to add the products to the shopping cart
3. Customers should be able to view the products listed on the shopping cart
4. Customers should be able to remove the products listed on the shopping cart
5. Customers should be able to checkout shopping cart and continue their transaction to payment

## Prerequisites

- Go 1.14
- PostgreSQL 12.3
- Make
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [go-swagger](https://goswagger.io)

## How to run test

Execute the following command:

```shell
$ make test
```

## Migrations

To setup database, first you have to create a new database in your PostgreSQL server.
For example, a database named `xen_electronic` is created

Then, using golang-migrate, execute the following command:

```shell
$ migrate -database 'postgres://user:pass@localhost:5432/xen_electronic?sslmode=disable' -path db/migrations up
```

## How to run

First, set database connection string through environment variable

```shell
$ export DSN='postgres://user:pass@localhost:5432/xen_electronic'
```

Then, compile as binary:

```shell
$ make build
```

The binary will be at `out/` directory.

To run the server, execute the following command:

```shell
$ ./out/xenelectronic-server --port=9000
```

## Getting Help

To know what kind of flags you can use, type the following command:

```shell
$ ./out/xenelectronic-server --help
```
