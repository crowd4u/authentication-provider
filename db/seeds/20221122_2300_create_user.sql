-- テスト用ユーザーデータ
insert into
    users(
        id,
        user_name,
        email,
        hash_password,
        given_name,
        family_name,
        locale,
        created_at
    )
values
    (
        "example-user-id-1",
        "hoge",
        "hoge@example.com",
        "password",
        "test",
        "hoge",
        "JP",
        CURRENT_TIMESTAMP
    );