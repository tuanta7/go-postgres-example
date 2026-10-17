-- name: GetEmployeeByID :one
SELECT *
FROM employees.employee
WHERE id = $1;

-- name: CountEmployeesByID :one
SELECT count(*)
FROM employees.employee
WHERE id = $1;
