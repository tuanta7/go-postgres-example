CREATE INDEX IF NOT EXISTS employee_last_name_idx
ON employees.employee USING btree (last_name);
