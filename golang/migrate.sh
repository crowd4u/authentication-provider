#!/bin/bash
if [ $# -ne 1 ]; then
  echo "Wrong Parameter!"
  exit 1
fi
time=`date +%s`; touch "${time}_$1.sql"
echo "Create migration file"
