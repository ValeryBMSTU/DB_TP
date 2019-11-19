CREATE SCHEMA forum;


ALTER SCHEMA forum OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: forum; Type: TABLE; Schema: forum; Owner: postgres
--

CREATE TABLE forum.forum (
    id integer NOT NULL,
    posts integer DEFAULT 0 NOT NULL,
    slug text NOT NULL,
    threads integer DEFAULT 0 NOT NULL,
    title text NOT NULL,
    "user" text NOT NULL
);


ALTER TABLE forum.forum OWNER TO postgres;

--
-- Name: forum_id_seq; Type: SEQUENCE; Schema: forum; Owner: postgres
--

CREATE SEQUENCE forum.forum_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE forum.forum_id_seq OWNER TO postgres;

--
-- Name: forum_id_seq; Type: SEQUENCE OWNED BY; Schema: forum; Owner: postgres
--

ALTER SEQUENCE forum.forum_id_seq OWNED BY forum.forum.id;


--
-- Name: post; Type: TABLE; Schema: forum; Owner: postgres
--

CREATE TABLE forum.post (
    id integer NOT NULL,
    author text NOT NULL,
    created timestamp with time zone,
    forum text NOT NULL,
    isedited boolean DEFAULT false NOT NULL,
    message text NOT NULL,
    parent integer DEFAULT 0 NOT NULL,
    thread integer NOT NULL
);


ALTER TABLE forum.post OWNER TO postgres;

--
-- Name: post_id_seq; Type: SEQUENCE; Schema: forum; Owner: postgres
--

CREATE SEQUENCE forum.post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE forum.post_id_seq OWNER TO postgres;

--
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: forum; Owner: postgres
--

ALTER SEQUENCE forum.post_id_seq OWNED BY forum.post.id;


--
-- Name: thread; Type: TABLE; Schema: forum; Owner: postgres
--

CREATE TABLE forum.thread (
    id integer NOT NULL,
    author text NOT NULL,
    created timestamp with time zone,
    forum text NOT NULL,
    message text NOT NULL,
    slug text,
    title text NOT NULL,
    votes integer DEFAULT 0 NOT NULL
);


ALTER TABLE forum.thread OWNER TO postgres;

--
-- Name: table_name_id_seq; Type: SEQUENCE; Schema: forum; Owner: postgres
--

CREATE SEQUENCE forum.table_name_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE forum.table_name_id_seq OWNER TO postgres;

--
-- Name: table_name_id_seq; Type: SEQUENCE OWNED BY; Schema: forum; Owner: postgres
--

ALTER SEQUENCE forum.table_name_id_seq OWNED BY forum.thread.id;


--
-- Name: user; Type: TABLE; Schema: forum; Owner: postgres
--

CREATE TABLE forum."user" (
    id integer NOT NULL,
    about text,
    email text NOT NULL,
    fullname text,
    nickname text NOT NULL
);


ALTER TABLE forum."user" OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: forum; Owner: postgres
--

CREATE SEQUENCE forum.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE forum.user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: forum; Owner: postgres
--

ALTER SEQUENCE forum.user_id_seq OWNED BY forum."user".id;


--
-- Name: forum id; Type: DEFAULT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.forum ALTER COLUMN id SET DEFAULT nextval('forum.forum_id_seq'::regclass);


--
-- Name: post id; Type: DEFAULT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.post ALTER COLUMN id SET DEFAULT nextval('forum.post_id_seq'::regclass);


--
-- Name: thread id; Type: DEFAULT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.thread ALTER COLUMN id SET DEFAULT nextval('forum.table_name_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum."user" ALTER COLUMN id SET DEFAULT nextval('forum.user_id_seq'::regclass);


--
-- Name: forum_id_seq; Type: SEQUENCE SET; Schema: forum; Owner: postgres
--

SELECT pg_catalog.setval('forum.forum_id_seq', 1235, true);


--
-- Name: post_id_seq; Type: SEQUENCE SET; Schema: forum; Owner: postgres
--

SELECT pg_catalog.setval('forum.post_id_seq', 26, true);


--
-- Name: table_name_id_seq; Type: SEQUENCE SET; Schema: forum; Owner: postgres
--

SELECT pg_catalog.setval('forum.table_name_id_seq', 1965, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: forum; Owner: postgres
--

SELECT pg_catalog.setval('forum.user_id_seq', 10484, true);


--
-- Name: forum forum_pk; Type: CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.forum
    ADD CONSTRAINT forum_pk PRIMARY KEY (id);


--
-- Name: post post_pk; Type: CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.post
    ADD CONSTRAINT post_pk PRIMARY KEY (id);


--
-- Name: thread table_name_pk; Type: CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.thread
    ADD CONSTRAINT table_name_pk PRIMARY KEY (id);


--
-- Name: user user_pk; Type: CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);


--
-- Name: forum_id_uindex; Type: INDEX; Schema: forum; Owner: postgres
--

CREATE UNIQUE INDEX forum_id_uindex ON forum.forum USING btree (id);


--
-- Name: forum_slug_uindex; Type: INDEX; Schema: forum; Owner: postgres
--

CREATE UNIQUE INDEX forum_slug_uindex ON forum.forum USING btree (slug);


--
-- Name: post_id_uindex; Type: INDEX; Schema: forum; Owner: postgres
--

CREATE UNIQUE INDEX post_id_uindex ON forum.post USING btree (id);


--
-- Name: table_name_id_uindex; Type: INDEX; Schema: forum; Owner: postgres
--

CREATE UNIQUE INDEX table_name_id_uindex ON forum.thread USING btree (id);


--
-- Name: table_name_slug_uindex; Type: INDEX; Schema: forum; Owner: postgres
--

CREATE UNIQUE INDEX table_name_slug_uindex ON forum.thread USING btree (slug);


--
-- Name: user_email_uindex; Type: INDEX; Schema: forum; Owner: postgres
--

CREATE UNIQUE INDEX user_email_uindex ON forum."user" USING btree (email);


--
-- Name: user_id_uindex; Type: INDEX; Schema: forum; Owner: postgres
--

CREATE UNIQUE INDEX user_id_uindex ON forum."user" USING btree (id);


--
-- Name: user_nickname_uindex; Type: INDEX; Schema: forum; Owner: postgres
--

CREATE UNIQUE INDEX user_nickname_uindex ON forum."user" USING btree (nickname);


--
-- Name: forum forum_user_nickname_fk; Type: FK CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.forum
    ADD CONSTRAINT forum_user_nickname_fk FOREIGN KEY ("user") REFERENCES forum."user"(nickname);


--
-- Name: post post_thread_id_fk; Type: FK CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.post
    ADD CONSTRAINT post_thread_id_fk FOREIGN KEY (thread) REFERENCES forum.thread(id);


--
-- Name: post post_user_nickname_fk; Type: FK CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.post
    ADD CONSTRAINT post_user_nickname_fk FOREIGN KEY (author) REFERENCES forum."user"(nickname);


--
-- Name: thread thread_forum_slug_fk; Type: FK CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.thread
    ADD CONSTRAINT thread_forum_slug_fk FOREIGN KEY (forum) REFERENCES forum.forum(slug);


--
-- Name: thread thread_user_nickname_fk; Type: FK CONSTRAINT; Schema: forum; Owner: postgres
--

ALTER TABLE ONLY forum.thread
    ADD CONSTRAINT thread_user_nickname_fk FOREIGN KEY (author) REFERENCES forum."user"(nickname);


--
-- PostgreSQL database dump complete
--
