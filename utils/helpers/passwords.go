package helpers

import "golang.org/x/crypto/bcrypt"

// Hash User Passwords
func Hash(password string) string {
	pwd := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		HandleErrors(err)
	}

	return string(hash)
}

// CompareHash is to compare the stored passwords with plain passwords
func CompareHash(hashedPwd string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))

	// if error is null it means the password is right otherwise you know what is means :)
	return err == nil
}
