create table clients(
    id varchar(255) not null UNIQUE,
    email text not null default '',
    client_name text not null default '',
    redirect_url text not null default '',
    user_secret text not null default '',
    expires_at DATETIME,
    created_at DATETIME not null,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);