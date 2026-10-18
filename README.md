# go-postgres-example

- Postgres Instance: [Neon Console](https://console.neon.tech/app/projects/winter-rice-15802949/branches/br-calm-heart-b3tkqzbn)
- Sample Data: [neondatabase/postgres-sample-dbs](https://github.com/neondatabase/postgres-sample-dbs)

## Examples

- http://localhost:8080/api/v1/employees/search/explain?last_name=Denti
- http://localhost:8080/api/v1/employees/search/explain-analyze?last_name=Denti

### Without Index

```json
{
  "plan": [
    "Gather  (cost=1000.00..5717.26 rows=182 width=35)",
    "  Workers Planned: 1",
    "  ->  Parallel Seq Scan on employee  (cost=0.00..4699.06 rows=107 width=35)",
    "        Filter: ((last_name)::text = 'Denti'::text)"
  ]
}

{
  "plan": [
    "Gather  (cost=1000.00..5717.26 rows=182 width=35) (actual time=0.397..18.302 rows=188.00 loops=1)",
    "  Workers Planned: 1",
    "  Workers Launched: 1",
    "  Buffers: shared hit=2493",
    "  ->  Parallel Seq Scan on employee  (cost=0.00..4699.06 rows=107 width=35) (actual time=0.269..12.820 rows=94.00 loops=2)",
    "        Filter: ((last_name)::text = 'Denti'::text)",
    "        Rows Removed by Filter: 149918",
    "        Buffers: shared hit=2493",
    "Planning:",
    "  Buffers: shared hit=6 dirtied=1",
    "Planning Time: 0.067 ms",
    "Execution Time: 18.343 ms"
  ]
}
```

### With Index

The execution plan contains two operations:

1. Bitmap Index Scan — Finds matching row locations using the index.
2. Bitmap Heap Scan — Retrieves the actual rows from the table.

```json
{
  "plan": [
    "Bitmap Heap Scan on employee  (cost=5.83..571.82 rows=182 width=35)",
    "  Recheck Cond: ((last_name)::text = 'Denti'::text)",
    "  ->  Bitmap Index Scan on employee_last_name_idx  (cost=0.00..5.79 rows=182 width=0)",
    "        Index Cond: ((last_name)::text = 'Denti'::text)"
  ]
}

{
  "plan": [
    "Bitmap Heap Scan on employee  (cost=5.83..571.82 rows=182 width=35) (actual time=0.157..0.544 rows=188.00 loops=1)",
    "  Recheck Cond: ((last_name)::text = 'Denti'::text)",
    "  Heap Blocks: exact=180",
    "  Buffers: shared hit=180 read=3",
    "  ->  Bitmap Index Scan on employee_last_name_idx  (cost=0.00..5.79 rows=182 width=0) (actual time=0.127..0.127 rows=188.00 loops=1)",
    "        Index Cond: ((last_name)::text = 'Denti'::text)",
    "        Index Searches: 1",
    "        Buffers: shared read=3",
    "Planning Time: 0.065 ms",
    "Execution Time: 0.585 ms"
  ]
}
```
