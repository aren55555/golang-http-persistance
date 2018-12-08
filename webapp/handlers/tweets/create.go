package tweets

import "net/http"

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	// TODO: accept a POST request with a single tweet JSON body. Deserilize the
	//       JSON body, construct an INSERT SQL query, validate the SQL query
	//       exectuted successfully and return the newly created tweet with its
	//       server assigned ID, use the approraite 201 response status code.
	// Sample Request Body:
	//  {
	//      "body": "hello world"
	//  }
	// Sample Response Body:
	//  {
	//      "id": 1,
	//      "body": "hello world"
	//  }
}
