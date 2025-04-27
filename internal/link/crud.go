package link

//	HANDLERS FOR Links

import (
	"log"
	"my_project/configs"

	"my_project/pkg/di"
	"my_project/pkg/middleware"
	"my_project/pkg/req"
	"my_project/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
	StatRepository di.IStatRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
	StatRepository di.IStatRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
		StatRepository: deps.StatRepository,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.Handle("GET /link", middleware.IsAuthed(handler.GetAll(), deps.Config))

}

// Принимает запрос на регистрацию и отправляет link обратно
func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			res.Json(w, "Cant Create Link", 401)
			return
		}

		link := NewLink(body.Url)

		existedLink, err := handler.LinkRepository.GetByHash(link.Hash)
		if existedLink != nil {
			link.Hash = string(GenerateHash(10))
		}

		// DATABASE HANDLER
		CreatedLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("Link created")
		res.Json(w, CreatedLink, 201)
		// DATABASE HANDLER

	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		log.Println("Update Handler Running")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, "Cant parse id to uint", http.StatusBadGateway)
			return
		}

		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			res.Json(w, "Cant Update Link", 401)
			return
		}

		emailM := r.Context().Value(middleware.ContextEmailKey).(string)

		log.Printf("Email from context  --> %s", emailM)

		// DATABASE HANDLER

		link, err := handler.LinkRepository.UpdateById(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		log.Println("Hash updated")
		res.Json(w, link, 200)
		// DATABASE HANDLER
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Cant parse id to uint", http.StatusBadGateway)
			return
		}

		// DATABASE HANDLER
		err = handler.LinkRepository.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = handler.LinkRepository.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), 200)
			return
		}
		log.Println(id, "deleted")
		res.Json(w, "deleted", 200)

		// DATABASE HANDLER

	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		// DATABASE HANDLER
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		// DATABASE HANDLER

		handler.StatRepository.AddClick((link.ID))

		log.Println("REDIRECTION...")
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}

}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			log.Println("Cant Atoi limit")
			http.Error(w, "Invalid limit type in query", 402)
			return
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			log.Println("Cant Atoi offset")
			http.Error(w, "Invalid offset type in query", 402)
			return
		}

		result := handler.LinkRepository.GetAll(limit, offset)
		counter := handler.LinkRepository.GetActive()

		data := GetLinksResponce{
			Links: result,
			Count: counter,
		}

		res.Json(w, data, 200)

	}

}
