package dto

type DashboardSummaryResponse struct {
    TotalIncome   float64            `json:"total_income"`
    TotalExpense  float64            `json:"total_expense"`
    NetBalance    float64            `json:"net_balance"`
    CategoryTotals map[string]float64 `json:"category_totals"`
    RecentCount   int64              `json:"recent_count"`
}