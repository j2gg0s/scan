CREATE DATABASE IF NOT EXISTS example;

USE example;

CREATE TABLE IF NOT EXISTS project (
    org VARCHAR(100) NOT NULL,
    project VARCHAR(100) NOT NULL,

    star INT NOT NULL,
    fork INT NOT NULL,

    PRIMARY KEY(org, project)
)
;

INSERT INTO project(org, project, star, fork) VALUES (
    'kubernetes', 'kubernetes', 82200, 30100);
INSERT INTO project(org, project, star, fork) VALUES (
    'kubernetes', 'client-go', 5300, 2100);
INSERT INTO project(org, project, star, fork) VALUES (
    'kubernetes', 'example', 4500, 3400);

INSERT INTO project(org, project, star, fork) VALUES (
    'uptrace', 'bun', 374, 22);

INSERT INTO project(org, project, star, fork) VALUES (
    'open-telemetry', 'opentelemetry-go-contrib', 267, 181);
INSERT INTO project(org, project, star, fork) VALUES (
    'open-telemetry', 'opentelemetry-go-', 2008, 439);
