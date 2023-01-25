# Ticket Assignment

- Assign a ticket to an agent, change the ticket status from PENDING to ASSIGNED.
- Change the agent status from INACTIVE to ACTIVE.
- Add a record to the table assignment.
- These three processes should happen in a single transaction.

## The Plan

- Write code to implement transactions, handle deadlocks and ultimately avoid deadlocks.
- TESTIFY : Thou shalt write tests.
- Test Coverage of 80% - 90%

## Schema

![simple-tickets](./Simple%20Tickets.png)

### Migrations

### On the local machine

- These assume psql is installed : pgAdmin used to create the database beforehand

  - init migration : create initial migration

           migrate create -ext sql -dir db/migration -seq init_schema

  - Run migration

           migrate -path db/migrations -database "postgresql://root:password@localhost:5432/ticket-assignment?sslmode=disable" -verbose up

### Using a Docker Container

- Pull postgres alpine image from Dockerhub :

        docker pull postgres:15-alpine

- then run the docker image pulled to create the container in detached mode

        docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

- to access the linux terminal inside the container, run

        docker exec -it postgres15 /bin/sh

- then run

        createdb --username=root --owner=root ticket-assignment

to create the database

- then run

        psql ticket-assignment

to access the PSQL tool

- OR use a single docker command to create the database

        docker exec -it postgres15 createdb --username=root --owner=root ticket-assignment

- then access the PSQL console without using the command shell

        docker exec -it postgres15 psql -U root ticket-assignment

### SQLC

- [SQLC](https://github.com/kyleconroy/sqlc/tree/v1.4.0) generates boiler plate code for CRUD operations on the database using the SQL queries written.

  1.  On the root folder init SQLC to create sql.yaml file

          sqlc init

  . . . the rest of the commands can be located in the Makefile

## Technologies

- Go
- Postgres
- Sqlc
- Others . . .
