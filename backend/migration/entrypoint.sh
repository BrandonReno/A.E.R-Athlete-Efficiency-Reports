#!/bin/bash

DBSTRING="host=postgres user=postgres dbname=postgres sslmode=disable"
echo $DBSTRING
goose postgres "$DBSTRING" up