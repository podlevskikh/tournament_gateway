package response_success

import "vollyemsk_tournament_gateway/models/users"

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

type UserResponse struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	NickName   string `json:"nickName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

func FromUsersResponse(us []*users.User) UsersResponse {
	urs := make([]UserResponse, 0, len(us))
	for _, u := range us {
		urs = append(urs, UserResponse{
			Id:         u.ID,
			FirstName:  u.FirstName,
			LastName:   u.LastName,
			MiddleName: u.MiddleName,
			NickName:   u.NickName,
			Email:      u.Email,
			Phone:      u.Phone,
		})
	}

	return UsersResponse{Users: urs}
}
