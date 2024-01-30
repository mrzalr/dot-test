package post

import "errors"

var ErrBadRequest = errors.New("post: invalid payload")
