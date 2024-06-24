//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../test/mock/mock_$GOPACKAGE/$GOFILE
package interfaces

import (
	"myapp/internal/entities"
)

type PostRepository interface {
	GetAll() ([]*entities.Post, error)
	GetById(postId string) (*entities.Post, error)
	Delete(postId string) error
}
