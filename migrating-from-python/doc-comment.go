// User returns a Userinfo containing the provided username
// and no password set.
func User(username string) *Userinfo {
        return &Userinfo{username, "", false}
}
