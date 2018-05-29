#!/bin/sh
#Generate application keypair
openssl genrsa -out sermng.rsa 1024
openssl rsa -in sermng.rsa -pubout > sermng.rsa.pub
