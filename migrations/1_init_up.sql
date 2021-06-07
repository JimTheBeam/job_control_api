CREATE DATABASE job_control;

\c job_control

CREATE TABLE tasks
(
    id            serial       not null unique,
    name          varchar(255) not null unique,
    description   text         not null,
    CONSTRAINT "pk_task_id" PRIMARY KEY (id)
);


CREATE TABLE sub_tasks
(
    id            serial       not null unique,
    name          varchar(255) not null unique,
    description   text         not null,
    task_name     varchar(255) not null,
    CONSTRAINT "pk_subtask_id" PRIMARY KEY (id)   
);


CREATE TABLE costs
(
    id            serial       not null unique,
    name          varchar(255) not null unique,
    costs         varchar(255) not null,
    subtask_name  varchar(255) not null,
    CONSTRAINT "pk_cost_id" PRIMARY KEY (id)   
);