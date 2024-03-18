#!/bin/bash

sed -i 's/{{LAMBDA_EXEC_PERM}}/$LAMBDA_EXEC_PERM/g' ./hf-lambda.tf;