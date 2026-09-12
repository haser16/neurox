package auth

type TokenService interface {
	Generate(userID int64) (string, error)
}
