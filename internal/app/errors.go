package app

import "errors"

var errStorageAlreadySet = errors.New("application instance already has s3 storage associated")
var errServerAlreadySet = errors.New("server already associated")
