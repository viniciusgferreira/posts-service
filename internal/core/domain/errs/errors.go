package errs

import "net/http"

var (
	// Internal errors
	InternalServerError = New(http.StatusInternalServerError, "An unexpected error happened. Please contact support.", InternalType)
	DatabaseError       = New(http.StatusInternalServerError, "An unexpected database error happened.", InternalType)

	// Validation errors
	RequestBinding   = New(http.StatusBadRequest, "The request body is malformed. Please refer to the API documentation.", ValidationType)
	PostNotFound     = New(http.StatusNotFound, "Post not found.", ValidationType)
	AuthorNotFound   = New(http.StatusNotFound, "Author not found.", ValidationType)
	DuplicateError   = New(http.StatusConflict, "Duplicate key value violates unique constraint.", ValidationType)
	TitleEmpty       = New(http.StatusBadRequest, "Title cannot be empty.", ValidationType)
	TitleTooLong     = New(http.StatusBadRequest, "Title cannot exceed 200 characters.", ValidationType)
	EmailEmpty       = New(http.StatusBadRequest, "Email cannot be empty.", ValidationType)
	EmailInvalid     = New(http.StatusBadRequest, "Invalid email format.", ValidationType)
	ContentEmpty     = New(http.StatusBadRequest, "Markdown content cannot be empty.", ValidationType)
	ContentTooShort  = New(http.StatusBadRequest, "Markdown content must be at least 10 characters.", ValidationType)
	ContentTooLong   = New(http.StatusBadRequest, "Markdown content cannot exceed 50000 characters.", ValidationType)
	CoverURLInvalid  = New(http.StatusBadRequest, "Invalid cover image URL provided.", ValidationType)
	CoverURLNotImage = New(http.StatusBadRequest, "URL must point to an image.", ValidationType)
	MissingAuthor    = New(http.StatusBadRequest, "Post must have an author.", ValidationType)

	// Permission errors
	ErrUnauthorized = New(http.StatusUnauthorized, "The provided token is not valid.", PermissionType)
	ErrForbidden    = New(http.StatusForbidden, "The provided token is not authorized.", PermissionType)
)
