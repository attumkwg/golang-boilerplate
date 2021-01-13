package domain

// Users はuserを管理するための定義です。
type Users struct {
  Id int
  ScreenName string
  DisplayName string
  Password string
  Email *string
  CreatedAt string
  UpdatedAt string
}

// UsersForGet .
type UsersForGet struct {
  Id int `json:"id"`
  ScreenName string `json:"screenName"`
  DisplayName string `json:"displayName"`
  Email *string `json:"email"`
}

// BuildForGet .
func (u *Users) BuildForGet() UsersForGet {
  user := UsersForGet{}
  user.Id = u.Id
  user.ScreenName = u.ScreenName
  user.DisplayName = u.DisplayName
  if u.Email != nil {
    user.Email = u.Email
  } else {
    empty := ""
    user.Email = &empty
  }
  return user
}
