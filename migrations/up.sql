create table bioInfo(
    ID                      BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text not null,
    surname text not null,
    patronymic text,
    age int not null,
    gender text not null,
    national text not null
);