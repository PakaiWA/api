create schema if not exists "device";

alter schema "device" owner to pakaiwa;

grant create, usage on schema "device" to kanggara;

drop table pakaiwa.device.user_devices;

create table if not exists pakaiwa.device.user_devices
(
    uuid                varchar(200) not null,
    name                varchar(100),
    status              varchar(25),
    phone_number        varchar(25),
    created_at          timestamp default LOCALTIMESTAMP,
    connected_at        timestamp,
    disconnected_at     timestamp,
    disconnected_reason varchar(100)
);

create unique index if not exists user_device_pk on pakaiwa.device.user_devices (uuid);
