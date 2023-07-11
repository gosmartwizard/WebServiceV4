package product

import "github.com/gosmartwizard/WebServiceV4/business/data/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByProdID, order.ASC)

// Set of fields that the results can be ordered by. These are the names
// that should be used by the application layer.
const (
	OrderByProdID   = "productid"
	OrderByName     = "name"
	OrderByCost     = "cost"
	OrderByQuantity = "quantity"
	OrderBySold     = "sold"
	OrderByRevenue  = "revenue"
	OrderByUserID   = "userid"
)
