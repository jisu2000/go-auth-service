package query

const (
	
	USER_REGISTER_QUERY_AS_NORMAL_USER string = 
			`INSERT INTO users 
				(name, email, password, mobile_number) 
				VALUES ($1, $2, $3, $4) 
			RETURNING id, created_at, name, email, mobile_number`
	
	
	FETCH_ALL_USER_QUERY string =  		
			`SELECT id,created_at,name,email,mobile_number FROM users`
)


