package mysql

import (
	mooc "api-http-hex-golang/internal"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CourseRepository_Save_RepositoryError(t *testing.T) {
	courseID := "006f1977-d2cc-44f0-8ef8-95a275cb91ea"
	courseName := "Test Course"
	courseDuration := "10 Months"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO courses (id, name, duration) VALUES (?,?,?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnError(errors.New("something failed"))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}
