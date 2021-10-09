package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type Servis interface {
	GenerateToken(userID int)( string , error)
	ValidateToken(token string) (*jwt.Token,error)
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

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	toke, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil,errors.New("Invalid token")
		}
		return []byte(SECRET),nil
	})
	if err != nil {
		return toke, err
	}
	return toke,nil
}
