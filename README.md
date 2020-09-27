# Fluoride (WIP)

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)

## About

Robust icon pack management service in the making. Source for the backend.

Written in go, dockerized with a postgres database.
In production, k8s will be used.

You may try out the master if you'd like but please remember that it is in no way ready for use.

## Prerequisites

- Docker (or just Go, if you prefer)
- Unix Shell
- GNU Make

## Usage

Build and run the server with the database using docker-compose

```sh
./run compose
```

Without docker, just run the go main file using

```sh
./run run
```

PS: the `run` file is just a shell script wrapping a Makefile

## Endpoints

`/icons` - `GET`, `POST`
`/icons/count` - `GET`


### NOTE

- Icon requests are referred to as just icons for succinctness.