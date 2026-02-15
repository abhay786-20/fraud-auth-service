package repository

import (
	"github.com/abhay786-20/fraud-auth-service/internal/models"
	"github.com/abhay786-20/fraud-auth-service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

// =============================================================================
// INTERFACE - The Contract
// =============================================================================
// UserRepository defines WHAT operations can be done with users.
// It's like a job description - anyone implementing this MUST have these methods.
//
// WHY INTERFACE?
// - Flexibility: Can swap Postgres for MongoDB without changing service layer
// - Testing: Can create mock repository for unit tests
// - Decoupling: Service layer depends on interface, not concrete implementation
//
// USAGE IN SERVICE:
//
//	type AuthService struct {
//	    userRepo UserRepository  // Uses interface, not struct
//	}
type UserRepository interface {
	// Create inserts a new user into the database
	// Returns error if insert fails (e.g., duplicate email)
	Create(user *models.User) error

	// GetByEmail finds a user by their email address
	// Returns (*User, nil) if found, (nil, error) if not found or error occurs
	GetByEmail(email string) (*models.User, error)
}

// =============================================================================
// STRUCT - The Actual Implementation
// =============================================================================
// PostgresUserRepository is the concrete implementation of UserRepository.
// It holds the dependencies (tools) needed to do the job.
//
// FIELDS:
// - db:  Database connection to execute SQL queries
// - log: Logger to record what's happening (debugging, monitoring)
//
// NOTE: This struct is private (lowercase would make it unexported),
// but we expose it through the interface via the constructor.
type PostgresUserRepository struct {
	db  *sqlx.DB       // Database connection pool
	log *logger.Logger // Logger for recording operations
}

// =============================================================================
// CONSTRUCTOR - Creating the Repository
// =============================================================================
// NewPostgresUserRepository creates a new PostgresUserRepository instance.
//
// PARAMETERS:
// - db:  Active database connection (from bootstrap)
// - log: Logger instance (from bootstrap)
//
// RETURNS:
// - UserRepository (interface) - NOT *PostgresUserRepository (struct)
//
// WHY RETURN INTERFACE?
// The caller only needs to know "this thing can Create and GetByEmail".
// They don't need to know it's specifically Postgres underneath.
// This allows swapping implementations without changing caller code.
//
// EXAMPLE USAGE:
//
//	db := db.NewPostgres(cfg)
//	log := logger.New()
//	userRepo := repository.NewPostgresUserRepository(db.DB, log)
//	// userRepo is type UserRepository, but contains PostgresUserRepository
func NewPostgresUserRepository(db *sqlx.DB, log *logger.Logger) UserRepository {
	return &PostgresUserRepository{
		db:  db,  // Assign database connection to struct field
		log: log, // Assign logger to struct field
	}
}

// =============================================================================
// METHOD - Create User
// =============================================================================
// Create inserts a new user into the database.
//
// RECEIVER: (r *PostgresUserRepository)
// - 'r' is like 'this' or 'self' in other languages
// - '*' means pointer - we're working with the actual struct, not a copy
// - This gives us access to r.db and r.log
//
// PARAMETERS:
// - user: Pointer to User model with Email and Password filled in
//
// RETURNS:
// - error: nil if success, error if failed
//
// SIDE EFFECT:
// - user.ID, user.CreatedAt, user.UpdatedAt are populated after successful insert
//   (because we use RETURNING clause and Scan into the same user object)
//
// SQL EXPLANATION:
// - $1, $2 are placeholders (prevents SQL injection)
// - RETURNING gives back the auto-generated values
// - Scan puts those values into the user struct fields
func (r *PostgresUserRepository) Create(user *models.User) error {
	r.log.Info("Creating user with email: " + user.Email)

	// SQL query with placeholders ($1, $2) for safe parameter binding
	// RETURNING clause fetches auto-generated values after insert
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	// Execute query and scan returned values into user struct
	// QueryRow: executes query expecting single row result
	// Scan: puts returned columns into provided variables
	// &user.ID: & means "put the value HERE" (pointer/reference)
	err := r.db.QueryRow(
		query,
		user.Email,    // $1 - first placeholder
		user.Password, // $2 - second placeholder
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	// Error handling - always check and log errors
	if err != nil {
		r.log.Error("Failed to create user: " + err.Error())
		return err
	}

	r.log.Info("User created successfully with ID: " + user.ID)
	return nil
}

// =============================================================================
// METHOD - Get User By Email
// =============================================================================
// GetByEmail finds a user by their email address.
//
// PARAMETERS:
// - email: The email to search for
//
// RETURNS:
// - *models.User: Pointer to user if found
// - error: nil if found, error if not found or query failed
//
// NOTE: Returns pointer (*models.User) so we can return nil if not found.
// If we returned models.User (not pointer), we'd have to return empty struct.
//
// COMMON ERRORS:
// - sql.ErrNoRows: No user with this email exists
// - Other errors: Database connection issues, etc.
func (r *PostgresUserRepository) GetByEmail(email string) (*models.User, error) {
	r.log.Info("Fetching user by email: " + email)

	// Declare variable to hold result
	// 'var' creates zero-valued struct (empty strings, zero ints, etc.)
	var user models.User

	// SQL query - $1 is placeholder for email parameter
	query := `
		SELECT id, email, password, created_at
		FROM users
		WHERE email=$1
	`

	// db.Get: executes query and scans result into struct
	// &user: pointer to user struct where data will be stored
	// Unlike QueryRow().Scan(), Get() automatically maps columns to struct fields
	err := r.db.Get(&user, query, email)
	if err != nil {
		r.log.Error("Failed to fetch user by email: " + err.Error())
		return nil, err // Return nil user and the error
	}

	r.log.Info("User fetched successfully with ID: " + user.ID)
	return &user, nil // Return pointer to user and nil error
}
