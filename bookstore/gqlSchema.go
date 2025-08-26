package bookstore

import "github.com/graphql-go/graphql"

var authorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"location": &graphql.Field{
			Type: graphql.String,
		},
		//"books": &graphql.Field{
		//	Type: graphql.NewList(booktype),
		//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//		authId, _ := p.Source.(Author)
		//		for _, book := range Books {
		//			if book.Author == authId.Id {
		//				return book, nil
		//			}
		//		}
		//		return nil, nil
		//	},
		//},
	},
})

var publisherType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Publisher",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"location": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"books": &graphql.Field{
			Type: graphql.NewList(booktype),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var booksUser []Book
				user := p.Source.(User)
				for _, i := range Bookstores {
					if i.UserId == user.Id {
						for _, book := range Books {
							if book.Id == i.BookId {
								booksUser = append(booksUser, book)
							}
						}
					}
				}
				return booksUser, nil
			},
		},
	},
})

var booktype = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: authorType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				book, _ := p.Source.(Book)
				for _, auth := range Authors {
					if auth.Id == book.Author {
						return auth, nil
					}
				}
				return nil, nil
			},
		},
		"publisher": &graphql.Field{
			Type: publisherType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				auth, _ := p.Source.(Book)
				for _, pub := range Publishers {
					if pub.Id == auth.Publisher {
						return pub, nil
					}
				}
				return nil, nil
			},
		},
		"published": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var bookstoreType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Bookstore",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"book": &graphql.Field{
			Type: booktype,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				bookId, _ := p.Source.(Bookstore)
				for _, book := range Books {
					if bookId.BookId == book.Author {
						return book, nil
					}

				}
				return nil, nil
			},
		},
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId, _ := p.Source.(Bookstore)
				for _, user := range Users {
					if user.Id == userId.UserId {
						return user, nil
					}
				}
				return nil, nil
			},
		},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getAllBooks": &graphql.Field{
			Type: graphql.NewList(booktype),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Books, nil
			},
		},
		"getBook": &graphql.Field{
			Type: booktype,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				bookId, _ := p.Args["id"].(int)
				for _, book := range Books {
					if bookId == book.Id {
						return book, nil
					}
				}
				return nil, nil
			},
		},
		"getAllAuthors": &graphql.Field{
			Type: graphql.NewList(authorType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Authors, nil
			},
		},
		"getAuthor": &graphql.Field{
			Type: authorType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorId, _ := p.Args["id"].(int)
				for _, author := range Authors {
					if author.Id == authorId {
						return author, nil
					}
				}
				return nil, nil
			},
		},
		"getAllPublishers": &graphql.Field{
			Type: graphql.NewList(publisherType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Publishers, nil
			},
		},
		"getPublisher": &graphql.Field{
			Type: publisherType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				publisherId, _ := p.Args["id"].(int)
				for _, publisher := range Publishers {
					if publisher.Id == publisherId {
						return publisher, nil
					}
				}
				return nil, nil
			},
		},
		"getUser": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId, _ := p.Args["id"].(int)
				for _, user := range Users {
					if user.Id == userId {
						return user, nil
					}
				}
				return nil, nil
			},
		},
	},
})
