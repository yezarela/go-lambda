package user

// ListUserQuery ...
const ListUserQuery = `
	SELECT 
		id, 
		name, 
		email,
		created_at, 
		updated_at
	FROM user
	ORDER BY %s %s 
	LIMIT %d OFFSET %d
`

// GetUserQuery ...
const GetUserQuery = `
	SELECT 
		id, 
		name, 
		email,
		created_at, 
		updated_at
	FROM user
	WHERE id = ?
`

// CreateUserQuery ...
const CreateUserQuery = `
	INSERT INTO
		user (
			name,
			email,
			created_at
		)
	VALUES
		( ? , ? , NOW() )
`

// UpdateUserQuery ...
const UpdateUserQuery = `
	UPDATE
		user
	SET
		name = ? ,
		email = ? ,
		updated_at = NOW()
	WHERE id = ?		
`

// DeleteUserQuery ...
const DeleteUserQuery = `
	DELETE FROM user WHERE id = ?
`
