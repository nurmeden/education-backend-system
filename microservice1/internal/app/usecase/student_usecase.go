package usecase

import (
	"context"
	"fmt"
	"task-backend/microservice1/internal/app/model"
	"task-backend/microservice1/internal/app/repository"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type StudentUsecase interface {
	CreateStudent(student *model.Student) (*model.Student, error)
	GetStudentByID(id string) (*model.Student, error)
	UpdateStudent(student *model.Student) (*model.Student, error)
	DeleteStudent(id string) error
	SignIn(signInData *model.SignInData) (*model.AuthToken, error)
}

type studentUsecase struct {
	studentRepo repository.StudentRepository
}

func NewStudentUsecase(studentRepo repository.StudentRepository) StudentUsecase {
	return &studentUsecase{
		studentRepo: studentRepo,
	}
}

func (u *studentUsecase) CreateStudent(student *model.Student) (*model.Student, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	fmt.Printf("hashedPassword: %v\n", hashedPassword)

	student.Password = string(hashedPassword)

	return u.studentRepo.Create(context.Background(), student)
}

func (u *studentUsecase) GetStudentByID(id string) (*model.Student, error) {
	return u.studentRepo.Read(context.Background(), id)
}

func (u *studentUsecase) UpdateStudent(student *model.Student) (*model.Student, error) {
	return u.studentRepo.Update(context.Background(), student)
}

func (u *studentUsecase) DeleteStudent(id string) error {
	return u.studentRepo.Delete(context.Background(), id)
}

func (u *studentUsecase) SignIn(signInData *model.SignInData) (*model.AuthToken, error) {
	// Проверяем данные аутентификации, например, сравниваем с данными в базе данных
	// и проводим другие необходимые проверки
	// Если аутентификация успешна, генерируем токен аутентификации

	// Пример генерации токена JWT с использованием библиотеки github.com/dgrijalva/jwt-go
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = signInData.UserID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()              // Устанавливаем время жизни токена
	tokenString, err := token.SignedString([]byte("your-secret-key")) // Здесь указываем ваш секретный ключ
	if err != nil {
		return nil, err
	}

	// Возвращаем сгенерированный токен аутентификации
	authToken := &model.AuthToken{
		UserID:    signInData.UserID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(time.Hour * 1), // Устанавливаем время истечения токена
	}
	return authToken, nil
}
