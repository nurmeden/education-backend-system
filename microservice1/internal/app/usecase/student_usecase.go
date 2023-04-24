package usecase

import (
	"context"
	"task-backend/microservice1/internal/app/model"
	"task-backend/microservice1/internal/app/repository"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
	// Реализация логики создания студента
	return u.studentRepo.Create(context.Background(), student)
}

func (u *studentUsecase) GetStudentByID(id string) (*model.Student, error) {
	// Реализация логики получения студента по ID
	return u.studentRepo.Read(context.Background(), id)
}

func (u *studentUsecase) UpdateStudent(student *model.Student) (*model.Student, error) {
	// Реализация логики обновления информации о студенте
	return u.studentRepo.Update(context.Background(), student)
}

func (u *studentUsecase) DeleteStudent(id string) error {
	// Реализация логики удаления студента
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
