# notes-api-demo

This is simple notes api for demo purpose.\
Copied from this tutorial: https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34

## Installation

1. Create .env file, example:

```
DB_HOST=localhost
DB_NAME=fiber-app
DB_USER=postgres
DB_PASSWORD=postgres
DB_PORT=5432
```

2. Setting up your database by running this command:

```
docker compose up -d
```

remember, database environtment in the docker-composeyml file must be the same as the .env file.

3. Run the app
   You can run this command:

```
go mod download
make build
make run
```

Or if you are using docker, you can try this:

```
docker build -t notes-api-demo:latest .
docker run -d -p 3000:3000 \
    -e DB_HOST=host.docker.internal \
    -e DB_NAME=fiber-app \
    -e DB_USER=postgres \
    -e DB_PASSWORD=postgres \
    -e DB_PORT=5432 notes-api-demo:v1

```

## Api Documentation

1. POST `api/note` - create a note.
2. GET `api/note` - get all notes.
3. GET `api/note/:noteId` - get a note by id.
4. PUT `api/note/:noteId` - edit a note by id.
5. DELETE `api/note/:noteId` - delete a note by id.
