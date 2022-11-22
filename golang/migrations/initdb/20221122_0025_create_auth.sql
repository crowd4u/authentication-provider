create table authentications(
    id varchar(255) UNIQUE,
    client_id text not null,
    hash_token text not null,
    token_type text not null,
    expires_at datetime not null default CURRENT_TIMESTAMP,
    created_at DATETIME not null default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);