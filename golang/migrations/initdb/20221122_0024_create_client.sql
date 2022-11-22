create table clients(
    id varchar(255) UNIQUE,
    client_name text,
    email text not null,
    redirect_url text,
    user_secret text,
    expires_at datetime,
    created_at DATETIME not null,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);