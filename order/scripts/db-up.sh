#!/bin/bash

docker run -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=secret-password \
	-e MYSQL_DATABASE=order mysql
