# sqlfmt

A small Go CLI formatter for PostgreSQL `CREATE TABLE` statements.

It formats table definitions using fixed-width columns:

```sql
CREATE TABLE IF NOT EXISTS notes (
    id                  BIGSERIAL           PRIMARY KEY,
    title               TEXT                NOT NULL,
    body                TEXT                NOT NULL DEFAULT '',
    created_at          TIMESTAMPTZ         NOT NULL DEFAULT now()
)
```

## Build

```bash
make build
```

## Format a file

```bash
./bin/sqlfmt migrations/notes.sql
```

## Test

```bash
make test
```

## VS Code

Install the **Run on Save** extension by `emeraldwalk`, then use:

```json
"emeraldwalk.runonsave": {
    "autoShowOutputPanel": "always",
    "commands": [
      {
        "match": "\\.sql$",
        "cmd": "sqlfmt \"${file}\""
      },
      {
        "match": ".*",
        "cmd": "make -C \"${workspaceFolder}\""
      }
    ]
  }
```