PRAGMA journal_mode = WAL;
PRAGMA synchronous = NORMAL;

INSERT INTO movies VALUES ((SELECT MAX(id) FROM movies)+1, (((SELECT MAX(id) FROM movies)+1) || ' -> Spider-Man: Homecoming - Generated by robots'), 2017);
INSERT INTO movie_genre VALUES ((SELECT max(id) FROM movie_genre)+1, (SELECT MAX(id) FROM movies), 17);
INSERT INTO movie_ratings VALUES ((SELECT max(id) FROM movie_ratings)+1, 1, (SELECT max(id) FROM movies), 9, 1412178746);