#!/bin/bash

DBSTRING="host=postgres user=postgres dbname=postgres sslmode=disable"
goose postgres "$DBSTRING" up