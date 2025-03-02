package constants

const (
	MONGODBURI   = "MONGODB_URI"
	RD_CONN      = "REDDIS_ADDR"
	REQUESTIDKEY = "requestID"
	RD_PASS      = "REDDIS_PASS"
	JWT_TOKEN    = "JWT_KEY"
	GMAIL_PASS   = "MAIL_SECURITY"
	OWNER_MAIL   = "OWNER_MAIL"
)

const ( //dbcrediantials
	DB_NAME = "hackmate"
)

const (
	COLLECTION_USERS = "users"
	COLLECTION_POSTS = "posts"
)

var MailHeader = map[int]string{
	1: "HireMate Connect",
	2: "Team Connect",
	3: "HireMate Support",
	4: "Emergency",
}

const ( //cloudinary crediantials
	CLOUDINARY_NAME           = "CLOUDINARY_NAME"
	CLOUDINARY_API_KEY        = "CLOUDINARY_API_KEY"
	CLOUDINARY_API_SECRET_KEY = "CLOUDINARY_API_SECRET_KEY"
)
