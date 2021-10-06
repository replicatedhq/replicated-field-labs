package fieldlabs

type MemberList struct {
	Id                string `json:"id"`
	Email             string `json:"email"`
	Is_Pending_Invite bool   `json:"is_pending_invite"`
}
