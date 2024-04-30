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
    "tmdb_id" int4,
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

INSERT INTO "public"."auditorium_seats" ("id", "seat_number", "type", "auditorium_id", "status", "created_at") VALUES
(1, 1, 0, 1, 1, '2024-03-20 08:43:22.318223');
INSERT INTO "public"."auditorium_seats" ("id", "seat_number", "type", "auditorium_id", "status", "created_at") VALUES
(2, 2, 0, 1, 1, '2024-03-20 08:43:22.318223');
INSERT INTO "public"."auditorium_seats" ("id", "seat_number", "type", "auditorium_id", "status", "created_at") VALUES
(3, 3, 0, 1, 1, '2024-03-20 08:43:22.318223');
INSERT INTO "public"."auditorium_seats" ("id", "seat_number", "type", "auditorium_id", "status", "created_at") VALUES
(4, 4, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(5, 5, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(6, 6, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(7, 7, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(8, 8, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(9, 9, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(10, 10, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(11, 11, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(12, 12, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(13, 13, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(14, 14, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(15, 15, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(16, 16, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(17, 17, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(18, 18, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(19, 19, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(20, 20, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(21, 21, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(22, 22, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(23, 23, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(24, 24, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(25, 25, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(26, 26, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(27, 27, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(28, 28, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(29, 29, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(30, 30, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(31, 31, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(32, 32, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(33, 33, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(34, 34, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(35, 35, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(36, 36, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(37, 37, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(38, 38, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(39, 39, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(40, 40, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(41, 41, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(42, 42, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(43, 43, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(44, 44, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(45, 45, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(46, 46, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(47, 47, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(48, 48, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(49, 49, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(50, 50, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(51, 51, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(52, 52, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(53, 53, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(54, 54, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(55, 55, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(56, 56, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(57, 57, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(58, 58, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(59, 59, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(60, 60, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(61, 61, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(62, 62, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(63, 63, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(64, 64, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(65, 65, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(66, 66, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(67, 67, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(68, 68, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(69, 69, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(70, 70, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(71, 71, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(72, 72, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(73, 73, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(74, 74, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(75, 75, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(76, 76, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(77, 77, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(78, 78, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(79, 79, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(80, 80, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(81, 81, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(82, 82, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(83, 83, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(84, 84, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(85, 85, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(86, 86, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(87, 87, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(88, 88, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(89, 89, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(90, 90, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(91, 91, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(92, 92, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(93, 93, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(94, 94, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(95, 95, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(96, 96, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(97, 97, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(98, 98, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(99, 99, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(100, 100, 0, 1, 1, '2024-03-20 08:43:22.318223'),
(101, 1, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(102, 2, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(103, 3, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(104, 4, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(105, 5, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(106, 6, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(107, 7, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(108, 8, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(109, 9, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(110, 10, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(111, 11, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(112, 12, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(113, 13, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(114, 14, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(115, 15, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(116, 16, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(117, 17, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(118, 18, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(119, 19, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(120, 20, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(121, 21, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(122, 22, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(123, 23, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(124, 24, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(125, 25, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(126, 26, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(127, 27, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(128, 28, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(129, 29, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(130, 30, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(131, 31, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(132, 32, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(133, 33, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(134, 34, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(135, 35, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(136, 36, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(137, 37, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(138, 38, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(139, 39, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(140, 40, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(141, 41, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(142, 42, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(143, 43, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(144, 44, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(145, 45, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(146, 46, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(147, 47, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(148, 48, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(149, 49, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(150, 50, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(151, 51, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(152, 52, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(153, 53, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(154, 54, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(155, 55, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(156, 56, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(157, 57, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(158, 58, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(159, 59, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(160, 60, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(161, 61, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(162, 62, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(163, 63, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(164, 64, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(165, 65, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(166, 66, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(167, 67, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(168, 68, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(169, 69, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(170, 70, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(171, 71, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(172, 72, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(173, 73, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(174, 74, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(175, 75, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(176, 76, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(177, 77, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(178, 78, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(179, 79, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(180, 80, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(181, 81, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(182, 82, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(183, 83, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(184, 84, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(185, 85, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(186, 86, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(187, 87, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(188, 88, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(189, 89, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(190, 90, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(191, 91, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(192, 92, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(193, 93, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(194, 94, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(195, 95, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(196, 96, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(197, 97, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(198, 98, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(199, 99, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(200, 100, 0, 2, 1, '2024-03-20 08:43:22.318223'),
(201, 1, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(202, 2, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(203, 3, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(204, 4, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(205, 5, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(206, 6, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(207, 7, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(208, 8, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(209, 9, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(210, 10, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(211, 11, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(212, 12, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(213, 13, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(214, 14, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(215, 15, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(216, 16, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(217, 17, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(218, 18, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(219, 19, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(220, 20, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(221, 21, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(222, 22, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(223, 23, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(224, 24, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(225, 25, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(226, 26, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(227, 27, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(228, 28, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(229, 29, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(230, 30, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(231, 31, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(232, 32, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(233, 33, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(234, 34, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(235, 35, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(236, 36, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(237, 37, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(238, 38, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(239, 39, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(240, 40, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(241, 41, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(242, 42, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(243, 43, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(244, 44, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(245, 45, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(246, 46, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(247, 47, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(248, 48, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(249, 49, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(250, 50, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(251, 51, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(252, 52, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(253, 53, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(254, 54, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(255, 55, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(256, 56, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(257, 57, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(258, 58, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(259, 59, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(260, 60, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(261, 61, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(262, 62, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(263, 63, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(264, 64, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(265, 65, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(266, 66, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(267, 67, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(268, 68, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(269, 69, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(270, 70, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(271, 71, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(272, 72, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(273, 73, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(274, 74, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(275, 75, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(276, 76, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(277, 77, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(278, 78, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(279, 79, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(280, 80, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(281, 81, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(282, 82, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(283, 83, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(284, 84, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(285, 85, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(286, 86, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(287, 87, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(288, 88, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(289, 89, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(290, 90, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(291, 91, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(292, 92, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(293, 93, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(294, 94, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(295, 95, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(296, 96, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(297, 97, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(298, 98, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(299, 99, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(300, 100, 0, 3, 1, '2024-03-20 08:43:22.318223'),
(301, 1, 0, 2, 1, '2024-04-30 16:22:43.175586'),
(302, 2, 0, 2, 1, '2024-04-30 16:22:43.175586');

INSERT INTO "public"."auditoriums" ("id", "name", "seats", "cinema_id", "status", "created_at", "updated_at") VALUES
(1, 'string', 0, 1, 1, '2024-04-30 16:22:05.326301', '2024-04-30 09:22:05.325917');
INSERT INTO "public"."auditoriums" ("id", "name", "seats", "cinema_id", "status", "created_at", "updated_at") VALUES
(2, 'string', 2, 3, 1, '2024-04-30 16:22:43.173299', '2024-04-30 09:22:43.17268');
INSERT INTO "public"."auditoriums" ("id", "name", "seats", "cinema_id", "status", "created_at", "updated_at") VALUES
(3, 'Auditorium 1', 100, 3, 1, '2024-04-30 10:40:05.317203', '2024-04-30 10:40:05.317203');
INSERT INTO "public"."auditoriums" ("id", "name", "seats", "cinema_id", "status", "created_at", "updated_at") VALUES
(4, 'Auditorium 2', 100, 3, 1, '2024-04-30 10:40:05.333123', '2024-04-30 10:40:05.333123'),
(5, 'Auditorium 3', 100, 3, 1, '2024-04-30 10:40:05.339624', '2024-04-30 10:40:05.339624'),
(6, 'Auditorium 4', 100, 3, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(7, 'Auditorium 5', 100, 3, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(8, 'Auditorium 6', 100, 3, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(9, 'Auditorium 7', 100, 3, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(10, 'Auditorium 8', 100, 3, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(11, 'Auditorium 9', 100, 3, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(12, 'Auditorium 10', 100, 3, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(13, 'Auditorium 1', 100, 4, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(14, 'Auditorium 2', 100, 4, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171'),
(15, 'Auditorium 3', 100, 4, 1, '2024-04-30 10:40:05.343171', '2024-04-30 10:40:05.343171');

INSERT INTO "public"."cinemas" ("id", "address", "capacity", "email", "name", "phone_number", "status", "created_at", "updated_at", "owner_id") VALUES
(3, '102 August Revolution, Ward 8, District 3, Ho Chi Minh City', 10, 'feliz@cinema.com', 'Feliz Navidad', 13291939021, 1, '2024-04-08 15:17:27.323234', '2024-04-08 15:17:27.332426', 2);
INSERT INTO "public"."cinemas" ("id", "address", "capacity", "email", "name", "phone_number", "status", "created_at", "updated_at", "owner_id") VALUES
(4, '123 Rose Street, Ward 7, Phu Nhuan District', 3, 'movieplex@cinema.com', 'Movieplex', 928183901, 1, '2024-04-08 15:17:27.323234', '2024-04-08 15:17:27.332426', 3);


INSERT INTO "public"."movies" ("imdb_id", "awards", "dvd", "imdb_rating", "imdb_votes", "metascore", "original_title", "plot", "poster", "rated", "released", "runtime", "tagline", "title", "type", "website", "year", "production", "tmdb_id", "box_office", "id", "created_at") VALUES
('tt9663764', 'N/A', NULL, NULL, NULL, NULL, 'Aquaman and the Lost Kingdom', 'Black Manta, still driven by the need to avenge his father''s death and wielding the power of the mythic Black Trident, will stop at nothing to take Aquaman down once and for all. To defeat him, Aquaman must turn to his imprisoned brother Orm, the former King of Atlantis, to forge an unlikely alliance in order to save the world from irreversible destruction.', 'https://image.tmdb.org/t/p/original/8xV47NDrjdZDpkVcCFqkdHa3T0C.jpg', 'PG-13', '2023-12-20', 124, 'The tide is turning.', 'Aquaman and the Lost Kingdom', 'movie', 'http://www.aquamanmovie.com', 2023, NULL, 572802, 0, 1, '2024-04-30 03:39:39.440288');
INSERT INTO "public"."movies" ("imdb_id", "awards", "dvd", "imdb_rating", "imdb_votes", "metascore", "original_title", "plot", "poster", "rated", "released", "runtime", "tagline", "title", "type", "website", "year", "production", "tmdb_id", "box_office", "id", "created_at") VALUES
('tt10545704', '1 win', NULL, 7, 2, NULL, 'Adieu Monsieur Haffmann', 'Paris 1942. François Mercier is an ordinary man who only aspires to start a family with the woman he loves, Blanche. He is also the employee of a talented jeweler, Mr. Haffmann. But faced with the German occupation, the two men will have no other choice but to conclude an agreement whose consequences, over the months, will upset the fate of our three characters.', 'https://image.tmdb.org/t/p/original/7PnzcmLH5j4gek0Z8OT3uTilAQT.jpg', 'TV-14', '2022-01-12', 115, '', 'Farewell Mister Haffmann', 'movie', '', 2022, NULL, 670243, 0, 2, '2024-04-30 03:39:39.440288');
INSERT INTO "public"."movies" ("imdb_id", "awards", "dvd", "imdb_rating", "imdb_votes", "metascore", "original_title", "plot", "poster", "rated", "released", "runtime", "tagline", "title", "type", "website", "year", "production", "tmdb_id", "box_office", "id", "created_at") VALUES
('tt0108052', 'Won 7 Oscars. 91 wins & 49 nominations total', '05 Mar 2013', 9, 1, 95, 'Schindler''s List', 'The true story of how businessman Oskar Schindler saved over a thousand Jewish lives from the Nazis while they worked as slaves in his factory during World War II.', 'https://image.tmdb.org/t/p/original/sF1U4EUQS8YHUYjNl3pMGNIQyr0.jpg', 'R', '1993-12-15', 195, 'Whoever saves one life, saves the world entire.', 'Schindler''s List', 'movie', 'http://www.schindlerslist.com/', 1993, NULL, 424, 321365567, 3, '2024-04-30 03:39:39.440288');
INSERT INTO "public"."movies" ("imdb_id", "awards", "dvd", "imdb_rating", "imdb_votes", "metascore", "original_title", "plot", "poster", "rated", "released", "runtime", "tagline", "title", "type", "website", "year", "production", "tmdb_id", "box_office", "id", "created_at") VALUES
('tt21064584', '2 wins & 3 nominations', NULL, NULL, NULL, NULL, 'The Iron Claw', 'The true story of the inseparable Von Erich brothers, who made history in the intensely competitive world of professional wrestling in the early 1980s. Through tragedy and triumph, under the shadow of their domineering father and coach, the brothers seek larger-than-life immortality on the biggest stage in sports.', 'https://image.tmdb.org/t/p/original/lsdismtCDga4vsKnmliz0h7CaZT.jpg', 'N/A', '2023-12-21', 132, 'Sons. Brothers. Champions.', 'The Iron Claw', 'movie', 'https://a24films.com/films/the-iron-claw', 2023, NULL, 850165, 0, 4, '2024-04-30 03:39:39.440288'),
('tt5537002', '6 wins & 3 nominations', NULL, 8, 98, 89, 'Killers of the Flower Moon', 'When oil is discovered in 1920s Oklahoma under Osage Nation land, the Osage people are murdered one by one—until the FBI steps in to unravel the mystery.', 'https://image.tmdb.org/t/p/original/dB6Krk806zeqd0YNp2ngQ9zXteH.jpg', 'R', '2023-10-18', 206, 'Greed is an animal that hungers for blood.', 'Killers of the Flower Moon', 'movie', 'https://www.killersoftheflowermoonmovie.com', 2023, NULL, 466420, 155500000, 5, '2024-04-30 03:39:39.440288'),
('tt15398776', '7 wins & 3 nominations', '21 Nov 2023', 8.5, 529, 88, 'Oppenheimer', 'The story of J. Robert Oppenheimer''s role in the development of the atomic bomb during World War II.', 'https://image.tmdb.org/t/p/original/8Gxv8gSFCU0XGDykEGv7zR1n2ua.jpg', 'R', '2023-07-19', 181, 'The world forever changes.', 'Oppenheimer', 'movie', 'http://www.oppenheimermovie.com', 2023, NULL, 872585, NULL, 6, '2024-04-30 03:39:39.440288'),
('tt1517268', '5 wins & 9 nominations', '12 Sep 2023', 7, 383, 80, 'Barbie', 'Barbie and Ken are having the time of their lives in the colorful and seemingly perfect world of Barbie Land. However, when they get a chance to go to the real world, they soon discover the joys and perils of living among humans.', 'https://image.tmdb.org/t/p/original/iuFNMS8U5cb6xfzi51Dbkovj7vM.jpg', 'PG-13', '2023-07-19', 114, 'She''s everything. He''s just Ken.', 'Barbie', 'movie', 'https://www.barbie-themovie.com', 2023, NULL, 346698, NULL, 7, '2024-04-30 03:39:39.440288'),
('tt29493007', 'N/A', NULL, NULL, 86, NULL, 'Đất Rừng Phương Nam', 'The film version of the film is inherited and developed from the novel of the same name by writer Doan Gioi. The film tells the adventure journey of An - a boy who unfortunately lost his mother on the way to find his father. Together with An, the audience will experience the richness of nature and the unique cultural beauty of the land of Nam Ky Luc Tinh, the chivalrousness of the farmers clinging to the forest and the patriotism against the French first. 20th century.', 'https://image.tmdb.org/t/p/original/apFcKjQlxzFj2MY9d0rxdyT0ykv.jpg', 'N/A', '2023-10-20', 110, '', 'Song Of The South', 'movie', '', 2023, NULL, 1124863, NULL, 8, '2024-04-30 03:39:39.440288'),
('tt10298810', '2 wins & 22 nominations', '03 Aug 2022', 6.1, 120, 60, 'Lightyear', 'Legendary Space Ranger Buzz Lightyear embarks on an intergalactic adventure alongside a group of ambitious recruits and his robot companion Sox.', 'https://image.tmdb.org/t/p/original/b9t3w1loraDh7hjdWmpc9ZsaYns.jpg', 'PG', '2022-06-15', 105, 'Infinity awaits.', 'Lightyear', 'movie', 'https://movies.disney.com/lightyear', 2022, NULL, 718789, NULL, 9, '2024-04-30 03:39:39.440288'),
('tt0109830', 'Won 6 Oscars. 50 wins & 74 nominations total', '01 Aug 2013', 8.8, 2, 82, 'Forrest Gump', 'A man with a low IQ has accomplished great things in his life and been present during significant historic events—in each case, far exceeding what anyone imagined he could do. But despite all he has achieved, his one true love eludes him.', 'https://image.tmdb.org/t/p/original/arw2vcBveWOVZr6pxd9XTd1TdQa.jpg', 'PG-13', '1994-06-23', 142, 'The world will never be the same once you''ve seen it through the eyes of Forrest Gump.', 'Forrest Gump', 'movie', 'https://www.paramountmovies.com/movies/forrest-gump', 1994, NULL, 13, 677387716, 10, '2024-04-30 03:39:39.440288'),
('tt13287846', '1 nomination', NULL, NULL, NULL, NULL, 'Napoleon', 'An epic that details the checkered rise and fall of French Emperor Napoleon Bonaparte and his relentless journey to power through the prism of his addictive, volatile relationship with his wife, Josephine.', 'https://image.tmdb.org/t/p/original/jE5o7y9K6pZtWNNMEw3IdpHuncR.jpg', 'R', '2023-11-22', 158, 'He came from nothing. He conquered everything.', 'Napoleon', 'movie', 'https://www.napoleon.movie', 2023, NULL, 753342, NULL, 11, '2024-04-30 03:39:39.440288'),
('tt14230458', '3 wins & 5 nominations', NULL, 8.5, 3, NULL, 'Poor Things', 'Brought back to life by an unorthodox scientist, a young woman runs off with a debauched lawyer on a whirlwind adventure across the continents. Free from the prejudices of her times, she grows steadfast in her purpose to stand for equality and liberation.', 'https://image.tmdb.org/t/p/original/mPdeQ1H6IXDAXtwQ2EdQuSCNmwV.jpg', 'R', '2023-12-07', 141, 'She''s like nothing you''ve ever seen.', 'Poor Things', 'movie', 'https://www.searchlightpictures.com/films/poor-things', 2023, NULL, 792307, 1400000, 12, '2024-04-30 03:39:39.440288'),
('tt0120338', 'Won 11 Oscars. 126 wins & 83 nominations total', '01 Jun 2014', 7.9, 1, 75, 'Titanic', '101-year-old Rose DeWitt Bukater tells the story of her life aboard the Titanic, 84 years later. A young Rose boards the ship with her mother and fiancé. Meanwhile, Jack Dawson and Fabrizio De Rossi win third-class tickets aboard the ship. Rose tells the whole story from Titanic''s departure through to its death—on its first and last voyage—on April 15, 1912.', 'https://image.tmdb.org/t/p/original/9xjZS2rlVxm8SFx8kPC3aIGCOYQ.jpg', 'PG-13', '1997-11-18', 194, 'Nothing on Earth could come between them.', 'Titanic', 'movie', 'https://www.paramountmovies.com/movies/titanic', 1997, NULL, 597, 2264162353, 13, '2024-04-30 03:39:39.440288'),
('tt6166392', '1 win & 4 nominations', NULL, NULL, 635, NULL, 'Wonka', 'Willy Wonka – chock-full of ideas and determined to change the world one delectable bite at a time – is proof that the best things in life begin with a dream, and if you’re lucky enough to meet Willy Wonka, anything is possible.', 'https://image.tmdb.org/t/p/original/qhb1qOilapbapxWQn9jtRCMwXJF.jpg', 'PG', '2023-12-06', 117, 'Every good thing in this world started with a dream.', 'Wonka', 'movie', 'https://www.wonkamovie.com', 2023, NULL, 787699, 151569322, 14, '2024-04-30 03:39:39.440288'),
('tt23289160', 'N/A', NULL, 8.4, 1, NULL, 'ゴジラ-1.0', 'In postwar Japan, a new terror rises. Will the devastated people be able to survive... let alone fight back?', 'https://image.tmdb.org/t/p/original/hkxxMIGaiCTmrEArK7J56JTKUlB.jpg', 'PG-13', '2023-11-03', 125, 'Postwar Japan. From zero to minus.', 'Godzilla Minus One', 'movie', 'https://tickets.godzilla.com', 2023, NULL, 940721, 64000000, 15, '2024-04-30 03:39:39.440288'),
('tt6587046', '8 wins & 26 nominations', NULL, 7.7, 16, NULL, '君たちはどう生きるか', 'While the Second World War rages, the teenage Mahito, haunted by his mother''s tragic death, is relocated from Tokyo to the serene rural home of his new stepmother Natsuko, a woman who bears a striking resemblance to the boy''s mother. As he tries to adjust, this strange new world grows even stranger following the appearance of a persistent gray heron, who perplexes and bedevils Mahito, dubbing him the "long-awaited one."', 'https://image.tmdb.org/t/p/original/jDQPkgzerGophKRRn7MKm071vCU.jpg', 'PG-13', '2023-07-14', 124, 'Where death comes to an end, life finds a new beginning.', 'The Boy and the Heron', 'movie', 'https://gkids.com/films/the-boy-and-the-heron/', 2023, NULL, 508883, 109000000, 16, '2024-04-30 03:39:39.440288'),
('tt3915174', 'Nominated for 1 Oscar. 7 wins & 55 nominations total', NULL, 7.8, 163, 73, 'Puss in Boots: The Last Wish', 'Puss in Boots discovers that his passion for adventure has taken its toll: He has burned through eight of his nine lives, leaving him with only one life left. Puss sets out on an epic journey to find the mythical Last Wish and restore his nine lives.', 'https://image.tmdb.org/t/p/original/kuf6dutpsT0vSVehic3EZIqkOBt.jpg', 'PG', '2022-12-07', 103, 'Live each adventure like it''s your last.', 'Puss in Boots: The Last Wish', 'movie', 'https://www.dreamworks.com/movies/puss-in-boots-the-last-wish', 2022, '20th Century Fox', 315162, 484700000, 17, '2024-04-30 03:39:39.440288'),
('tt28814949', '3 nominations', NULL, 8.3, 13, NULL, 'TAYLOR SWIFT | THE ERAS TOUR', 'The cultural phenomenon continues on the big screen! Immerse yourself in this once-in-a-lifetime concert film experience with a breathtaking, cinematic view of the history-making tour.', 'https://image.tmdb.org/t/p/original/jf3YO8hOqGHCupsREf5qymYq1n.jpg', 'PG-13', '2023-10-13', 169, 'It’s been a long time coming.', 'TAYLOR SWIFT | THE ERAS TOUR', 'movie', 'https://www.tstheerastourfilm.com', 2023, NULL, 1160164, 24960000, 18, '2024-04-30 03:39:39.440288'),
('tt10676048', '2 wins & 4 nominations', NULL, 5.9, 64, NULL, 'The Marvels', 'Carol Danvers, aka Captain Marvel, has reclaimed her identity from the tyrannical Kree and taken revenge on the Supreme Intelligence. But unintended consequences see Carol shouldering the burden of a destabilized universe. When her duties send her to an anomalous wormhole linked to a Kree revolutionary, her powers become entangled with that of Jersey City super-fan Kamala Khan, aka Ms. Marvel, and Carol’s estranged niece, now S.A.B.E.R. astronaut Captain Monica Rambeau. Together, this unlikely trio must team up and learn to work in concert to save the universe.', 'https://image.tmdb.org/t/p/original/Ag3D9qXjhJ2FUkrlJ0Cv1pgxqYQ.jpg', 'PG-13', '2023-11-08', 105, 'Higher. Further. Faster. Together.', 'The Marvels', 'movie', 'https://www.marvel.com/movies/the-marvels', 2023, NULL, 609681, 199400000, 19, '2024-04-30 03:39:39.440288'),
('string', 'string', 'string', 0, 0, 0, 'string', 'string', 'string', 'string', 'string', 0, 'string', 'string', 'string', 'string', 0, 'string', 0, 0, 20, '2024-04-30 03:43:27.905888'),
('strisng', 'string', 'string', 0, 0, 0, 'string', 'string', 'string', 'string', 'string', 0, 'string', 'string', 'string', 'string', 0, 'string', 0, 0, 22, '2024-04-30 03:48:43.32857');

INSERT INTO "public"."shows" ("id", "date", "end_time", "start_time", "auditorium_id", "imdb_id", "created_at", "updated_at") VALUES
(128, '2023-12-29', '10:38:00', '08:00:00', 1, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991');
INSERT INTO "public"."shows" ("id", "date", "end_time", "start_time", "auditorium_id", "imdb_id", "created_at", "updated_at") VALUES
(129, '2023-12-30', '11:53:00', '09:15:00', 2, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991');
INSERT INTO "public"."shows" ("id", "date", "end_time", "start_time", "auditorium_id", "imdb_id", "created_at", "updated_at") VALUES
(130, '2023-12-31', '13:08:00', '10:30:00', 3, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991');
INSERT INTO "public"."shows" ("id", "date", "end_time", "start_time", "auditorium_id", "imdb_id", "created_at", "updated_at") VALUES
(131, '2024-01-01', '15:23:00', '12:45:00', 4, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(132, '2024-01-02', '16:38:00', '14:00:00', 5, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(133, '2024-01-03', '18:53:00', '16:15:00', 6, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(134, '2024-01-04', '21:08:00', '18:30:00', 7, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(135, '2024-01-05', '23:23:00', '20:45:00', 8, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(136, '2024-01-06', '22:38:00', '20:00:00', 10, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(137, '2023-12-30', '12:37:00', '10:00:00', 12, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(138, '2023-12-30', '15:37:00', '13:00:00', 11, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(139, '2023-12-30', '18:37:00', '16:00:00', 10, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(140, '2023-12-30', '21:37:00', '19:00:00', 9, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(141, '2023-12-30', '17:37:00', '15:00:00', 8, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(142, '2023-12-30', '10:37:00', '08:00:00', 8, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(143, '2023-12-30', '13:37:00', '11:00:00', 9, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(144, '2023-12-30', '16:37:00', '14:00:00', 10, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(145, '2023-12-30', '19:37:00', '17:00:00', 11, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(146, '2023-12-30', '22:37:00', '20:00:00', 12, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(147, '2024-03-15', '10:38:00', '08:00:00', 1, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(148, '2024-03-16', '11:53:00', '09:15:00', 2, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(149, '2024-03-17', '13:08:00', '10:30:00', 3, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(150, '2024-03-18', '15:23:00', '12:45:00', 4, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(151, '2024-03-19', '16:38:00', '14:00:00', 5, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(152, '2024-03-20', '18:53:00', '16:15:00', 6, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(153, '2024-03-21', '21:08:00', '18:30:00', 7, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(154, '2024-03-22', '23:23:00', '20:45:00', 8, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(155, '2024-03-23', '22:38:00', '20:00:00', 10, 'tt13287846', '2024-04-08 15:19:00.888991', '2024-04-08 15:19:00.888991'),
(1, '2024-05-01', '08:00:00', '09:00:00', 0, 'test', '2024-04-30 16:51:04.193578', '2024-04-30 16:51:04.193578'),
(156, '2024-04-30', '08:00:00', '09:00:00', 2, 'test', '2024-04-30 17:03:13.390019', '2024-04-30 17:03:13.390019');

INSERT INTO "public"."tickets" ("id", "seat_number", "status", "timestamp", "show_id", "user_id", "created_at", "updated_at") VALUES
(304, 2, 1, '0001-01-01 00:00:00', 156, 17, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984');
INSERT INTO "public"."tickets" ("id", "seat_number", "status", "timestamp", "show_id", "user_id", "created_at", "updated_at") VALUES
(303, 1, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984');
INSERT INTO "public"."tickets" ("id", "seat_number", "status", "timestamp", "show_id", "user_id", "created_at", "updated_at") VALUES
(305, 3, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984');
INSERT INTO "public"."tickets" ("id", "seat_number", "status", "timestamp", "show_id", "user_id", "created_at", "updated_at") VALUES
(306, 4, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(307, 5, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(308, 6, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(309, 7, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(310, 8, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(311, 9, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(312, 10, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(313, 11, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(314, 12, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(315, 13, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(316, 14, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(317, 15, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(318, 16, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(319, 17, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(320, 18, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(321, 19, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(322, 20, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(323, 21, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(324, 22, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(325, 23, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(326, 24, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(327, 25, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(328, 26, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(329, 27, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(330, 28, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(331, 29, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(332, 30, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(333, 31, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(334, 32, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(335, 33, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(336, 34, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(337, 35, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(338, 36, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(339, 37, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(340, 38, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(341, 39, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(342, 40, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(343, 41, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(344, 42, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(345, 43, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(346, 44, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(347, 45, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(348, 46, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(349, 47, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(350, 48, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(351, 49, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(352, 50, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(353, 51, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(354, 52, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(355, 53, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(356, 54, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(357, 55, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(358, 56, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(359, 57, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(360, 58, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(361, 59, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(362, 60, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(363, 61, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(364, 62, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(365, 63, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(366, 64, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(367, 65, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(368, 66, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(369, 67, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(370, 68, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(371, 69, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(372, 70, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(373, 71, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(374, 72, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(375, 73, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(376, 74, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(377, 75, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(378, 76, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(379, 77, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(380, 78, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(381, 79, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(382, 80, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(383, 81, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(384, 82, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(385, 83, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(386, 84, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(387, 85, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(388, 86, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(389, 87, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(390, 88, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(391, 89, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(392, 90, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(393, 91, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(394, 92, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(395, 93, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(396, 94, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(397, 95, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(398, 96, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(399, 97, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(400, 98, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(401, 99, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(402, 100, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(403, 1, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984'),
(404, 2, 1, '0001-01-01 00:00:00', 156, NULL, '2024-04-30 10:03:13.366984', '2024-04-30 10:03:13.366984');

INSERT INTO "public"."users" ("id", "name", "gender", "email", "password", "phone_number", "date_of_birth", "tier", "salt", "status", "created_at", "updated_at") VALUES
(1, 'Alice', 'Female', 'alice@example.com', 'password123', 1234567890, '2000-01-15', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948');
INSERT INTO "public"."users" ("id", "name", "gender", "email", "password", "phone_number", "date_of_birth", "tier", "salt", "status", "created_at", "updated_at") VALUES
(2, 'Bob', 'Male', 'bob@example.com', 'pass456', 9876543210, '1995-08-22', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948');
INSERT INTO "public"."users" ("id", "name", "gender", "email", "password", "phone_number", "date_of_birth", "tier", "salt", "status", "created_at", "updated_at") VALUES
(3, 'Charlie', 'Male', 'charlie@example.com', '123pass', 5555555555, '1998-05-10', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948');
INSERT INTO "public"."users" ("id", "name", "gender", "email", "password", "phone_number", "date_of_birth", "tier", "salt", "status", "created_at", "updated_at") VALUES
(4, 'Bob Johnson', 'Male', 'bob.johnson@example.com', 'pass456', 1112223333, '1982-07-08', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(5, 'Emily Davis', 'Female', 'emily.davis@example.com', 'letmein', 4443332222, '1998-03-20', 'VIP', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(6, 'Charlie Brown', 'Male', 'charlie.brown@example.com', 'peanuts', 7778889999, '1975-12-05', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(7, 'Eva Martinez', 'Female', 'eva.martinez@example.com', 'eva123', 6667778888, '1993-09-18', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(8, 'George White', 'Male', 'george.white@example.com', 'georgepass', 2221114444, '1980-06-12', 'VIP', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(9, 'Olivia Turner', 'Female', 'olivia.turner@example.com', 'olivia456', 3334445555, '1978-04-25', 'ADMIN', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(10, 'David Lee', 'Male', 'david.lee@example.com', 'davidpass', 8889990000, '1992-09-30', 'VIP', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(11, 'Sophia Kim', 'Female', 'sophia.kim@example.com', 'sophia123', 1239876543, '1985-11-08', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(12, 'Michael Smith', 'Male', 'michael.smith@example.com', 'mikepass', 9991112222, '1973-02-17', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(13, 'Grace Miller', 'Female', 'grace.miller@example.com', 'gracepass', 7778889999, '1990-08-03', 'ADMIN', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(14, 'Nguyen Van A', 'Male', 'nguyenvana@gmail.com', '12345678', 123456, '2003-10-07', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(15, 'Nguyen Van A', 'Male', 'nguyenvanaBC@gmail.com', '12345678', 123456, '2003-10-07', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(16, 'Nguyen Van A', 'Male', 'nguyenvabcd@gmail.com', '12345678', 123456, '2003-10-07', 'MEMBER', NULL, 1, '2024-04-30 01:31:27.287309', '2024-04-30 01:31:27.304948'),
(17, 'test', NULL, 'nn@nn.com', '834267424243b5959036ec49adca19ca', NULL, NULL, 'admin', 'sSWMTduZdctwwnmPpKebIJLCKGLqHrsfkIYYijtfuNdepUUaNL', 1, '2024-04-30 08:46:27.99509', '2024-04-30 01:46:27.995218');

