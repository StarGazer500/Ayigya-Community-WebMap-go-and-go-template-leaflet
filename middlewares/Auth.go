package middlewares

import (
	"fmt"
	"os"
	"time"

	// "github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v5"
)

func CheckTokenExpiration(claims jwt.MapClaims) error {
	// Extract the expiration time (exp) from the claims
	if exp, ok := claims["exp"].(float64); ok {
		// Convert exp to Unix timestamp
		expirationTime := time.Unix(int64(exp), 0)

		// Check if the token has expired
		if time.Now().After(expirationTime) {
			return fmt.Errorf("token has expired at %s", expirationTime)
		}

		// Token is valid and not expired
		fmt.Printf("Token is valid. Expiration time: %s\n", expirationTime)
		return nil
	} else {
		return fmt.Errorf("missing expiration claim in token")
	}
}

// VerifyAccessToken verifies the access token and checks expiration
func VerifyAccessToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing access token: %w", err)
	}

	// If the token is valid, check expiration manually
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the token has expired
		if err := CheckTokenExpiration(claims); err != nil {
			fmt.Println("Access Expiry has happened")
			return nil, err
		}

		// Token is valid and not expired
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid access token")
	}
}

// VerifyRefreshToken verifies the refresh token and checks expiration
func VerifyRefreshToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing refresh token: %w", err)
	}

	// If the token is valid, check expiration manually
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the token has expired
		if err := CheckTokenExpiration(claims); err != nil {
			fmt.Println("Refresh Expiry has happened")
			return nil, err
		}

		// Token is valid and not expired
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid refresh token")
	}
}

// LoadKeyAndReturnByte loads the secret key from the environment
func LoadKeyAndReturnByte() ([]byte, error) {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		return nil, fmt.Errorf("missing JWT secret key in the environment")
	}
	return []byte(key), nil
}

// TokenClaimStruct defines the structure of the token claims

// GenerateExpiryTime generates the Unix timestamp for the expiration time based on the duration (in seconds).
// You can pass a duration in seconds, days, or hours.
func GenerateExpiryTime(durationInSeconds int64) int64 {
	// Add the given duration to the current time and return the Unix timestamp
	expirationTime := time.Now().Add(time.Duration(durationInSeconds) * time.Second)
	return expirationTime.Unix()
}

type TokenClaimStruct struct {
	MyAuthServer    string
	AuthUserEmail   string
	AuthUserSurname string
	AuthUserId      string
	AuthExp         int64 // Expiration time as Unix timestamp (int64)
}

var Claim TokenClaimStruct

// GenerateAccessToken generates an access JWT token using HS256 (HMAC with SHA-256)
func GenerateAccessToken(claim TokenClaimStruct) (string, error) {
	// Load the secret key from the environment
	key, err := LoadKeyAndReturnByte()
	if err != nil {
		return "", fmt.Errorf("error loading the secret key: %w", err)
	}
	accessTokenExpiry := GenerateExpiryTime(120)
	// Define claims for the access token
	claims := jwt.MapClaims{
		"iss":     claim.MyAuthServer,    // Issuer
		"sub":     claim.AuthUserEmail,   // Subject (user email)
		"surname": claim.AuthUserSurname, // User surname
		"id":      claim.AuthUserId,      // User ID
		"exp":     accessTokenExpiry,         // Expiration time (Unix timestamp)
	}

	// Create the token with signing method HS256 (HMAC with SHA-256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token and get the signed string
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing the access token: %w", err)
	}

	return tokenString, nil
}

// GenerateRefreshToken generates a refresh JWT token using HS256 (HMAC with SHA-256)
func GenerateRefreshToken(claim TokenClaimStruct) (string, error) {
	// Load the secret key from the environment
	key, err := LoadKeyAndReturnByte()
	if err != nil {
		return "", fmt.Errorf("error loading the secret key: %w", err)
	}

	// Calculate the expiration time for the refresh token (30 days validity)
	refreshTokenExpiry := GenerateExpiryTime(86400 * 30) // 30 days * 24 hours * 60 minutes * 60 seconds

	// Define claims for the refresh token (longer expiration time)
	claims := jwt.MapClaims{
		"iss":     claim.MyAuthServer,    // Issuer
		"sub":     claim.AuthUserEmail,   // Subject (user email)
		"surname": claim.AuthUserSurname, // User surname
		"id":      claim.AuthUserId,      // User ID
		"exp":     refreshTokenExpiry,    // Expiration time (Unix timestamp)
	}

	// Create the token with signing method HS256 (HMAC with SHA-256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token and get the signed string
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing the refresh token: %w", err)
	}

	return tokenString, nil
}
