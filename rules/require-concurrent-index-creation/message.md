# Problem

Ensure all index creations use the `CONCURRENTLY` option.
This rule ignores indexes added to tables created in the same transaction.
During a normal index creation updates are blocked. `CONCURRENTLY` avoids the issue of blocking.

[SQL-CREATEINDEX-CONCURRENTLY](https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY)

# Solutions

Instead of:
```sql
CREATE INDEX "email_idx" ON "app_user" ("email");
```

Use:
```sql
CREATE INDEX CONCURRENTLY "email_idx" ON "app_user" ("email");
```
