#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e 

# Check location
if ( ! test -f "Makefile" ) then
    echo "Please use this script in Project Root, your are in $(pwd) now.";
    exit;
fi

DEPLOY_ENV=$1;
DATABASE_DIR="db";
DATABASE_HOST="172.18.0.2"
DATABASE_DB="webtest"
DATABASE_USER="root"
DATABASE_PASSWORD="123456"

echo "[`date "+%Y/%m/%d %H:%M:%S"`] Start Migrating.";

cd $DATABASE_DIR

# https://github.com/rubenv/sql-migrate 
sql-migrate up -env=${1:-development};
mysql -h$DATABASE_HOST -u$DATABASE_USER -p$DATABASE_PASSWORD $DATABASE_DB < trigger/up/*.sql > /dev/null 2>&1

cd ..

echo "[`date "+%Y/%m/%d %H:%M:%S"`] Finished.";