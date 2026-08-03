package main

// Main Models
type Book struct {
	ID			int		`json:"id"`
	Title		string	`json:"title"`
	Author		string	`json:"author"`
	Category	string	`json:"category"`
	Stock		int		`json:"stock"`
}

type Member struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
}

type BorrowRecord struct {
	ID			int		`json:"id"`
	BookID		int		`json:"bookId"`
	MemberID	int		`json:"memberId"`
	BorrowDate	string	`json:"borrowDate"`
	ReturnDate	string	`json:"returnDate"`
	Returned 	bool	`json:"returned"`
}