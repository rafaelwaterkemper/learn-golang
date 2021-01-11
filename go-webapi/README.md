```
docker run \
    --name pggolang \
    -e POSTGRES_USER=golang \
    -e POSTGRES_PASSWORD=golang \
    -e POSTGRES_DB=godb \
    -p 5432:5432 \
    -d \
    postgres

    create table produtos(
        id serial primary key,
        name varchar,
        description varchar,
        price decimal,
        quantity integer
    );
```