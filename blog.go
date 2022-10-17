package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const blogsBasePath = "blogs"

// BlogService is an interface for interfacing with the blogs endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/online_store/blog
type BlogService interface {
	List(interface{}) ([]Blog, error)
	ListWithPagination(interface{}) ([]Blog, *Pagination, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*Blog, error)
	Create(Blog) (*Blog, error)
	Update(Blog) (*Blog, error)
	Delete(int64) error
	GetOrderList() []string
}

// BlogServiceOp handles communication with the blog related methods of
// the Shopify API.
type BlogServiceOp struct {
	client *Client
}

// Blog represents a Shopify blog
type Blog struct {
	ID                 int64      `json:"id"`
	Handle             string     `json:"handle"`
	Title              string     `json:"title"`
	UpdatedAt          *time.Time `json:"updated_at"`
	Commentable        string     `json:"commentable"`
	Feedburner         string     `json:"feedburner"`
	FeedburnerLocation string     `json:"feedburner_location"`
	CreatedAt          *time.Time `json:"created_at"`
	TemplateSuffix     string     `json:"template_suffix"`
	Tags               string     `json:"tags"`
	AdminGraphqlAPIID  string     `json:"admin_graphql_api_id"`
}

// BlogsResource is the result from the blogs.json endpoint
type BlogsResource struct {
	Blogs []Blog `json:"blogs"`
}

// Represents the result from the blogs/X.json endpoint
type BlogResource struct {
	Blog *Blog `json:"blog"`
}

// List all blogs
func (s *BlogServiceOp) List(options interface{}) ([]Blog, error) {
	path := fmt.Sprintf("%s.json", blogsBasePath)
	resource := new(BlogsResource)
	err := s.client.Get(path, resource, options)
	return resource.Blogs, err
}

func (s *BlogServiceOp) ListWithPagination(options interface{}) ([]Blog, *Pagination, error) {
	path := fmt.Sprintf("%s.json", blogsBasePath)
	resource := new(BlogsResource)
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

	return resource.Blogs, pagination, nil
}

// Count blogs
func (s *BlogServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", blogsBasePath)
	return s.client.Count(path, options)
}

// Get single blog
func (s *BlogServiceOp) Get(blogID int64, options interface{}) (*Blog, error) {
	path := fmt.Sprintf("%s/%d.json", blogsBasePath, blogID)
	resource := new(BlogResource)
	err := s.client.Get(path, resource, options)
	return resource.Blog, err
}

// Create a new blog
func (s *BlogServiceOp) Create(blog Blog) (*Blog, error) {
	path := fmt.Sprintf("%s.json", blogsBasePath)
	wrappedData := BlogResource{Blog: &blog}
	resource := new(BlogResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Blog, err
}

// Update an existing blog
func (s *BlogServiceOp) Update(blog Blog) (*Blog, error) {
	path := fmt.Sprintf("%s/%d.json", blogsBasePath, blog.ID)
	wrappedData := BlogResource{Blog: &blog}
	resource := new(BlogResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Blog, err
}

// Delete an blog
func (s *BlogServiceOp) Delete(blogID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", blogsBasePath, blogID))
}

func (s *BlogServiceOp) GetOrderList() []string {
	str := new(Blog)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
