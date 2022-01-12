package mooc

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

type Course struct {
	id       CourseID
	name     string
	duration string
}

var ErrInvalidCourseID = errors.New("invalid Course ID")

type CourseID struct {
	value string
}

func (id CourseID) String() string {
	return id.value
}

func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseID{
		value: v.String(),
	}, nil
}

func NewCourse(id, name, duration string) (Course, error) {
	idNc, err := NewCourseID(id)

	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idNc,
		name:     name,
		duration: duration,
	}, nil
}

func (c Course) ID() CourseID {
	return c.id
}

func (c Course) Name() string {
	return c.name
}

func (c Course) Duration() string {
	return c.duration
}
