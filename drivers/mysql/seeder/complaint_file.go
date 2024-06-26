package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedComplaintFile(db *gorm.DB) {
	if err := db.First(&entities.ComplaintFile{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		complaintFiles := []entities.ComplaintFile{
			{
				ComplaintID: "C-81j9aK9280",
				Path:        "complaint-files/example1.jpg",
			},
			{
				ComplaintID: "C-81j9aK9280",
				Path:        "complaint-files/example2.jpg",
			},
			{
				ComplaintID: "C-8ksh&s9280",
				Path:        "complaint-files/example3.jpg",
			},
			{
				ComplaintID: "C-8ksh&s9280",
				Path:        "complaint-files/example3.jpg",
			},
			{
				ComplaintID: "C-81jas92581",
				Path:        "complaint-files/example3.jpg",
			},
			{
				ComplaintID: "C-81jas92581",
				Path:        "complaint-files/example1.jpg",
			},
			{
				ComplaintID: "C-271j9ak280",
				Path:        "complaint-files/example2.jpg",
			},
			{
				ComplaintID: "C-123j9ak280",
				Path:        "complaint-files/example2.jpg",
			},
			{
				ComplaintID: "C-123j9ak280",
				Path:        "complaint-files/example1.jpg",
			},
		}

		if err := db.CreateInBatches(&complaintFiles, len(complaintFiles)).Error; err != nil {
			panic(err)
		}
	}
}
