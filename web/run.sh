#!/bin/bash

go build -o booking ./*.go
./booking -dbname=bookings -dbuser=booking -cache=false -production=false -dbpass=test123