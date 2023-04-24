package usecase

import (
	"task-backend/microservice2/internal/app/model"
	"task-backend/microservice2/internal/app/repository"
)

type CourseUsecase interface {
	GetCourses() ([]model.Course, error)
	GetCourseByID(id string) (*model.Course, error)
	CreateCourse(courseInput *model.CourseInput) (*model.Course, error)
	UpdateCourse(id string, courseUpdateInput *model.CourseUpdateInput) (*model.Course, error)
	DeleteCourse(id string) error
	AddStudentToCourse(courseID string, studentID string) (*model.Course, error)
	RemoveStudentFromCourse(courseID string, studentID string) (*model.Course, error)
}

// CourseService представляет сервис для работы с курсами
type courseUsecase struct {
	courseRepo *repository.CourseRepository
}

// NewCourseService создает новый экземпляр сервиса курсов
func NewCourseService(courseRepo *repository.CourseRepository) *courseUsecase {
	return &courseUsecase{
		courseRepo: courseRepo,
	}
}

// CreateCourse создает новый курс на основе входных параметров
func (cs *courseUsecase) CreateCourse(courseInput *model.CourseInput) (*model.Course, error) {
	course, err := model.NewCourse(courseInput)
	if err != nil {
		return nil, err
	}
	_, err = cs.courseRepo.CreateCourse(course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

// UpdateCourse обновляет информацию о курсе на основе входных параметров
func (cs *courseUsecase) UpdateCourse(courseID string, courseInput *model.CourseUpdateInput) (*model.Course, error) {
	course, err := cs.courseRepo.GetCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	err = course.Update(courseInput)
	if err != nil {
		return nil, err
	}

	_, err = cs.courseRepo.UpdateCourse(course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

// DeleteCourse удаляет курс по его идентификатору
func (cs *courseUsecase) DeleteCourse(courseID string) error {
	err := cs.courseRepo.DeleteCourse(courseID)
	if err != nil {
		return err
	}

	return nil
}

// GetCourseByID возвращает курс по его идентификатору
func (cs *courseUsecase) GetCourseByID(courseID string) (*model.Course, error) {
	course, err := cs.courseRepo.GetCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	return course, nil
}

// GetAllCourses возвращает все курсы
func (cs *courseUsecase) GetAllCourses() ([]*model.Course, error) {
	courses, err := cs.courseRepo.GetAllCourses()
	if err != nil {
		return nil, err
	}

	return courses, nil
}
