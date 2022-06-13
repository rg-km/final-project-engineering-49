package helper

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GetFileName(id int, originalName string) string {
	idString := strconv.Itoa(id)
	microTime := strconv.Itoa(int(time.Now().UnixNano() / int64(time.Millisecond)))
	count := strings.Count(originalName, "")
	limit := math.Min(float64(count), 10)
	return string(originalName[:int(limit-1)]) + microTime + "-" + idString
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func StringWithCharset(length int, charset string) string {
  var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  charset := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  return StringWithCharset(length, charset)
}
