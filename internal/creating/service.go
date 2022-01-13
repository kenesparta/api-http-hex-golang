package creating

import (
	mooc "api-http-hex-golang/internal"
	"context"
)

type CourseService struct {
	courseRepository mooc.CourseRepository
}

func NewCourseService(cr mooc.CourseRepository) CourseService {
	return CourseService{courseRepository: cr}
}

func (cs CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}
	return cs.courseRepository.Save(ctx, course)
}
