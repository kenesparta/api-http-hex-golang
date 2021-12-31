package mysql

import (
	mooc "api-http-hex-golang/internal"
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

const (
	sqlCourseTable = "courses"
)

type sqlCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}

// CourseRepository is a MySQL mooc.CourseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID(),
		Name:     course.Name(),
		Duration: course.Duration(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}