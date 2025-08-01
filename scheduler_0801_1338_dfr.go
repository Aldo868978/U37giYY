// 代码生成时间: 2025-08-01 13:38:14
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/robfig/cron/v3"
)

// Scheduler is the main struct for the scheduler application
type Scheduler struct {
    Cron *cron.Cron
}

// NewScheduler creates a new instance of Scheduler
func NewScheduler() *Scheduler {
    return &Scheduler{
        Cron: cron.New(),
    }
}

// AddJob adds a new job to the scheduler
func (s *Scheduler) AddJob(spec string, cmd func()) error {
    _, err := s.Cron.AddFunc(spec, cmd)
    return err
}

// Start starts the scheduler
func (s *Scheduler) Start() {
    s.Cron.Start()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
    s.Cron.Stop()
}

// RunJob defines a function that will be executed by the scheduler
func RunJob() {
    fmt.Println("Job is running...")
}

func main() {
    // Create a new scheduler instance
    scheduler := NewScheduler()

    // Add a job that runs every minute
    if err := scheduler.AddJob("* * * * *", RunJob); err != nil {
        log.Fatalf("Failed to schedule job: %v", err)
    }

    // Start the scheduler
    scheduler.Start()

    // Run the scheduler in an infinite loop to keep the application running
    select{}
}
