package errs

var (
	// Internal errors (PS001-PS009)
	InternalServerError = New("PS001", "An unexpected error happened. Please contact support.", InternalType)
	DatabaseError       = New("PS002", "An unexpected database error happened.", InternalType)

	// Validation errors (PS010-PS029)
	RequestBinding   = New("PS010", "The request body is malformed. Please refer to the API documentation.", ValidationType)
	PostNotFound     = New("PS011", "Post not found.", ValidationType)
	AuthorNotFound   = New("PS012", "Author not found.", ValidationType)
	DuplicateError   = New("PS013", "Duplicate key value violates unique constraint.", ValidationType)
	InvalidTitle     = New("PS014", "Invalid title provided.", ValidationType)
	InvalidEmail     = New("PS015", "Invalid email format.", ValidationType)
	InvalidContent   = New("PS016", "Invalid content provided.", ValidationType)
	InvalidCoverURL  = New("PS017", "Invalid cover image URL provided.", ValidationType)
	MissingAuthor    = New("PS018", "Post must have an author.", ValidationType)

	// Permission errors (PS030-PS039)
	ErrUnauthorized = New("PS030", "The provided token is not valid.", PermissionType)
	ErrForbidden    = New("PS031", "The provided token is not authorized.", PermissionType)
)
