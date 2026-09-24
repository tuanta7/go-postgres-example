I want to build a learning project focused on PostgreSQL query optimization using Go.

**Project:** E-commerce marketplace benchmark

**Tech stack:**

- Go for backend APIs
- PostgreSQL for persistence
- pgx/v5 and sqlc for database access
- golang-migrate for schema migrations
- k6 for HTTP load testing

**Primary goal:**
Explore real-world PostgreSQL performance problems rather than build a production-ready e-commerce application.

**Core modules:**

1. Product browsing, searching, filtering and pagination
2. Order history and customer activity
3. Seller analytics, revenue reports and leaderboards
4. Inventory reservation and concurrent updates

**Database optimization topics:**

- Sequential scans versus index scans
- B-tree, composite, partial and covering indexes
- COUNT queries and aggregation
- GROUP BY and JOIN optimization
- Window functions: ROW_NUMBER, RANK, LAG, running SUM
- OFFSET versus keyset pagination
- CTEs and subqueries
- Materialized views and pre-aggregation
- Partitioning
- Transaction contention and locking
- Read-versus-write performance tradeoffs

**Benchmarking requirements:**

- Generate realistic datasets, potentially millions of rows.
- Use k6 to measure p50, p95, p99 latency, throughput and errors.
- Use EXPLAIN (ANALYZE, BUFFERS) to inspect query plans.
- Use pg_stat_statements to analyze database query statistics.
- Compare baseline and optimized queries under equivalent workloads.
- Record optimization results and tradeoffs.

**Project conventions:**

- `db/migrations` for SQL migrations
- `db/queries` for sqlc queries
- `db/seed` for dataset generation
- `internal/db/sqlc` for generated Go code
- `tests/load` for k6 scripts
- `benchmarks/results` for benchmark outputs
- `docs` for technical documentation

**Instructions:**

1. Start by designing a minimal but realistic database schema.
2. Propose a set of optimization experiments covering the topics above.
3. Keep the HTTP API implementation simple.
4. Structure each benchmark so that baseline and optimized variants can be compared.
5. Do not optimize queries prematurely; preserve measurable baseline implementations.
6. Explain the underlying PostgreSQL concepts and tradeoffs in the documentation.

First, propose the architecture, schema, and implementation phases. Do not implement the entire system immediately.
