package repository

// ReportRepository contains the functions of data logic for domain report
type ReportRepository interface {
}

type reportRepository struct {
}

// ReportInit initializes the data logic / repository for domain report
func ReportInit() ReportRepository {
	return reportRepository{}
}
