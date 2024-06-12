DROP TABLE IF EXISTS "public"."auditorium_seats";
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

DROP TABLE IF EXISTS "public"."auditoriums";
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

DROP TABLE IF EXISTS "public"."cinemas";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS cinemas_cinema_id_seq;

-- Table Definition
CREATE TABLE "public"."cinemas" (
    "id" int8 NOT NULL DEFAULT nextval('cinemas_cinema_id_seq'::regclass),
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

DROP TABLE IF EXISTS "public"."movies";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS movies_id_seq;

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
    "tmdb_id" int4 NOT NULL,
    "box_office" float8,
    "id" int4 NOT NULL DEFAULT nextval('movies_id_seq'::regclass),
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("imdb_id")
);

DROP TABLE IF EXISTS "public"."shows";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS shows_show_id_seq2;

-- Table Definition
CREATE TABLE "public"."shows" (
    "id" int4 NOT NULL DEFAULT nextval('shows_show_id_seq2'::regclass),
    "date" date,
    "end_time" time,
    "start_time" time,
    "auditorium_id" int8,
    "imdb_id" varchar(20),
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."tickets";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS tickets_id_seq;

-- Table Definition
CREATE TABLE "public"."tickets" (
    "id" int4 NOT NULL DEFAULT nextval('tickets_id_seq'::regclass),
    "seat_number" int4,
    "status" int2 NOT NULL DEFAULT 1,
    "timestamp" timestamp,
    "show_id" int8,
    "user_id" int8,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."users";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_user_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_user_id_seq'::regclass),
    "name" varchar(255),
    "gender" varchar(255),
    "email" varchar(255),
    "password" varchar(255),
    "phone_number" int8,
    "date_of_birth" date,
    "tier" varchar(255) NOT NULL DEFAULT USER,
    "salt" varchar(255),
    "status" int4 NOT NULL DEFAULT 1,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE OR REPLACE FUNCTION public.update_updated_at_user_task()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$ BEGIN NEW.updated_at = now();
RETURN NEW;
END;
$function$
;

