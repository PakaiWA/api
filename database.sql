drop table if exists pakaiwa.management.user_devices;
drop table if exists pakaiwa.management.users;
-- =======================================================================
create table if not exists pakaiwa.management.users
(
    uuid       varchar(100) primary key,
    email      varchar(100) unique not null,
    password   varchar(100)        not null,
    created_at timestamp default LOCALTIMESTAMP
);
-- =======================================================================
create table if not exists pakaiwa.management.user_devices
(
    uuid                varchar(100) primary key,
    user_uuid           varchar(100) not null,
    name                varchar(100) not null,
    status              varchar(20) default 'disconnected',
    phone_number        varchar(20) default '',
    created_at          timestamp   default LOCALTIMESTAMP,
    connected_at        timestamp,
    disconnected_at     timestamp,
    disconnected_reason varchar(100),
    foreign key (user_uuid) references pakaiwa.management.users (uuid) on delete cascade
);
CREATE INDEX IF NOT EXISTS user_devices_user_idx on pakaiwa.management.user_devices (user_uuid);
CREATE UNIQUE INDEX IF NOT EXISTS user_device_unique_name_per_user ON pakaiwa.management.user_devices (user_uuid, name);
-- =======================================================================
EXPLAIN ANALYZE
SELECT uuid
FROM pakaiwa.management.users
WHERE email = 'Jerald_Pfannerstill@yahoo.com';

SELECT u.uuid AS user_uuid,
       u.email,
       d.uuid AS device_uuid,
       d.name,
       d.status
FROM pakaiwa.management.users u
         JOIN pakaiwa.management.user_devices d ON d.user_uuid = u.uuid
WHERE u.email = 'Jerald_Pfannerstill@yahoo.com';

INSERT INTO management.user_devices (uuid, user_uuid, name)
VALUES ('random', (SELECT uuid
                   FROM pakaiwa.management.users
                   WHERE email = 'Jerald_Pfannerstill@yahoo.com'), 'Test')
RETURNING name, status, created_at;

select *
from pakaiwa.management.user_devices;


SELECT name, status, phone_number, created_at, connected_at, disconnected_at, disconnected_reason
FROM management.user_devices
WHERE name = $1
  AND user_uuid = (SELECT uuid FROM pakaiwa.management.users WHERE email = $2)