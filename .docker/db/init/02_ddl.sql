--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6
-- Dumped by pg_dump version 14.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: super_schema; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA super_schema;


ALTER SCHEMA super_schema OWNER TO postgres;

--
-- Name: user_schema; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA user_schema;


ALTER SCHEMA user_schema OWNER TO postgres;

--
-- Name: SCHEMA super_schema; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA super_schema TO app_admin;


--
-- Name: SCHEMA user_schema; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA user_schema TO app_admin;
GRANT ALL ON SCHEMA user_schema TO app_user;


--
-- PostgreSQL database dump complete
--

