# Creating Instagram API using Golang

Aim: To create a basic API for instagram.

#### The API consists of following endpoints: <br/>
1. /users GET Request: Get all users
2. /users PUT Request: Create a user
3. /users/id GET Request: Get user by id
4. /posts GET Request: Get all posts
5. /posts PUT Request: Create a post
6. /posts/id GET Request: Get post by id

<br/>
<br/>

#### The model of the API looks like:

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
}

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string             `json:"caption,omitempty"`
	ImageUrl  string             `json:"imageUrl,omitempty"`
	Timestamp time.Time          `json:"timestamp,omitempty"`
	
}
