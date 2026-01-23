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
	TitleEmpty       = New("PS014", "Title cannot be empty.", ValidationType)
	TitleTooLong     = New("PS015", "Title cannot exceed 200 characters.", ValidationType)
	EmailEmpty       = New("PS016", "Email cannot be empty.", ValidationType)
	EmailInvalid     = New("PS017", "Invalid email format.", ValidationType)
	ContentEmpty     = New("PS018", "Markdown content cannot be empty.", ValidationType)
	ContentTooShort  = New("PS019", "Markdown content must be at least 10 characters.", ValidationType)
	ContentTooLong   = New("PS020", "Markdown content cannot exceed 50000 characters.", ValidationType)
	CoverURLInvalid  = New("PS021", "Invalid cover image URL provided.", ValidationType)
	CoverURLNotImage = New("PS022", "URL must point to an image.", ValidationType)
	MissingAuthor    = New("PS023", "Post must have an author.", ValidationType)

	// Permission errors (PS030-PS039)
	ErrUnauthorized = New("PS030", "The provided token is not valid.", PermissionType)
	ErrForbidden    = New("PS031", "The provided token is not authorized.", PermissionType)
)
