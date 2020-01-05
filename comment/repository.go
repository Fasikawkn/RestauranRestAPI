package comment

import "github.com/fasikawkn/sample-restaurant-rest-api-master/entity"

// CommentRepository specifies customer comment related database operations
type CommentRepository interface {
	Comments() ([]entity.Comment, []error)
	Comment(id uint) (*entity.Comment, []error)
	UpdateComment(comment *entity.Comment) (*entity.Comment, []error)
	DeleteComment(id uint) (*entity.Comment, []error)
	StoreComment(comment *entity.Comment) (*entity.Comment, []error)
}

// func (c CommentRepository) DeleteRole(id uint) (*entity.Role, []error) {
// 	panic("implement me")
// }

// func (c CommentRepository) Role(id uint) (*entity.Role, []error) {
// 	panic("implement me")
// }

// func (c CommentRepository) Roles() ([]entity.Role, []error) {
// 	panic("implement me")
// }

// func (c CommentRepository) StoreRole(role *entity.Role) (*entity.Role, []error) {
// 	panic("implement me")
// }

// func (c CommentRepository) UpdateRole(role *entity.Role) (*entity.Role, []error) {
// 	panic("implement me")
// }
