# Go CRUD

A CRUD example written in Go.
This example uses MySQL for the database server.

## Usage

```
Import the go-crud-posts.postman_collection.json file into Postman.
```

Change into the package directory:

```bash
cd go-crud/
```

To create the required database tables:

```bash
go run migrate/migrate.go
```

To run the package using CompileDaemon:

```bash
CompileDaemon -command="./go-crud"
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
