package user

const listUserQuery = `
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

const getUserQuery = `
	SELECT 
		id, 
		name, 
		email,
		created_at, 
		updated_at
	FROM user
	WHERE id = ?
`

const createUserQuery = `
	INSERT INTO
		user (
			name,
			email,
			created_at
		)
	VALUES
		( ? , ? , NOW() )
`

const updateUserQuery = `
	UPDATE
		user
	SET
		name = ? ,
		email = ? ,
		updated_at = NOW()
	WHERE id = ?		
`

const deleteUserQuery = `
	DELETE FROM user WHERE id = ?
`
