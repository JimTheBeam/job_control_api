CREATE TABLE tasks
(
    id            serial       not null unique,
    name          text not null unique,
    CONSTRAINT "pk_task_id" PRIMARY KEY (id)
);