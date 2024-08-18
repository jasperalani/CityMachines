CREATE TABLE machines
(
    id          INTEGER
        CONSTRAINT machines_pk NOT NULL
        PRIMARY KEY autoincrement,
    type        VARCHAR(25) DEFAULT 'atm' NOT NULL,
    title       VARCHAR(255),
    description TEXT,
    lat         FLOAT(8),
    long        FLOAT(8)
);

CREATE TABLE discussion
(
    id      INTEGER
        CONSTRAINT discussion_pk NOT NULL
        PRIMARY KEY autoincrement,
    mid     INTEGER                            NOT NULL
        CONSTRAINT discussion_machines_id_fk
            REFERENCES machines,
    created DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user    VARCHAR(50),
    content TEXT
);

