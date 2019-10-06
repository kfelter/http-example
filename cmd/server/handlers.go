package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) handleStatus() httprouter.Handle {
	//setup here

	// This gives us a closure environment in which our handler can operate
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		v, err := s.db.Get("statusCheckCount")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`STATUS: DOWN`))
			s.Logf("status check: %s", err)
			return
		}

		count, ok := v.(float64)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`STATUS: DOWN`))
			s.Logf("status check: count not ok: [%T] %v %v", v, v, ok)
			return
		}

		count++

		err = s.db.Set("statusCheckCount", count)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`STATUS: DOWN`))
			s.Logf("status check: %s", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`STATUS: UP`))
		s.Logf("status check: success #%d", int(count))
		return
	}
}
