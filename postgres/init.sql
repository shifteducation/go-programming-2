create table if not exists users
(
    id uuid primary key default gen_random_uuid(),
    first_Name text not null,
    last_Name text not null,
    age smallint not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone
);

create table if not exists addresses
(
    id uuid primary key default gen_random_uuid(),
    city text not null,
    street text not null,
    building text not null,
    apartment text not null,
    user_id uuid references users(id),
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone
);
