package entity

// User :
type User struct {
	Name, Password, Email, Phone string
}

func (m_user User) init(tName, tPassword, tEmail, tPhone string) {
	m_user.Name = tName
	m_user.Password = tPassword
	m_user.Email = tEmail
	m_user.Phone = tPhone
}

// CopyUser = copy User
func (m_user User) CopyUser(tuser User) {
	m_user.Name = tuser.Name
	m_user.Password = tuser.Password
	m_user.Email = tuser.Email
	m_user.Phone = tuser.Phone
}

// GetName = getname
func (m_user User) GetName() string {
	return m_user.Name
}

// SetName = setname
func (m_user User) SetName(tname string) {
	m_user.Name = tname
}

// GetPassword = GetPassword
func (m_user User) GetPassword() string {
	return m_user.Password
}

// SetPassword = SetPassword
func (m_user User) SetPassword(tpassword string) {
	m_user.Password = tpassword
}

// GetEmail = GetPassword
func (m_user User) GetEmail() string {
	return m_user.Email
}

// SetEmail = SetPassword
func (m_user User) SetEmail(temail string) {
	m_user.Email = temail
}

// GetPhone = GetPassword
func (m_user User) GetPhone() string {
	return m_user.Phone
}

// SetPhone = SetPassword
func (m_user User) SetPhone(tphone string) {
	m_user.Phone = tphone
}
