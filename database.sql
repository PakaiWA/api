create schema if not exists "management";
alter schema "management" owner to pakaiwa;
grant create, usage on schema "management" to kanggara;

drop table if exists pakaiwa.management.users;
create table if not exists pakaiwa.management.users
(
    uuid       varchar(100) primary key,
    email      varchar(100) unique not null,
    password   varchar(100)        not null,
    created_at timestamp default LOCALTIMESTAMP
);
grant all on pakaiwa.management.users to kanggara;

-- =======================================================================
drop table if exists pakaiwa.management.user_devices;
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
    foreign key (user_uuid) references pakaiwa.management.users(uuid) on delete cascade
);
create index if not exists user_devices_user_idx on pakaiwa.management.user_devices (user_uuid);
grant all on pakaiwa.management.user_devices to kanggara;

