package daemon

import (
	"github.com/gorilla/mux"
)

/*
	User
	-----

	/search
	Post request
	{
		medicine_name:
		lat:
		lon:
	}
	/user/upload_presciption

	/user/order
		{
			user_id
			medicines:[

			]
			store_id
			prescription:"blob_url"
		}

		returns
		{
			qr_code_url:""
		}

		Store
		------
		/store/orders
			{
				"orders":[

				]

			}

		/store/order/attend - Add medicines to users current course
		send a POST Request with body
		{
			"user_id":
			"medicines":[
				{
					name:
					couse:
				}
			]
		}

*/
func initRoutes() *mux.Router {
	m := mux.NewRouter()
	m.Handle("/search", SearchForMedicines()).Methods("POST")
	m.Handle("/user/upload_prescription", UploadPrescription()).Methods("POST")
	m.Handle("/user/order", UserOrder()).Methods("POST")
	m.Handle("/user/checkout_order/{user_id}", UserOrder()).Methods("POST")
	return m
}
