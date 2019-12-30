#!/usr/bin/env bash
# create admin user who has root permissions
# create develop user to use
echo "Creating mongo users..."
mongo admin --host localhost -u root -p root --eval "db.createUser({user: 'admin', pwd: 'admin', roles: [{role: 'userAdminAnyDatabase', db: 'admin'}]});"
mongo admin -u admin -p admin << EOF
use develop
db.createUser({user: 'develop', pwd: 'develop', roles:[{role:'readWrite',db:'develop'}]})
use test
db.createUser({user: 'test', pwd: 'test', roles:[{role:'readWrite',db:'test'}]})
EOF
echo "Mongo users created."