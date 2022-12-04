create table users(
    id varchar(255) UNIQUE,
    user_name text,
    email text not null,
    hash_password text not null,
    sub text,
    given_name text,
    family_name text,
    locale text,
    expires_at datetime,
    created_at DATETIME not null,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

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