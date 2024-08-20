CREATE DATABASE atm;
USE atm;

CREATE TABLE nodes
(
    id          INTEGER                   NOT NULL AUTO_INCREMENT,
    type        VARCHAR(25) DEFAULT 'atm' NOT NULL,
    title       VARCHAR(255),
    description TEXT,
    lat         FLOAT(8),
    `long`      FLOAT(8),
    CONSTRAINT machines_pk PRIMARY KEY (id)
);

CREATE TABLE discussion
(
    id      INTEGER                            NOT NULL AUTO_INCREMENT,
    nid     INTEGER                            NOT NULL,
    created DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user    VARCHAR(50),
    content TEXT,
    CONSTRAINT discussion_pk PRIMARY KEY (id),
    CONSTRAINT discussion_machines_id_fk FOREIGN KEY (nid) REFERENCES nodes(id)
);

