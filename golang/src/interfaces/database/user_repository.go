package database

import "notchman8600/authentication-provider/domain"

type UserRepository struct {
	DBHandler
}

// func (repository *OAuthRepository) Store(client domain.Client) (err error) {
// 	// TODO これってちゃんとPrepared Statementになってるの？
// 	statement := `insert into clients (client_id, email, name, secret, expires_at) values($1,$2,$3,$4,$5)`
// 	_, err = repository.Execute(statement, client.Id, client.Email, client.Name, client.Secret, client.ExpiresAt)
// 	return err
// }

func (repo *OAuthRepository) FindByUserId(userId string) (client domain.Client, err error) {
	rows, err := repo.Query("select (client_id, email, name, secret, expires_at) from clients where client_id=$1 order by created_at desc limit 1", clientId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var email string
		var secret string
		var expiresAt int64

		if err = rows.Scan(&id, &email, &name, &secret, &expiresAt); err != nil {
			return
		}
		client.Id = id
		client.Name = name
		client.Email = email
		client.Secret = secret
		client.ExpiresAt = expiresAt
	}
	return client, err
}
