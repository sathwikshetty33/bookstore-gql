// package bookstore

// import "github.com/graphql-go/graphql"

// var authorType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Author",
// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.Int,
// 		},
// 		"name": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"location": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		//"books": &graphql.Field{
// 		//	Type: graphql.NewList(booktype),
// 		//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		//		authId, _ := p.Source.(Author)
// 		//		for _, book := range Books {
// 		//			if book.Author == authId.Id {
// 		//				return book, nil
// 		//			}
// 		//		}
// 		//		return nil, nil
// 		//	},
// 		//},
// 	},
// })

// var publisherType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Publisher",
// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.Int,
// 		},
// 		"name": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"location": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 	},
// })

// var userType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "User",
// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.Int,
// 		},
// 		"name": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"email": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"books": &graphql.Field{
// 			Type: graphql.NewList(booktype),
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				var booksUser []Book
// 				user := p.Source.(User)
// 				for _, i := range Bookstores {
// 					if i.UserId == user.Id {
// 						for _, book := range Books {
// 							if book.Id == i.BookId {
// 								booksUser = append(booksUser, book)
// 							}
// 						}
// 					}
// 				}
// 				return booksUser, nil
// 			},
// 		},
// 	},
// })

// var booktype = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Book",
// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.Int,
// 		},
// 		"title": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"author": &graphql.Field{
// 			Type: authorType,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				book, _ := p.Source.(Book)
// 				for _, auth := range Authors {
// 					if auth.Id == book.Author {
// 						return auth, nil
// 					}
// 				}
// 				return nil, nil
// 			},
// 		},
// 		"publisher": &graphql.Field{
// 			Type: publisherType,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				auth, _ := p.Source.(Book)
// 				for _, pub := range Publishers {
// 					if pub.Id == auth.Publisher {
// 						return pub, nil
// 					}
// 				}
// 				return nil, nil
// 			},
// 		},
// 		"published": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 	},
// })

// var bookstoreType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Bookstore",
// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.Int,
// 		},
// 		"book": &graphql.Field{
// 			Type: booktype,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				bookId, _ := p.Source.(Bookstore)
// 				for _, book := range Books {
// 					if bookId.BookId == book.Author {
// 						return book, nil
// 					}

// 				}
// 				return nil, nil
// 			},
// 		},
// 		"user": &graphql.Field{
// 			Type: userType,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				userId, _ := p.Source.(Bookstore)
// 				for _, user := range Users {
// 					if user.Id == userId.UserId {
// 						return user, nil
// 					}
// 				}
// 				return nil, nil
// 			},
// 		},
// 	},
// })

// var RootQuery = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "RootQuery",
// 	Fields: graphql.Fields{
// 		"getAllBooks": &graphql.Field{
// 			Type: graphql.NewList(booktype),
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				return Books, nil
// 			},
// 		},
// 		"getBook": &graphql.Field{
// 			Type: booktype,
// 			Args: graphql.FieldConfigArgument{
// 				"id": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.Int),
// 				},
// 			},
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				bookId, _ := p.Args["id"].(int)
// 				for _, book := range Books {
// 					if bookId == book.Id {
// 						return book, nil
// 					}
// 				}
// 				return nil, nil
// 			},
// 		},
// 		"getAllAuthors": &graphql.Field{
// 			Type: graphql.NewList(authorType),
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				return Authors, nil
// 			},
// 		},
// 		"getAuthor": &graphql.Field{
// 			Type: authorType,
// 			Args: graphql.FieldConfigArgument{
// 				"id": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.Int),
// 				},
// 			},
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				authorId, _ := p.Args["id"].(int)
// 				for _, author := range Authors {
// 					if author.Id == authorId {
// 						return author, nil
// 					}
// 				}
// 				return nil, nil
// 			},
// 		},
// 		"getAllPublishers": &graphql.Field{
// 			Type: graphql.NewList(publisherType),
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				return Publishers, nil
// 			},
// 		},
// 		"getPublisher": &graphql.Field{
// 			Type: publisherType,
// 			Args: graphql.FieldConfigArgument{
// 				"id": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.Int),
// 				},
// 			},
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				publisherId, _ := p.Args["id"].(int)
// 				for _, publisher := range Publishers {
// 					if publisher.Id == publisherId {
// 						return publisher, nil
// 					}
// 				}
// 				return nil, nil
// 			},
// 		},
// 		"getUser": &graphql.Field{
// 			Type: userType,
// 			Args: graphql.FieldConfigArgument{
// 				"id": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.Int),
// 				},
// 			},
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				userId, _ := p.Args["id"].(int)
// 				for _, user := range Users {
// 					if user.Id == userId {
// 						return user, nil
// 					}
// 				}
// 				return nil, nil
// 			},
// 		},
// 	},
// })

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
		"books": &graphql.Field{
			Type: graphql.NewList(booktype),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authorBooks []Book
				author := p.Source.(Author)
				for _, book := range Books {
					if book.Author == author.Id {
						authorBooks = append(authorBooks, book)
					}
				}
				return authorBooks, nil
			},
		},
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
		"books": &graphql.Field{
			Type: graphql.NewList(booktype),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var publisherBooks []Book
				publisher := p.Source.(Publisher)
				for _, book := range Books {
					if book.Publisher == publisher.Id {
						publisherBooks = append(publisherBooks, book)
					}
				}
				return publisherBooks, nil
			},
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
				book, _ := p.Source.(Book)
				for _, pub := range Publishers {
					if pub.Id == book.Publisher {
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
				bookstore, _ := p.Source.(Bookstore)
				for _, book := range Books {
					if bookstore.BookId == book.Id {
						return book, nil
					}
				}
				return nil, nil
			},
		},
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				bookstore, _ := p.Source.(Bookstore)
				for _, user := range Users {
					if user.Id == bookstore.UserId {
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
		// Existing queries
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
		
		// NEW QUERIES ADDED BELOW
		
		"getAllUsers": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Users, nil
			},
		},
		
		"getAllBookstores": &graphql.Field{
			Type: graphql.NewList(bookstoreType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Bookstores, nil
			},
		},
		
		"getBooksByAuthor": &graphql.Field{
			Type: graphql.NewList(booktype),
			Args: graphql.FieldConfigArgument{
				"authorId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorId, _ := p.Args["authorId"].(int)
				var authorBooks []Book
				for _, book := range Books {
					if book.Author == authorId {
						authorBooks = append(authorBooks, book)
					}
				}
				return authorBooks, nil
			},
		},
		
		"getBooksByPublisher": &graphql.Field{
			Type: graphql.NewList(booktype),
			Args: graphql.FieldConfigArgument{
				"publisherId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				publisherId, _ := p.Args["publisherId"].(int)
				var publisherBooks []Book
				for _, book := range Books {
					if book.Publisher == publisherId {
						publisherBooks = append(publisherBooks, book)
					}
				}
				return publisherBooks, nil
			},
		},
		
		"getUserBooks": &graphql.Field{
			Type: graphql.NewList(booktype),
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId, _ := p.Args["userId"].(int)
				var userBooks []Book
				for _, bookstore := range Bookstores {
					if bookstore.UserId == userId {
						for _, book := range Books {
							if book.Id == bookstore.BookId {
								userBooks = append(userBooks, book)
							}
						}
					}
				}
				return userBooks, nil
			},
		},
		
		"searchBooks": &graphql.Field{
			Type: graphql.NewList(booktype),
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				title, hasTitle := p.Args["title"].(string)
				var matchingBooks []Book
				
				for _, book := range Books {
					if hasTitle && title != "" {
						// Simple case-insensitive substring search
						if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
							matchingBooks = append(matchingBooks, book)
						}
					}
				}
				return matchingBooks, nil
			},
		},
		
		"searchAuthors": &graphql.Field{
			Type: graphql.NewList(authorType),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"location": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name, hasName := p.Args["name"].(string)
				location, hasLocation := p.Args["location"].(string)
				var matchingAuthors []Author
				
				for _, author := range Authors {
					match := true
					
					if hasName && name != "" {
						if !strings.Contains(strings.ToLower(author.Name), strings.ToLower(name)) {
							match = false
						}
					}
					
					if hasLocation && location != "" {
						if !strings.Contains(strings.ToLower(author.Location), strings.ToLower(location)) {
							match = false
						}
					}
					
					if match && (hasName || hasLocation) {
						matchingAuthors = append(matchingAuthors, author)
					}
				}
				return matchingAuthors, nil
			},
		},
		
		"getBooksByYear": &graphql.Field{
			Type: graphql.NewList(booktype),
			Args: graphql.FieldConfigArgument{
				"year": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				year, _ := p.Args["year"].(string)
				var yearBooks []Book
				for _, book := range Books {
					if strings.Contains(book.Published, year) {
						yearBooks = append(yearBooks, book)
					}
				}
				return yearBooks, nil
			},
		},
		
		"getUsersByBook": &graphql.Field{
			Type: graphql.NewList(userType),
			Args: graphql.FieldConfigArgument{
				"bookId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				bookId, _ := p.Args["bookId"].(int)
				var bookUsers []User
				userMap := make(map[int]bool) // To avoid duplicates
				
				for _, bookstore := range Bookstores {
					if bookstore.BookId == bookId {
						if !userMap[bookstore.UserId] {
							for _, user := range Users {
								if user.Id == bookstore.UserId {
									bookUsers = append(bookUsers, user)
									userMap[bookstore.UserId] = true
									break
								}
							}
						}
					}
				}
				return bookUsers, nil
			},
		},
	},
})
