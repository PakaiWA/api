create schema if not exists "device";

alter schema "device" owner to pakaiwa;

grant create, usage on schema "device" to kanggara;

drop table pakaiwa.device.user_devices;

create table if not exists pakaiwa.device.user_devices
(
    uuid                varchar(200) not null,
    name                varchar(100) not null,
    status              varchar(25) default 'disconnected',
    phone_number        varchar(25) default '',
    created_at          timestamp   default LOCALTIMESTAMP,
    connected_at        timestamp,
    disconnected_at     timestamp,
    disconnected_reason varchar(100)
);

create unique index if not exists user_device_pk on pakaiwa.device.user_devices (uuid);

grant all on pakaiwa.device.user_devices to kanggara;


create schema if not exists "management";

alter schema "management" owner to pakaiwa;

grant create, usage on schema "management" to kanggara;

drop table pakaiwa.management.users;

create table if not exists pakaiwa.management.users
(
    uuid       varchar(200) not null,
    email      varchar(100) not null,
    password   varchar(100) not null,
    created_at timestamp default LOCALTIMESTAMP
);

create unique index if not exists users_pk on pakaiwa.management.users (uuid);
create unique index if not exists user_email_idx on pakaiwa.management.users (email);

grant all on pakaiwa.management.users to kanggara;