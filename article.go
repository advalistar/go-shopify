package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const (
	articlesBasePath    = "articles"
	articlesAuthorsPath = "authors"
	articlesTagsPath    = "tags"
	articlesBlogsPath   = "blogs"
)

type ArticleService interface {
	AuthorsList(interface{}) ([]string, error)
	TagsList(interface{}) ([]string, error)
	List(int64, interface{}) ([]Article, error)
	ListWithPagination(int64, interface{}) ([]Article, *Pagination, error)
	GetOrderList() []string
}

// ArticleServiceOp handles communication with the article related methods of
// the Shopify API.
type ArticleServiceOp struct {
	client *Client
}

type ArticlesAuthorsResource struct {
	ArticlesAuthors []string `json:"authors"`
}

type ArticlesTagsResource struct {
	ArticlesTags []string `json:"tags"`
}

// Article represents a Shopify article
type Article struct {
	ID                int64      `json:"id"`
	Title             string     `json:"title"`
	CreatedAt         *time.Time `json:"created_at"`
	BodyHTML          string     `json:"body_html"`
	BlogID            int64      `json:"blog_id"`
	Author            string     `json:"author"`
	UserID            int64      `json:"user_id"`
	PublishedAt       *time.Time `json:"published_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	SummaryHTML       string     `json:"summary_html"`
	TemplateSuffix    string     `json:"template_suffix"`
	Handle            string     `json:"handle"`
	Tags              string     `json:"tags"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id"`
}

// ArticlesResource is the result from the articles.json endpoint
type ArticlesResource struct {
	Articles []Article `json:"articles"`
}

// Represents the result from the articles/X.json endpoint
type ArticleResource struct {
	Article *Article `json:"article"`
}

func (s *ArticleServiceOp) AuthorsList(options interface{}) ([]string, error) {
	path := fmt.Sprintf("%s/%s.json", articlesBasePath, articlesAuthorsPath)
	resource := new(ArticlesAuthorsResource)
	err := s.client.Get(path, resource, options)
	return resource.ArticlesAuthors, err
}

func (s *ArticleServiceOp) TagsList(options interface{}) ([]string, error) {
	path := fmt.Sprintf("%s/%s.json", articlesBasePath, articlesTagsPath)
	resource := new(ArticlesTagsResource)
	err := s.client.Get(path, resource, options)
	return resource.ArticlesTags, err
}

func (s *ArticleServiceOp) List(blogID int64, options interface{}) ([]Article, error) {
	path := fmt.Sprintf("%s/%d/%s.json", articlesBlogsPath, blogID, articlesBasePath)
	resource := new(ArticlesResource)
	err := s.client.Get(path, resource, options)
	return resource.Articles, err
}

func (s *ArticleServiceOp) ListWithPagination(blogID int64, options interface{}) ([]Article, *Pagination, error) {
	path := fmt.Sprintf("%s/%d/%s.json", articlesBlogsPath, blogID, articlesBasePath)
	resource := new(ArticlesResource)
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

	return resource.Articles, pagination, nil
}

func (s *ArticleServiceOp) GetOrderList() []string {
	str := new(Article)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
