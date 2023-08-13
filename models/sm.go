package models

type Secret struct {
	Host     string `jason:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	JWTSign  string `json:"jwtsign"`
	Database string `json:"database"` /*las backstage `` se hacen con altgr y corchete*/
}
