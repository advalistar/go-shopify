package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const (
	commentsBasePath = "comments"
)

type CommentService interface {
	List(interface{}) ([]Comment, error)
	ListWithPagination(interface{}) ([]Comment, *Pagination, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*Comment, error)
	GetOrderList() []string
}

// CommentServiceOp handles communication with the comment related methods of the
// Shopify API.
type CommentServiceOp struct {
	client *Client
}

// Comment represents a Shopify comment.
type Comment struct {
	ID          int64      `json:"id"`
	Body        string     `json:"body"`
	BodyHTML    string     `json:"body_html"`
	Author      string     `json:"author"`
	Email       string     `json:"email"`
	Status      string     `json:"status"`
	ArticleID   int64      `json:"article_id"`
	BlogID      int64      `json:"blog_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Ip          string     `json:"ip"`
	UserAgent   string     `json:"user_agent"`
	PublishedAt *time.Time `json:"published_at"`
}

// CommentResource represents the result from the comments/X.json endpoint
type CommentResource struct {
	Comment *Comment `json:"comment"`
}

// CommentsResource represents the result from the comments.json endpoint
type CommentsResource struct {
	Comments []Comment `json:"comments"`
}

// List comments
func (s *CommentServiceOp) List(options interface{}) ([]Comment, error) {
	path := fmt.Sprintf("%s.json", commentsBasePath)
	resource := new(CommentsResource)
	err := s.client.Get(path, resource, options)
	return resource.Comments, err
}

func (s *CommentServiceOp) ListWithPagination(options interface{}) ([]Comment, *Pagination, error) {
	path := fmt.Sprintf("%s.json", commentsBasePath)
	resource := new(CommentsResource)
	headers := http.Header{}

	headers, err := s.client.createAndDoGetHeaders("GET", path, nil, options, resource)
	if err != nil {
		return nil, nil, err
	}

	// Extract pagination info from header
	linkHeader := headers.Get("Link")

	pagination, err := extractPagination(linkHeader)
	if err != nil {
		return nil, nil, err
	}

	return resource.Comments, pagination, nil
}

// Count comments
func (s *CommentServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", commentsBasePath)
	return s.client.Count(path, options)
}

// Get individual comment
func (s *CommentServiceOp) Get(commentID int64, options interface{}) (*Comment, error) {
	path := fmt.Sprintf("%s/%d.json", commentsBasePath, commentID)
	resource := new(CommentResource)
	err := s.client.Get(path, resource, options)
	return resource.Comment, err
}

func (s *CommentServiceOp) GetOrderList() []string {
	str := new(Comment)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
