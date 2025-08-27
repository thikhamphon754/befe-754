package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Book struct
type Book struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	ISBN        string  `json:"isbn"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	PublishYear int     `json:"publish_year"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var books = []Book{
	{
		ID:          "1",
		Title:       "Harry Potter and the Philosopher's Stone",
		Author:      "J.K. Rowling",
		ISBN:        "978-0747532699",
		Price:       350.00,
		Category:    "Fantasy",
		PublishYear: 1997,
		Stock:       15,
		Description: "เรื่องราวของเด็กชายที่ได้รับจดหมายเข้าโรงเรียนเวทมนตร์",
		Rating:      4.8,
	},
	{
		ID:          "2",
		Title:       "The Lord of the Rings",
		Author:      "J.R.R. Tolkien",
		ISBN:        "978-0544003415",
		Price:       450.00,
		Category:    "Fantasy",
		PublishYear: 1954,
		Stock:       8,
		Description: "การผจญภัยในมิดเดิลเอิร์ธเพื่อทำลายแหวนครองพิภพ",
		Rating:      4.9,
	},
	{
		ID:          "3",
		Title:       "To Kill a Mockingbird",
		Author:      "Harper Lee",
		ISBN:        "978-0060935467",
		Price:       280.00,
		Category:    "Classic Literature",
		PublishYear: 1960,
		Stock:       12,
		Description: "นิยายคลาสสิกเกี่ยวกับการเติบโตและความยุติธรรม",
		Rating:      4.7,
	},
	{
		ID:          "4",
		Title:       "1984",
		Author:      "George Orwell",
		ISBN:        "978-0451524935",
		Price:       320.00,
		Category:    "Science Fiction",
		PublishYear: 1949,
		Stock:       20,
		Description: "นิยายดิสโทเปียเกี่ยวกับสังคมที่ถูกควบคุม",
		Rating:      4.6,
	},
	{
		ID:          "5",
		Title:       "The Great Gatsby",
		Author:      "F. Scott Fitzgerald",
		ISBN:        "978-0743273565",
		Price:       290.00,
		Category:    "Classic Literature",
		PublishYear: 1925,
		Stock:       18,
		Description: "เรื่องราวความรักและความฝันในยุค Jazz Age",
		Rating:      4.4,
	},
}

// GET /api/v1/books - ดึงหนังสือทั้งหมด พร้อม query parameters
func getBooks(c *gin.Context) {
	// Query parameters
	categoryQuery := c.Query("category")
	authorQuery := c.Query("author")
	minPriceQuery := c.Query("min_price")
	maxPriceQuery := c.Query("max_price")
	publishYearQuery := c.Query("publish_year")

	filteredBooks := []Book{}

	for _, book := range books {
		// Filter by category
		if categoryQuery != "" && book.Category != categoryQuery {
			continue
		}

		// Filter by author
		if authorQuery != "" && book.Author != authorQuery {
			continue
		}

		// Filter by publish year
		if publishYearQuery != "" && fmt.Sprint(book.PublishYear) != publishYearQuery {
			continue
		}

		// Filter by minimum price
		if minPriceQuery != "" {
			minPrice, err := strconv.ParseFloat(minPriceQuery, 64)
			if err == nil && book.Price < minPrice {
				continue
			}
		}

		// Filter by maximum price
		if maxPriceQuery != "" {
			maxPrice, err := strconv.ParseFloat(maxPriceQuery, 64)
			if err == nil && book.Price > maxPrice {
				continue
			}
		}

		filteredBooks = append(filteredBooks, book)
	}

	// ถ้าไม่มี query parameters หรือไม่มีการกรอง ส่งข้อมูลทั้งหมด
	if len(filteredBooks) == 0 && (categoryQuery != "" || authorQuery != "" || minPriceQuery != "" || maxPriceQuery != "" || publishYearQuery != "") {
		c.JSON(http.StatusOK, gin.H{
			"message": "No books found matching the criteria",
			"data":    []Book{},
		})
		return
	}

	// ถ้าไม่มีการกรองเลย ส่งข้อมูลทั้งหมด
	if categoryQuery == "" && authorQuery == "" && minPriceQuery == "" && maxPriceQuery == "" && publishYearQuery == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Books retrieved successfully",
			"total":   len(books),
			"data":    books,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Books retrieved successfully",
		"total":   len(filteredBooks),
		"data":    filteredBooks,
	})
}

// GET /api/v1/books/:id - ดึงหนังสือตาม ID
func getBookByID(c *gin.Context) {
	bookID := c.Param("id")

	for _, book := range books {
		if book.ID == bookID {
			// เพิ่มข้อมูลเสริม
			response := gin.H{
				"message": "Book found successfully",
				"data": gin.H{
					"id":             book.ID,
					"title":          book.Title,
					"author":         book.Author,
					"isbn":           book.ISBN,
					"price":          book.Price,
					"category":       book.Category,
					"publish_year":   book.PublishYear,
					"stock":          book.Stock,
					"description":    book.Description,
					"rating":         book.Rating,
					"is_available":   book.Stock > 0,
					"stock_status":   getStockStatus(book.Stock),
					"price_with_vat": book.Price * 1.07, // VAT 7%
				},
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Book not found",
		"error":   fmt.Sprintf("Book with ID %s does not exist", bookID),
	})
}

// Helper function to determine stock status
func getStockStatus(stock int) string {
	if stock > 10 {
		return "In Stock"
	} else if stock > 0 {
		return "Limited Stock"
	}
	return "Out of Stock"
}

func main() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is healthy",
			"status":  "OK",
		})
	})

	// API v1 routes
	api := r.Group("/api/v1")
	{
		api.GET("/books", getBooks)       // ดึงหนังสือทั้งหมด
		api.GET("/books/:id", getBookByID) // ดึงหนังสือตาม ID
	}

	fmt.Println("🚀 Book Store API Server starting on port 8080...")
	fmt.Println("📚 Available endpoints:")
	fmt.Println("   GET  /health")
	fmt.Println("   GET  /api/v1/books")
	fmt.Println("   GET  /api/v1/books/:id")
	fmt.Println("📖 Query parameters for /api/v1/books:")
	fmt.Println("   ?category=Fantasy")
	fmt.Println("   ?author=J.K. Rowling")
	fmt.Println("   ?min_price=300")
	fmt.Println("   ?max_price=400")
	fmt.Println("   ?publish_year=1997")

	r.Run(":8080")
}