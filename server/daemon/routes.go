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
	m.Handle("/pong", PongHandler()).Methods("GET")

	return m
}
