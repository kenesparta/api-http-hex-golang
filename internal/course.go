package mooc

import "context"

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

type Course struct {
	id       string
	name     string
	duration string
}

func NewCourse(id, name, duration string) Course {
	return Course{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (c Course) ID() string {
	return c.id
}

func (c Course) Name() string {
	return c.name
}

func (c Course) Duration() string {
	return c.duration
}
