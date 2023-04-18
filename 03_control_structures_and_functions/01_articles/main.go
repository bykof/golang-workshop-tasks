package main

import "fmt"

type Article struct {
	ID    int64
	Name  string
	Price int64
}

type ArticleRepository struct {
	articles []Article
}

func (ar ArticleRepository) ArticleById(id int64) Article {
	for _, article := range ar.articles {
		if article.ID == id {
			return article
		}
	}
	return Article{}
}

func NewArticleRepository() ArticleRepository {
	return ArticleRepository{
		articles: []Article{
			{
				ID:    0,
				Name:  "Mentos Mint",
				Price: 199,
			},
			{
				ID:    1,
				Name:  "Mentos fruit",
				Price: 149,
			},
			{
				ID:    2,
				Name:  "Mentos Pure Fresh Chewing Gums",
				Price: 199,
			},
			{
				ID:    3,
				Name:  "Schnuckeltüte",
				Price: 99,
			},
		},
	}
}

type OrderItem struct {
	Amount  int64
	Article Article
}

type Order struct {
	OrderItems []OrderItem
}

func addArticle(order Order, article Article, amount int64) Order {
	panic("Implement me")
}

func removeOrderItem(order Order, article Article, amount int64) Order {
	panic("Implement me")
}

func sumItems(order Order) int64 {
	panic("Implement me")
}

func main() {
	order := Order{}
	articleRepository := NewArticleRepository()

	// Add following articles to the order:
	// ID: 0 = 2x
	// ID: 1 = 5x
	// ID: 2 = 12x
	// ID: 3 = 4x

	if sumItems(order) != 3927 {
		panic("wrong")
	} else {
		fmt.Println("Correct ✅")
	}
}
