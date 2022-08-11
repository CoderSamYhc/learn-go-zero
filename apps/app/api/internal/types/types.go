// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Category    string `json:"category"`
	Status      int64  `json:"status"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

type Comment struct {
	ID        int64    `json:"id"`
	ProductID int64    `json:"product_id"`
	Content   string   `json:"content"`
	Images    []*Image `json:"images"`
	User      *User    `json:"user"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

type Image struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type CartProduct struct {
	Product *Product `json:"product"`
	Count   int64    `json:"count"`
}

type Order struct {
	OrderID      string `json:"order_id"`
	Quantity     int64  `json:"quantity"`
	ProductID    int64  `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
}

type RecommendRequest struct {
	Cursor int64 `form:"cursor"`
	Ps     int64 `form:"ps,default=20"` // 分页大小
}

type RecommendResponse struct {
	Products      []*Product `json:"products"`
	IsEnd         bool       `json:"is_end"`
	RecommendTime int64      `json:"recommend_time"`
}

type FlashSaleResponse struct {
	StartTime int64      `json:"start_time"`
	Products  []*Product `json:"products"`
}

type CategoryListRequest struct {
	Cursor   int64  `form:"cursor"`
	Ps       int64  `form:"ps,default=20"` // 分页大小
	Category string `form:"category"`
	Sort     string `form:"sort"`
}

type CategoryListResponse struct {
	Products []*Product `json:"products"`
	IsEnd    bool       `json:"is_end"`
	LastVal  int64      `json:"last_val"`
}

type CartListRequest struct {
	UID int64 `json:"uid"`
}

type CartListResponse struct {
	Products []*Product `json:"products"`
}

type ProductCommentRequest struct {
	ProductID int64 `form:"product_id"`
	Cursor    int64 `form:"cursor"`
	Ps        int64 `form:"ps,default=20"`
}

type ProductCommentResponse struct {
	Comments    []*Comment `json:"comments"`
	IsEnd       bool       `json:"is_end"`
	CommentTime int64      `json:"comment_time"`
}

type OrderListRequest struct {
	UID    int64 `json:"uid"`
	Cursor int64 `json:"cursor"`
	Ps     int64 `json:"ps,default=20"` // 分页大小
	Status int32 `json:"status"`
}

type OrderListResponse struct {
	Orders    []*Order `json:"orders"`
	IsEnd     bool     `json:"is_end"`
	OrderTime int64    `json:"order_time"`
}
