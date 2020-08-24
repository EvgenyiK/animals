package app

import ("net/http"
u "github.com/EvgenyiK/animals/utils/utils"
)




var JwtAuthentication = func (next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		notAuth:= []string{"api/user/new", "api/user/login"}//List of endpoints that doesn't require auth
		requestPath:= r.URL.Path //current request path
		 //check if request does not need authentication, serve the request if it doesn't need it

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w,r)
				return
			}
		}
		
		response := make(map[string] interface{})
		tokenHeader:= r.Header.Get("Autorization") //Grab the token from the header

		if tokenheader == "" {//Token is missing, returns with error code 403 Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type","application/json")
			u.Responde(w, response)
			return
		}

	})
}