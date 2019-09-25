CREATE DATABASE blog;
USE blog;

DROP TABLE IF EXISTS PostItems;
CREATE TABLE PostItems (
                        Id SERIAL PRIMARY KEY,
                        Title VARCHAR ( 255 ),
                        Text TEXT
);

DROP TABLE IF EXISTS Labels;
CREATE TABLE Labels (
                     Id SERIAL PRIMARY KEY,
                     Name VARCHAR ( 255 )
);

DROP TABLE IF EXISTS Post_Label;
CREATE TABLE Post_Label (
                               Id SERIAL PRIMARY KEY,
                               OrderId INT UNSIGNED,
                               ProductId INT UNSIGNED
);

INSERT INTO PostItems VALUES (DEFAULT, 'Post 1', '1 Post Text');
INSERT INTO PostItems VALUES (DEFAULT, 'Post 2', '3 Post Text');
INSERT INTO PostItems VALUES (DEFAULT, 'Post 3', '3 Post Text');

INSERT INTO Labels VALUES (DEFAULT, 'Label 1');
INSERT INTO Labels VALUES (DEFAULT, 'Label 2');
INSERT INTO Labels VALUES (DEFAULT, 'Label 3');
INSERT INTO Labels VALUES (DEFAULT, 'Label 4');


INSERT INTO Post_Label VALUES (DEFAULT, 1, 1);
INSERT INTO Post_Label VALUES (DEFAULT, 1, 2);
INSERT INTO Post_Label VALUES (DEFAULT, 2, 3);
INSERT INTO Post_Label VALUES (DEFAULT, 2, 4);
