package database

import (
	"log"
	"notchman8600/authentication-provider/domain"
)

type OAuthRepository struct {
	DBHandler
}

func (repository *OAuthRepository) Store(client domain.Client) (err error) {
	// TODO これってちゃんとPrepared Statementになってるの？
	statement := `insert into clients (client_id, email, name, secret, expires_at) values(?,?,?,?,?)`
	_, err = repository.Execute(statement, client.Id, client.Email, client.Name, client.Secret, client.ExpiresAt)
	return err
}

func (repo *OAuthRepository) FindByClientId(clientId string) (client domain.Client, err error) {
	statement := "select id, email, client_name, user_secret, expires_at from clients where id=? order by created_at desc limit 1"
	rows, err := repo.Query(statement, clientId)
	if err != nil {
		log.Println(err)
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
