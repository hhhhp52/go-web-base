set -e

if( ! test -f Makefile) then
    echo "Please use this script in Project Root,you are in $(pwd) now."
    exit;
fi

cp config/config.yml

scripts/build.sh

scripts/migrate.sh