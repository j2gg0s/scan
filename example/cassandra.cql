CREATE KEYSPACE IF NOT EXISTS example
    WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 3};

USE example;

CREATE TABLE IF NOT EXISTS project_by_org (
    org TEXT,
    project TEXT,

    star INT,
    fork INT,

    PRIMARY KEY(org, project)
)
;

INSERT INTO project_by_org(org, project, star, fork) VALUES (
    'kubernetes', 'kubernetes', 82200, 30100);
INSERT INTO project_by_org(org, project, star, fork) VALUES (
    'kubernetes', 'client-go', 5300, 2100);
INSERT INTO project_by_org(org, project, star, fork) VALUES (
    'kubernetes', 'example', 4500, 3400);

INSERT INTO project_by_org(org, project, star, fork) VALUES (
    'uptrace', 'bun', 374, 22);

INSERT INTO project_by_org(org, project, star, fork) VALUES (
    'open-telemetry', 'opentelemetry-go-contrib', 267, 181);
INSERT INTO project_by_org(org, project, star, fork) VALUES (
    'open-telemetry', 'opentelemetry-go-', 2008, 439);
