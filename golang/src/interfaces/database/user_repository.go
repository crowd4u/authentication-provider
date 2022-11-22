package database

import "notchman8600/authentication-provider/domain"

type UserRepository struct {
	DBHandler
}

func (repository *UserRepository) Store(user domain.User) (err error) {
	// TODO これってちゃんとPrepared Statementになってるの？
	statement := `insert into users (id,email,user_name,given_name,family_name,sub,locale) values($1,$2,$3,$4,$5,$6,$7)`
	_, err = repository.Execute(statement, user.Id, user.Email, user.Name, user.GivenName, user.FamilyName, user.Sub, user.Locale)
	return err
}

func (repo *UserRepository) FindByUserId(userId string) (user domain.User, err error) {
	rows, err := repo.Query("select (id,email,user_name,given_name,family_name,sub,locale) from users where user_id=$1 order by created_at desc limit 1", userId)
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
