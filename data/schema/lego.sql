--
-- PostgreSQL database dump
--

\restrict 2nAex4PynVUdSZDgzamAnk2tqiLN8k84PZ2QI7o8vhrdns1yBLraQtdYAu4kEQn

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
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: lego_colors; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lego_colors (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    rgb character varying(6) NOT NULL,
    is_trans character(1) NOT NULL
);


--
-- Name: lego_colors_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.lego_colors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: lego_colors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.lego_colors_id_seq OWNED BY public.lego_colors.id;


--
-- Name: lego_inventories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lego_inventories (
    id integer NOT NULL,
    version integer NOT NULL,
    set_num character varying(255) NOT NULL
);


--
-- Name: lego_inventories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.lego_inventories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: lego_inventories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.lego_inventories_id_seq OWNED BY public.lego_inventories.id;


--
-- Name: lego_inventory_parts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lego_inventory_parts (
    inventory_id integer NOT NULL,
    part_num character varying(255) NOT NULL,
    color_id integer NOT NULL,
    quantity integer NOT NULL,
    is_spare boolean NOT NULL
);


--
-- Name: lego_inventory_sets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lego_inventory_sets (
    inventory_id integer NOT NULL,
    set_num character varying(255) NOT NULL,
    quantity integer NOT NULL
);


--
-- Name: lego_part_categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lego_part_categories (
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


--
-- Name: lego_part_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.lego_part_categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: lego_part_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.lego_part_categories_id_seq OWNED BY public.lego_part_categories.id;


--
-- Name: lego_parts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lego_parts (
    part_num character varying(255) NOT NULL,
    name text NOT NULL,
    part_cat_id integer NOT NULL
);


--
-- Name: lego_sets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lego_sets (
    set_num character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    year integer,
    theme_id integer,
    num_parts integer
);


--
-- Name: lego_themes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lego_themes (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    parent_id integer
);


--
-- Name: lego_themes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.lego_themes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: lego_themes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.lego_themes_id_seq OWNED BY public.lego_themes.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: lego_colors id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_colors ALTER COLUMN id SET DEFAULT nextval('public.lego_colors_id_seq'::regclass);


--
-- Name: lego_inventories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_inventories ALTER COLUMN id SET DEFAULT nextval('public.lego_inventories_id_seq'::regclass);


--
-- Name: lego_part_categories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_part_categories ALTER COLUMN id SET DEFAULT nextval('public.lego_part_categories_id_seq'::regclass);


--
-- Name: lego_themes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_themes ALTER COLUMN id SET DEFAULT nextval('public.lego_themes_id_seq'::regclass);


--
-- Name: lego_colors lego_colors_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_colors
    ADD CONSTRAINT lego_colors_pkey PRIMARY KEY (id);


--
-- Name: lego_inventories lego_inventories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_inventories
    ADD CONSTRAINT lego_inventories_pkey PRIMARY KEY (id);


--
-- Name: lego_part_categories lego_part_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_part_categories
    ADD CONSTRAINT lego_part_categories_pkey PRIMARY KEY (id);


--
-- Name: lego_parts lego_parts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_parts
    ADD CONSTRAINT lego_parts_pkey PRIMARY KEY (part_num);


--
-- Name: lego_sets lego_sets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_sets
    ADD CONSTRAINT lego_sets_pkey PRIMARY KEY (set_num);


--
-- Name: lego_themes lego_themes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lego_themes
    ADD CONSTRAINT lego_themes_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- PostgreSQL database dump complete
--

\unrestrict 2nAex4PynVUdSZDgzamAnk2tqiLN8k84PZ2QI7o8vhrdns1yBLraQtdYAu4kEQn

