#!/bin/sh

# 确保脚本返回非零的status时，也能正确退出
set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$MIGRATE_DB_URL" -verbose up

echo "start the app"
# 将所有参数传入脚本并执行
exec "$@"