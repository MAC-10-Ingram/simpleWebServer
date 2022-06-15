package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Content 1"},
	{ID: 2, Title: "Article 2", Content: "Content 2"},
}

func getAllArticles() []article {
	return articleList
}
