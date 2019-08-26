set -e

if( ! test -f Makefile) then
    echo "Please use this script in Project Root,you are in $(pwd) now."
    exit;
fi


go get -v github.com/rubenv/sql-migrate/...


scripts/build.sh

scripts/migrate.sh