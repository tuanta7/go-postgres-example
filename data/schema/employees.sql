--
-- PostgreSQL database dump
--

\restrict EhO3A0MmkU6zU2J02P5SKjdzWhb2KrLWNpBFt4leDVtpT2LV8ZKpM5D9qeUnR0N

-- Dumped from database version 18.6 (653b019)
-- Dumped by pg_dump version 18.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: employees; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA employees;


--
-- Name: employee_gender; Type: TYPE; Schema: employees; Owner: -
--

CREATE TYPE employees.employee_gender AS ENUM (
    'M',
    'F'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: department; Type: TABLE; Schema: employees; Owner: -
--

CREATE TABLE employees.department (
    id character(4) NOT NULL,
    dept_name character varying(40) NOT NULL
);


--
-- Name: department_employee; Type: TABLE; Schema: employees; Owner: -
--

CREATE TABLE employees.department_employee (
    employee_id bigint NOT NULL,
    department_id character(4) NOT NULL,
    from_date date NOT NULL,
    to_date date NOT NULL
);


--
-- Name: department_manager; Type: TABLE; Schema: employees; Owner: -
--

CREATE TABLE employees.department_manager (
    employee_id bigint NOT NULL,
    department_id character(4) NOT NULL,
    from_date date NOT NULL,
    to_date date NOT NULL
);


--
-- Name: employee; Type: TABLE; Schema: employees; Owner: -
--

CREATE TABLE employees.employee (
    id bigint NOT NULL,
    birth_date date NOT NULL,
    first_name character varying(14) NOT NULL,
    last_name character varying(16) NOT NULL,
    gender employees.employee_gender NOT NULL,
    hire_date date NOT NULL
);


--
-- Name: id_employee_seq; Type: SEQUENCE; Schema: employees; Owner: -
--

CREATE SEQUENCE employees.id_employee_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: id_employee_seq; Type: SEQUENCE OWNED BY; Schema: employees; Owner: -
--

ALTER SEQUENCE employees.id_employee_seq OWNED BY employees.employee.id;


--
-- Name: salary; Type: TABLE; Schema: employees; Owner: -
--

CREATE TABLE employees.salary (
    employee_id bigint NOT NULL,
    amount bigint NOT NULL,
    from_date date NOT NULL,
    to_date date NOT NULL
);


--
-- Name: title; Type: TABLE; Schema: employees; Owner: -
--

CREATE TABLE employees.title (
    employee_id bigint NOT NULL,
    title character varying(50) NOT NULL,
    from_date date NOT NULL,
    to_date date
);


--
-- Name: employee id; Type: DEFAULT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.employee ALTER COLUMN id SET DEFAULT nextval('employees.id_employee_seq'::regclass);


--
-- Name: department idx_16979_primary; Type: CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.department
    ADD CONSTRAINT idx_16979_primary PRIMARY KEY (id);


--
-- Name: department_employee idx_16982_primary; Type: CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.department_employee
    ADD CONSTRAINT idx_16982_primary PRIMARY KEY (employee_id, department_id);


--
-- Name: department_manager idx_16985_primary; Type: CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.department_manager
    ADD CONSTRAINT idx_16985_primary PRIMARY KEY (employee_id, department_id);


--
-- Name: employee idx_16988_primary; Type: CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.employee
    ADD CONSTRAINT idx_16988_primary PRIMARY KEY (id);


--
-- Name: salary idx_16991_primary; Type: CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.salary
    ADD CONSTRAINT idx_16991_primary PRIMARY KEY (employee_id, from_date);


--
-- Name: title idx_16994_primary; Type: CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.title
    ADD CONSTRAINT idx_16994_primary PRIMARY KEY (employee_id, title, from_date);


--
-- Name: idx_16979_dept_name; Type: INDEX; Schema: employees; Owner: -
--

CREATE UNIQUE INDEX idx_16979_dept_name ON employees.department USING btree (dept_name);


--
-- Name: idx_16982_dept_no; Type: INDEX; Schema: employees; Owner: -
--

CREATE INDEX idx_16982_dept_no ON employees.department_employee USING btree (department_id);


--
-- Name: idx_16985_dept_no; Type: INDEX; Schema: employees; Owner: -
--

CREATE INDEX idx_16985_dept_no ON employees.department_manager USING btree (department_id);


--
-- Name: department_employee dept_emp_ibfk_1; Type: FK CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.department_employee
    ADD CONSTRAINT dept_emp_ibfk_1 FOREIGN KEY (employee_id) REFERENCES employees.employee(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: department_employee dept_emp_ibfk_2; Type: FK CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.department_employee
    ADD CONSTRAINT dept_emp_ibfk_2 FOREIGN KEY (department_id) REFERENCES employees.department(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: department_manager dept_manager_ibfk_1; Type: FK CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.department_manager
    ADD CONSTRAINT dept_manager_ibfk_1 FOREIGN KEY (employee_id) REFERENCES employees.employee(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: department_manager dept_manager_ibfk_2; Type: FK CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.department_manager
    ADD CONSTRAINT dept_manager_ibfk_2 FOREIGN KEY (department_id) REFERENCES employees.department(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: salary salaries_ibfk_1; Type: FK CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.salary
    ADD CONSTRAINT salaries_ibfk_1 FOREIGN KEY (employee_id) REFERENCES employees.employee(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: title titles_ibfk_1; Type: FK CONSTRAINT; Schema: employees; Owner: -
--

ALTER TABLE ONLY employees.title
    ADD CONSTRAINT titles_ibfk_1 FOREIGN KEY (employee_id) REFERENCES employees.employee(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict EhO3A0MmkU6zU2J02P5SKjdzWhb2KrLWNpBFt4leDVtpT2LV8ZKpM5D9qeUnR0N

