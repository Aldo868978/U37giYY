// 代码生成时间: 2025-09-20 07:27:49
package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/spf13/cobra"
)

// DbConfig contains the database configuration parameters
type DbConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// DBPool represents the database connection pool
type DBPool struct {
    *sql.DB
}

// NewDBPool creates a new database connection pool
func NewDBPool(cfg DbConfig) (*DBPool, error) {
    // Construct the DSN (Data Source Name)
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.DBName,
    )

    // Open the database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Set the maximum number of connections in the idle connection pool
    db.SetMaxIdleConns(10)

    // Set the maximum number of open connections to the database
    db.SetMaxOpenConns(100)

    // Set the maximum amount of time a connection may be reused
    db.SetConnMaxLifetime(3600 * time.Second) // 1 hour

    // Verify the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &DBPool{db}, nil
}

// Close closes the underlying database connection池
func (p *DBPool) Close() error {
    return p.DB.Close()
}

// NewCommand creates a new cobra command
func NewCommand() *cobra.Command {
    var cfg DbConfig
    var command = &cobra.Command{
        Use:   "dbpool",
        Short: "Manage database connection pool",
        Run: func(cmd *cobra.Command, args []string) {
            dbPool, err := NewDBPool(cfg)
            if err != nil {
                log.Fatalf("Failed to create DB pool: %v", err)
            }
            defer dbPool.Close()

            fmt.Println("Database connection pool is ready.")
        },
    }

    command.PersistentFlags().StringVar(&cfg.Host, "host", "localhost", "Database host")
    command.PersistentFlags().IntVar(&cfg.Port, "port", 3306, "Database port")
    command.PersistentFlags().StringVar(&cfg.User, "user", "root", "Database user")
    command.PersistentFlags().StringVar(&cfg.Password, "password", "", "Database password")
    command.PersistentFlags().StringVar(&cfg.DBName, "dbname", "test", "Database name")

    return command
}

func main() {
    cmd := NewCommand()
    if err := cmd.Execute(); err != nil {
        fmt.Println(err)
    }
}