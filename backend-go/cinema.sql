-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS auditorium_seats_id_seq;
-- Table Definition
CREATE TABLE "public"."auditorium_seats" (
    "id" int4 NOT NULL DEFAULT nextval('auditorium_seats_id_seq'::regclass),
    "seat_number" int4,
    "type" int2,
    "auditorium_id" int8,
    "status" int4 NOT NULL DEFAULT 1,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS auditoriums_id_seq;
-- Table Definition
CREATE TABLE "public"."auditoriums" (
    "id" int4 NOT NULL DEFAULT nextval('auditoriums_id_seq'::regclass),
    "name" varchar(255),
    "seats" int4,
    "cinema_id" int8,
    "status" int4 NOT NULL DEFAULT 1,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."cast_member" (
    "imdb_id" varchar(20) NOT NULL,
    "person_id" int4 NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS cinemas_id_seq;
-- Table Definition
CREATE TABLE "public"."cinemas" (
    "id" int4 NOT NULL DEFAULT nextval('cinemas_id_seq'::regclass),
    "address" text,
    "capacity" int4,
    "email" varchar(255),
    "name" varchar(255),
    "phone_number" int8,
    "status" int4 NOT NULL DEFAULT 1,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "owner_id" int8,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS comments_id_seq;
-- Table Definition
CREATE TABLE "public"."comments" (
    "id" int4 NOT NULL DEFAULT nextval('comments_id_seq'::regclass),
    "content" text,
    "timestamp" timestamp,
    "imdb_id" varchar(20),
    "ticket_id" int8,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS companies_id_seq;
-- Table Definition
CREATE TABLE "public"."companies" (
    "id" int4 NOT NULL DEFAULT nextval('companies_id_seq'::regclass),
    "name" varchar(255),
    "logo_path" varchar(255),
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "user_id" int8,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."companies_countries" (
    "company_id" int4 NOT NULL,
    "iso_3166_1" varchar(3) NOT NULL
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."countries" (
    "iso_3166_1" varchar(255) NOT NULL,
    "name" varchar(255),
    PRIMARY KEY ("iso_3166_1")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."genres" (
    "id" int4 NOT NULL,
    "name" varchar(255),
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."languages" (
    "iso_639_1" varchar(2) NOT NULL,
    "english_name" varchar(255),
    "name" varchar(255),
    PRIMARY KEY ("iso_639_1")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."movies" (
    "imdb_id" varchar(20) NOT NULL,
    "awards" varchar(255),
    "dvd" varchar(255),
    "imdb_rating" float8,
    "imdb_votes" int4,
    "metascore" int4,
    "original_title" varchar(255),
    "plot" text,
    "poster" text,
    "rated" varchar(50),
    "released" varchar(255),
    "runtime" int4,
    "tagline" varchar(255),
    "title" varchar(255),
    "type" varchar(50),
    "website" varchar(255) DEFAULT 'N/A'::character varying,
    "year" int4,
    "production" varchar(255),
    "tmdb_id" int4,
    "box_office" float8,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("imdb_id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."movies_companies" (
    "imdb_id" varchar(20) NOT NULL,
    "company_id" int4 NOT NULL,
    PRIMARY KEY ("imdb_id", "company_id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."movies_countries" (
    "imdb_id" varchar(20) NOT NULL,
    "iso_3166_1" varchar(3) NOT NULL,
    PRIMARY KEY ("imdb_id", "iso_3166_1")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."movies_genres" (
    "imdb_id" varchar(20) NOT NULL,
    "genre_id" int4 NOT NULL
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."movies_languages" (
    "imdb_id" varchar(20) NOT NULL,
    "iso_639_1" varchar(2) NOT NULL,
    PRIMARY KEY ("imdb_id", "iso_639_1")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."movies_ratings" (
    "imdb_id" varchar(20) NOT NULL,
    "writer_id" int8 NOT NULL
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS people_id_seq1;
-- Table Definition
CREATE TABLE "public"."people" (
    "id" int4 NOT NULL DEFAULT nextval('people_id_seq1'::regclass),
    "name" varchar(255),
    "role" varchar(255),
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."production_company" (
    "imdb_id" varchar(20) NOT NULL,
    "company_id" int4 NOT NULL,
    PRIMARY KEY ("imdb_id", "company_id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Table Definition
CREATE TABLE "public"."production_country" (
    "imdb_id" varchar(20) NOT NULL,
    "iso_3166_1" varchar(3) NOT NULL,
    PRIMARY KEY ("imdb_id", "iso_3166_1")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ratings_id_seq1;
-- Table Definition
CREATE TABLE "public"."ratings" (
    "id" int4 NOT NULL DEFAULT nextval('ratings_id_seq1'::regclass),
    "value" varchar(255),
    "source" varchar(255),
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS show_seats_id_seq;
-- Table Definition
CREATE TABLE "public"."show_seats" (
    "id" int4 NOT NULL DEFAULT nextval('show_seats_id_seq'::regclass),
    "status" int2,
    "auditorium_seat_id" int8,
    "ticket_id" int8,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS shows_show_id_seq;
-- Table Definition
CREATE TABLE "public"."shows" (
    "id" int4 NOT NULL DEFAULT nextval('shows_show_id_seq'::regclass),
    "date" date,
    "end_time" time,
    "start_time" time,
    "auditorium_id" int8,
    "imdb_id" varchar(20),
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS tickets_id_seq;
-- Table Definition
CREATE TABLE "public"."tickets" (
    "id" int4 NOT NULL DEFAULT nextval('tickets_id_seq'::regclass),
    "seat_number" int4,
    "status" int2,
    "timestamp" timestamp,
    "show_id" int8,
    "user_id" int8,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;
-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "name" varchar(255),
    "gender" varchar(255),
    "email" varchar(255),
    "password" varchar(255),
    "phone_number" int8,
    "date_of_birth" date,
    "tier" varchar(255) NOT NULL DEFAULT 0,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
CREATE OR REPLACE FUNCTION public.update_updated_at_user_task() RETURNS trigger LANGUAGE plpgsql AS $function$ BEGIN NEW.updated_at = now();
RETURN NEW;
END;
$function$;