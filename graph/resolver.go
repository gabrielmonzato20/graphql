package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import "github.com/gabrielmonzato20/graphql/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
