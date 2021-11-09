// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

// Categories is the golang structure for table categories.
type Categories struct {
	Id          uint      `orm:"id,primary"  json:"id"`          //
	Name        string    `orm:"name"        json:"name"`        //
	Description string    `orm:"description" json:"description"` //
	Topics      []*Topics `orm:"with:category_id=id"`
}

// Topics is the golang structure for table topics.
type Topics struct {
	Id              uint        `orm:"id,primary"         json:"id"`              //
	Title           string      `orm:"title"              json:"title"`           //
	Body            string      `orm:"body"               json:"body"`            //
	UserId          int         `orm:"user_id"            json:"userId"`          //
	CategoryId      int         `orm:"category_id"        json:"categoryId"`      //
	ReplyCount      int         `orm:"reply_count"        json:"replyCount"`      //
	ViewCount       int         `orm:"view_count"         json:"viewCount"`       //
	LastReplyUserId int         `orm:"last_reply_user_id" json:"lastReplyUserId"` //
	Order           int         `orm:"order"              json:"order"`           //
	Excerpt         string      `orm:"excerpt"            json:"excerpt"`         //
	Slug            string      `orm:"slug"               json:"slug"`            //
	Categories      *Categories `orm:"with:id=category_id"`
}