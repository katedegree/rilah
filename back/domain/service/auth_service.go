package service

import (
	"back/domain/repository"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

// トークン検証
func (s *authService) ValidateToken(token string, tokenRepository repository.AccessTokenRepository) bool {
	tokenEntity, err := tokenRepository.FindByToken(token)
	if err != nil {
		return false
	}
	if tokenEntity == nil {
		return false
	}
	if tokenEntity.ExpiresAt.Before(time.Now()) {
		return false
	}
	return true
}

// ログイン時のパスワード検証
func (s *authService) ValidatePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

// アカウントコードの規則検証
// TODO: user_service.go に移動
func (s *authService) IsValidAccountCode(code string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return re.MatchString(code)
}

// アカウントコードの重複検証
func (s *authService) IsAccountCodeDuplicate(accountCode string, oldAccountCode *string, userRepository repository.UserRepository) bool {
	if oldAccountCode == nil || accountCode == *oldAccountCode {
		return false
	}

	user, err := userRepository.FindByAccountCode(accountCode)
	if err != nil {
		return false
	}

	return user != nil
}
