--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: users; Type: TABLE; Schema: public; Owner: leonard; Tablespace: 
--

CREATE TABLE users (
    id integer NOT NULL,
    name character varying(64),
    type character varying(32),
    likes integer[]
);


ALTER TABLE users OWNER TO leonard;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: leonard
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO leonard;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: leonard
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: leonard
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: leonard
--

COPY users (id, name, type, likes) FROM stdin;
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: leonard
--

SELECT pg_catalog.setval('users_id_seq', 1, false);


--
-- PostgreSQL database dump complete
--

