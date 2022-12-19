create table clients(
    id varchar(255) not null UNIQUE,
    email text not null,
    client_name text not null,
    redirect_url text not null,
    user_secret text not null,
    expires_at DATETIME,
    created_at DATETIME not null,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- テスト用シードデータ
insert into
    clients(
        id,
        client_name,
        email,
        redirect_url,
        user_secret,
        expires_at,
        created_at
    )
values
    (
        "example-client-id-1",
        "Sample App",
        "hoge@example.com",
        "http://localhost",
        "secret",
        "9999-12-31 23:59:59",
        CURRENT_TIMESTAMP
    );
