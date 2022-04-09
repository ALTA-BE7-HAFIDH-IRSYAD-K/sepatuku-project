package response

import "golang.org/x/crypto/bcrypt"

func ResponseSuccess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func ResponseUser(message string, data interface{}, product interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":       "success",
		"message":      message,
		"data_user":    data,
		"data_product": product,
	}
}

func ResponseSuccessWithoutData(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
	}
}

func ResponseFailed(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
		"message": message,
	}
}

func HashPassword(password string) (string, int, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), 0, err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
