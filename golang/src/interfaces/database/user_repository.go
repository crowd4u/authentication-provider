package database

import "notchman8600/authentication-provider/domain"

type UserRepository struct {
	DBHandler
}

// func (repository *OAuthRepository) Store(user domain.user) (err error) {
// 	// TODO これってちゃんとPrepared Statementになってるの？
// 	statement := `insert into users (user_id, email, name, secret, expires_at) values($1,$2,$3,$4,$5)`
// 	_, err = repository.Execute(statement, user.Id, user.Email, user.Name, user.Secret, user.ExpiresAt)
// 	return err
// }

func (repo *OAuthRepository) FindByUserId(userId string) (user domain.User, err error) {
	rows, err := repo.Query("select (id, email, user_name, given_name, family_name,sub,locale) from users where user_id=$1 order by created_at desc limit 1", userId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var user_name string
		var given_name string
		var family_name string
		var sub string
		var locale string

		if err = rows.Scan(&id, &user_name, &given_name, &family_name, &sub, &locale); err != nil {
			return
		}
		user.Id = id
		user.Name = user_name
		user.FamilyName = family_name
		user.GivenName = given_name
		user.Sub = sub
		user.Locale = locale
		// hash passwordは返却しない
		user.Password = ""
	}
	return user, err
}
