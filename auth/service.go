package auth

import "github.com/dgrijalva/jwt-go"

type Servis interface {
	GenerateToken(userID int)( string , error)
}
type jwtService struct {
}
func NewJwtService() *jwtService {
	return &jwtService{}
}
var SECRET = []byte("t4mmysury4n4")
func (s *jwtService)GenerateToken (userID int) (string, error){
claim := jwt.MapClaims{}
claim ["user_id"] = userID

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedSTOKEN, err := token.SignedString(SECRET)
	if err != nil {
		return signedSTOKEN, err
	}
	return signedSTOKEN, err
}

