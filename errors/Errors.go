package errors
import(
	"errors"
)
var(
	NotFoundError = errors.New("Does Not Exist")
	BadRequestError = errors.New("Bad Request")
	InternalServerError = errors.New("Internal Server Error")
	//Validation errors can still be introduced 
	//Un-Authorized Error when I will have roles in Fiber Mart considering User Panels which could add products
)