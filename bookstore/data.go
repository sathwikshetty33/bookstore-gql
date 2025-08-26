package bookstore

var Authors = []Author{
	{1, "J.K. Rowling", "British author, best known for the Harry Potter series."},
	{2, "George R.R. Martin", "American novelist and short story writer, author of A Song of Ice and Fire."},
	{3, "Haruki Murakami", "Japanese writer known for surrealistic and melancholic novels."},
	{4, "Chinua Achebe", "Nigerian novelist, poet, professor, and critic."},
	{5, "Agatha Christie", "English writer known for detective novels and short stories."},
}

var Publishers = []Publisher{
	{1, "Bloomsbury", "London, UK"},
	{2, "Bantam Books", "New York, USA"},
	{3, "Vintage International", "Tokyo, Japan"},
	{4, "Heinemann", "Lagos, Nigeria"},
	{5, "HarperCollins", "London, UK"},
}

var Books = []Book{
	{1, "Harry Potter and the Philosopher's Stone", 1, 499, 1, "1997-06-26"},
	{2, "Harry Potter and the Chamber of Secrets", 1, 599, 1, "1998-07-02"},
	{3, "A Game of Thrones", 2, 799, 2, "1996-08-06"},
	{4, "A Clash of Kings", 2, 899, 2, "1998-11-16"},
	{5, "Kafka on the Shore", 3, 699, 3, "2002-09-12"},
	{6, "Norwegian Wood", 3, 499, 3, "1987-09-04"},
	{7, "Things Fall Apart", 4, 399, 4, "1958-06-17"},
	{8, "No Longer at Ease", 4, 350, 4, "1960-10-01"},
	{9, "Murder on the Orient Express", 5, 450, 5, "1934-01-01"},
	{10, "And Then There Were None", 5, 480, 5, "1939-11-06"},
}

var Users = []User{
	{1, "Alice Johnson", "alice@example.com"},
	{2, "Bob Smith", "bob@example.com"},
	{3, "Charlie Brown", "charlie@example.com"},
	{4, "Diana Prince", "diana@example.com"},
	{5, "Ethan Hunt", "ethan@example.com"},
}

var Bookstores = []Bookstore{
	{1, 1, 1},   // Alice bought Harry Potter 1
	{2, 3, 2},   // Bob bought A Game of Thrones
	{3, 5, 3},   // Charlie bought Kafka on the Shore
	{4, 7, 4},   // Diana bought Things Fall Apart
	{5, 9, 5},   // Ethan bought Murder on the Orient Express
	{6, 2, 1},   // Alice also bought Harry Potter 2
	{7, 4, 2},   // Bob also bought A Clash of Kings
	{8, 6, 3},   // Charlie also bought Norwegian Wood
	{9, 8, 4},   // Diana also bought No Longer at Ease
	{10, 10, 5}, // Ethan also bought And Then There Were None
}
