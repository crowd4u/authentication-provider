-- テスト用クライアント挿入データ
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
        "example-user-id-1",
        "hoge",
        "hoge@example.com",
        "http://localhost",
        "secret",
        "9999-12-31 23:59:59",
        CURRENT_TIMESTAMP
    );
