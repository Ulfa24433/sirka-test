DROP TABLE IF EXISTS users;
CREATE TABLE "users" (
    "userid" VARCHAR (50),
    "name" VARCHAR(50)
);

INSERT INTO users (userid, name) 
VALUES ('01', 'Budi'), ('02', 'Nano');