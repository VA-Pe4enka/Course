-- Tx1: transfer $10 from account 1 to account 2
BEGIN;

UPDATE accounts SET balance = balance - 10 WHERE id = 1 RETURNING *;
UPDATE accounts SET balance = balance + 10 WHERE id = 2 RETURNING *;

ROLLBACK;


-- Tx2: transfer $10 from account 2 to account 1
BEGIN;

UPDATE accounts SET balance = balance + 10 WHERE id = 1 RETURNING *;	
UPDATE accounts SET balance = balance - 10 WHERE id = 2 RETURNING *;

ROLLBACK;

-- Show lock in PostgresSQL
SELECT
    a.application_name,
    l.relation::regclass,
    l.transactionid,
    l.mode,
    l.locktype,
    l.GRANTED,
    a.usename,
    a.query,
    a.pid
FROM
    pg_stat_activity a
    JOIN pg_locks l ON
    l.pid = a.pid
WHERE
    a.application_name = 'psql'
ORDER BY
    a.pid;
