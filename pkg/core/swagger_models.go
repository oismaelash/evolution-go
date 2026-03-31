package core

import instance_model "github.com/EvolutionAPI/evolution-go/pkg/instance/model"

// --- Generic Responses ---

// SuccessResponse represents a standard success response without data
type SuccessResponse struct {
	Message string `json:"message" example:"success"`
}

// GenericResponse is kept for reference but explicit structs are preferred for better Swagger UI rendering
type GenericResponse[T any] struct {
	Message string `json:"message" example:"success"`
	Data    T      `json:"data,omitempty"`
}

// --- Error Responses ---

// ErrorMeta represents the metadata of the error response
type ErrorMeta struct {
	Timestamp string `json:"timestamp" example:"2024-01-15T10:30:00Z"`
	Path      string `json:"path" example:"/api/path"`
	Method    string `json:"method" example:"GET"`
}

// 400 Bad Request
type Error400Detail struct {
	Code    string `json:"code" example:"BAD_REQUEST"`
	Message string `json:"message" example:"Invalid request data"`
}
type Error400 struct {
	Success bool           `json:"success" example:"false"`
	Error   Error400Detail `json:"error"`
	Meta    ErrorMeta      `json:"meta"`
}

// 401 Unauthorized
type Error401Detail struct {
	Code    string `json:"code" example:"UNAUTHORIZED"`
	Message string `json:"message" example:"Invalid or missing API key"`
}
type Error401 struct {
	Success bool           `json:"success" example:"false"`
	Error   Error401Detail `json:"error"`
	Meta    ErrorMeta      `json:"meta"`
}

// 403 Forbidden
type Error403Detail struct {
	Code    string `json:"code" example:"FORBIDDEN"`
	Message string `json:"message" example:"Insufficient permissions"`
}
type Error403 struct {
	Success bool           `json:"success" example:"false"`
	Error   Error403Detail `json:"error"`
	Meta    ErrorMeta      `json:"meta"`
}

// 404 Not Found
type Error404Detail struct {
	Code    string `json:"code" example:"NOT_FOUND"`
	Message string `json:"message" example:"Resource not found"`
}
type Error404 struct {
	Success bool           `json:"success" example:"false"`
	Error   Error404Detail `json:"error"`
	Meta    ErrorMeta      `json:"meta"`
}

// 500 Internal Server Error
type Error500Detail struct {
	Code    string `json:"code" example:"INTERNAL_SERVER_ERROR"`
	Message string `json:"message" example:"An unexpected error occurred"`
}
type Error500 struct {
	Success bool           `json:"success" example:"false"`
	Error   Error500Detail `json:"error"`
	Meta    ErrorMeta      `json:"meta"`
}

// --- Module Specific Responses ---

// License
type StatusResponse struct {
	Status     string `json:"status" example:"active"`
	InstanceID string `json:"instance_id" example:"inst-12345"`
	APIKey     string `json:"api_key,omitempty" example:"evol...xyz"`
}

type RegisterResponse struct {
	Status      string `json:"status" example:"pending"`
	RegisterURL string `json:"register_url,omitempty" example:"https://app.evolution-api.com/register/12345"`
	Message     string `json:"message,omitempty" example:"License is already active"`
}

type ActivateResponse struct {
	Status  string `json:"status" example:"active"`
	Message string `json:"message" example:"License activated successfully!"`
}

// Call
type CallRejectResponse struct {
	Message string `json:"message" example:"success"`
}

// Chat
type ChatActionData struct {
	Timestamp int64 `json:"timestamp" example:"1705314600"`
}

type ChatActionResponse struct {
	Message string         `json:"message" example:"success"`
	Data    ChatActionData `json:"data"`
}

type HistorySyncData struct {
	MessageID string `json:"messageId" example:"3EB00000000000000000"`
}

type HistorySyncResponse struct {
	Message string          `json:"message" example:"success"`
	Data    HistorySyncData `json:"data"`
}

// Message / Send Message
type MessageKey struct {
	RemoteJid string `json:"remoteJid" example:"5511999999999@s.whatsapp.net"`
	FromMe    bool   `json:"fromMe" example:"true"`
	ID        string `json:"id" example:"3EB00000000000000000"`
}

type MessageSendData struct {
	Key              MessageKey `json:"key"`
	MessageTimestamp int64      `json:"messageTimestamp" example:"1705314600"`
	Status           string     `json:"status" example:"PENDING"`
}

type SendMessageResponse struct {
	Message string          `json:"message" example:"success"`
	Data    MessageSendData `json:"data"`
}

// Community
type CommunityResponse struct {
	JID string `json:"jid" example:"1234567890@g.us"`
}

// CommunityFullResponse
type CommunityFullResponse struct {
	Message string            `json:"message" example:"success"`
	Data    CommunityResponse `json:"data"`
}

// Newsletter
type NewsletterMetadata struct {
	ID              string `json:"id" example:"1234567890@newsletter"`
	Name            string `json:"name" example:"Evolution API Channel"`
	Description     string `json:"description" example:"Updates about Evolution API"`
	SubscriberCount int    `json:"subscriberCount" example:"150"`
	InviteCode      string `json:"inviteCode" example:"AbCdEfGh1234"`
}

type NewsletterResponse struct {
	Message string             `json:"message" example:"success"`
	Data    NewsletterMetadata `json:"data"`
}

type NewsletterListResponse struct {
	Message string               `json:"message" example:"success"`
	Data    []NewsletterMetadata `json:"data"`
}

type NewsletterMessage struct {
	ID        string `json:"id" example:"3EB00000000000000000"`
	Text      string `json:"text" example:"Hello World"`
	Timestamp int64  `json:"timestamp" example:"1705314600"`
}

type NewsletterMessagesResponse struct {
	Message string              `json:"message" example:"success"`
	Data    []NewsletterMessage `json:"data"`
}

// Polls
type PollOptionCounts struct {
	Option1Hash int `json:"option_1_hash" example:"5"`
	Option2Hash int `json:"option_2_hash" example:"3"`
}

// PollResultsData represents the detailed results of a poll
type PollResultsData struct {
	PollMessageID string           `json:"pollMessageId" example:"3EB00000000000000000"`
	PollChatJid   string           `json:"pollChatJid" example:"5511999999999@s.whatsapp.net"`
	TotalVotes    int              `json:"totalVotes" example:"10"`
	OptionCounts  PollOptionCounts `json:"optionCounts"`
	Voters        []VoterInfo      `json:"voters"`
}

type PollResultsResponse struct {
	Message string          `json:"message" example:"success"`
	Data    PollResultsData `json:"data"`
}

type VoterInfo struct {
	Jid             string   `json:"jid" example:"5511999999999@s.whatsapp.net"`
	Name            string   `json:"name" example:"John Doe"`
	SelectedOptions []string `json:"selectedOptions" example:"option_1_hash"`
	VotedAt         string   `json:"votedAt" example:"2024-01-15T10:30:00Z"`
}

// User
type PrivacySettingsData struct {
	ReadReceipts string `json:"readreceipts" example:"all"`
	Profile      string `json:"profile" example:"all"`
	Status       string `json:"status" example:"all"`
	Online       string `json:"online" example:"all"`
	LastSeen     string `json:"last" example:"all"`
	GroupAdd     string `json:"groupadd" example:"all"`
}

type PrivacySettingsResponse struct {
	Message string              `json:"message" example:"success"`
	Data    PrivacySettingsData `json:"data"`
}

type UserBlockData struct {
	JID string `json:"jid" example:"5511999999999@s.whatsapp.net"`
}

type UserBlockResponse struct {
	Message string        `json:"message" example:"success"`
	Data    UserBlockData `json:"data"`
}

type BlocklistData struct {
	JIDs []string `json:"jids" example:"5511999999999@s.whatsapp.net,5511988888888@s.whatsapp.net"`
}

type BlocklistResponse struct {
	Message string        `json:"message" example:"success"`
	Data    BlocklistData `json:"data"`
}

type UserProfileData struct {
	Timestamp int64 `json:"timestamp" example:"1705314600"`
}

type UserProfileResponse struct {
	Message string          `json:"message" example:"success"`
	Data    UserProfileData `json:"data"`
}

type AvatarData struct {
	URL string `json:"url" example:"https://pps.whatsapp.net/v/t61.24694-24/12345678_123456789012345_1234567890123456789_n.jpg"`
}

type AvatarResponse struct {
	Message string     `json:"message" example:"success"`
	Data    AvatarData `json:"data"`
}

type UserInfoData struct {
	VerifiedName string `json:"verifiedName,omitempty" example:"John Doe"`
	Status       string `json:"status" example:"Hey there! I am using WhatsApp."`
	PictureId    string `json:"pictureId,omitempty" example:"1234567890"`
}

type UserInfoResponse struct {
	Message string         `json:"message" example:"success"`
	Data    []UserInfoData `json:"data"`
}

type ContactInfoData struct {
	PushName string `json:"pushName" example:"John Doe"`
	Jid      string `json:"jid" example:"5511999999999@s.whatsapp.net"`
}

type ContactListResponse struct {
	Message string            `json:"message" example:"success"`
	Data    []ContactInfoData `json:"data"`
}

type IsOnWhatsAppResponse struct {
	Exists bool   `json:"exists" example:"true"`
	Jid    string `json:"jid" example:"5511999999999@s.whatsapp.net"`
}

type IsOnWhatsAppListResponse struct {
	Message string                 `json:"message" example:"success"`
	Data    []IsOnWhatsAppResponse `json:"data"`
}

// InstanceResponse
type InstanceResponse struct {
	Message string                  `json:"message" example:"success"`
	Data    instance_model.Instance `json:"data"`
}

// InstanceListResponse
type InstanceListResponse struct {
	Message string                    `json:"message" example:"success"`
	Data    []instance_model.Instance `json:"data"`
}

// AdvancedSettingsResponse
type AdvancedSettingsResponse struct {
	Message string                          `json:"message" example:"success"`
	Data    instance_model.AdvancedSettings `json:"data"`
}

// Server
type ServerOkResponse struct {
	Status string `json:"status" example:"ok"`
}

// Instance
type ConnectResponseData struct {
	Qrcode      string `json:"qrcode,omitempty" example:"1@...|..."`
	PairingCode string `json:"pairingCode,omitempty" example:"ABC1DEF2"`
}

type ConnectResponse struct {
	Message string              `json:"message" example:"success"`
	Data    ConnectResponseData `json:"data"`
}

type QRResponse struct {
	Message string `json:"message" example:"success"`
	Data    struct {
		Qrcode string `json:"qrcode" example:"1@...|..."`
	} `json:"data"`
}

type PairResponse struct {
	Message string `json:"message" example:"success"`
	Data    struct {
		PairingCode string `json:"pairingCode" example:"ABC1DEF2"`
	} `json:"data"`
}

// Group

// GroupInviteResponse
type GroupInviteResponse struct {
	Message string `json:"message" example:"success"`
	Data    string `json:"data" example:"https://chat.whatsapp.com/..."`
}

// GroupPhotoResponse
type GroupPhotoResponse struct {
	Message string `json:"message" example:"success"`
	Data    struct {
		PictureID string `json:"pictureId" example:"1234567890"`
	} `json:"data"`
}

// GroupInfoResponse
type GroupInfoResponse struct {
	Message string    `json:"message" example:"success"`
	Data    GroupInfo `json:"data"`
}

type GroupInfo struct {
	JID      string `json:"jid" example:"1234567890@g.us"`
	Name     string `json:"name" example:"Group Name"`
	Owner    string `json:"owner" example:"5511999999999@s.whatsapp.net"`
	ReadOnly bool   `json:"isReadOnly" example:"false"`
}

type GroupListResponse struct {
	Message string      `json:"message" example:"success"`
	Data    []GroupInfo `json:"data"`
}

// Instance
type InstanceStatusData struct {
	Connected bool   `json:"connected" example:"true"`
	LoggedIn  bool   `json:"loggedIn" example:"true"`
	MyJid     string `json:"myJid" example:"5511999999999@s.whatsapp.net"`
	Name      string `json:"name" example:"Instance Name"`
}

type InstanceStatusResponse struct {
	Message string             `json:"message" example:"success"`
	Data    InstanceStatusData `json:"data"`
}

type QRData struct {
	Qrcode string `json:"qrcode" example:"1@...|..."`
	Code   string `json:"code" example:"1234567890"`
}

type QRFullResponse struct {
	Message string `json:"message" example:"success"`
	Data    QRData `json:"data"`
}

type ConnectData struct {
	JID         string `json:"jid" example:"5511999999999@s.whatsapp.net"`
	WebhookURL  string `json:"webhookUrl" example:"http://localhost:8080/webhook"`
	EventString string `json:"eventString" example:"MESSAGE,GROUP_UP"`
}

type ConnectFullResponse struct {
	Message string      `json:"message" example:"success"`
	Data    ConnectData `json:"data"`
}

// Label
type LabelData struct {
	ID           string `json:"id" example:"uuid-string"`
	InstanceID   string `json:"instance_id" example:"uuid-string"`
	LabelID      string `json:"label_id" example:"1"`
	LabelName    string `json:"label_name" example:"Work"`
	LabelColor   string `json:"label_color" example:"#dfaef0"`
	PredefinedID string `json:"predefined_id" example:"1"`
}

type LabelListResponse struct {
	Message string      `json:"message" example:"success"`
	Data    []LabelData `json:"data"`
}

// MapData is a fallback for generic JSON objects
type MapData map[string]interface{}
