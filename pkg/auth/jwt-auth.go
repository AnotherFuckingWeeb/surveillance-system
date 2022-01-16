package auth

import (
	"crypto/rsa"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {
	privateRsa, err := ioutil.ReadFile("./private.rsa")

	if err != nil {
		log.Fatal("Couldn't read file private.rsa")
	}

	publicRsa, err := ioutil.ReadFile("./public.rsa.pub")

	if err != nil {
		log.Fatal("Couldn't read file public.rsa.pub")
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateRsa)

	if err != nil {
		log.Fatal("Could not parse private key")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicRsa)

	if err != nil {
		log.Fatal("Could not parse public key")
	}
}

type UserClaims struct {
	ID       int    `json:"id"`
	Role     int    `json:"role"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	jwt.StandardClaims
}

func GenerateJWT(userClaims *UserClaims) string {
	claims := UserClaims{
		ID:             userClaims.ID,
		Role:           userClaims.Role,
		Name:           userClaims.Name,
		Lastname:       userClaims.Lastname,
		StandardClaims: userClaims.StandardClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(PrivateKey)

	if err != nil {
		log.Fatal("Something went wrong signing the token")
	}

	return result
}

func ValidateJWT(c *gin.Context) (*jwt.Token, error) {
	token, err := request.ParseFromRequestWithClaims(c.Request, request.OAuth2Extractor, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

/*func ValidateJWT(c *gin.Context) bool {
	token, err := request.ParseFromRequestWithClaims(c.Request, request.OAuth2Extractor, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				log.Fatal("token expired")
			case jwt.ValidationErrorSignatureInvalid:
				log.Fatal("bad signature")
			default:
				log.Fatal("invalid token")
			}
		default:
			log.Fatal("invalid token")
		}

		return false
	}

	return token.Valid
}*/
