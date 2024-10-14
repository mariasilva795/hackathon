package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mariasilva795/go-api-rest/helpers/auth"
	"github.com/mariasilva795/go-api-rest/models"
	"github.com/mariasilva795/go-api-rest/repository"
	"github.com/mariasilva795/go-api-rest/server"
	"github.com/segmentio/ksuid"
)

type UpsertPostRequest struct {
	PostContent string `json:"postContent"`
}

// type PostResponse struct {
// 	Id          string `json:"id"`
// 	PostContent string `json:"postContent"`
// }

type PostResponse struct {
	Id             string   `json:"id"`
	Status         string   `json:"status"`         //Estado emocional del usuario (ej. "Feliz", "Ansioso", "Triste").
	ActivitiesDone []string `json:"activitiesDone"` //Lista de actividades realizadas (ej. ["Trabajo", "Ejercicio", "Tiempo con amigos"]).
	LevelEnergy    int      `json:"levelEnergy"`    //Nivel de energía reportado por el usuario (valor entre 1 y 10).
	Category       string   `json:"category"`       //Categoría asignada al registro, puede ser "Ingreso" o "Egreso" (o definida por la AI).
	Tags           []string `json:"tags"`           //Etiquetas personalizadas que el usuario agrega al registro (ej. ["Familia", "Trabajo"]).
	Photos         string   `json:"photos"`
}

type PostDeletedResponse struct {
	Message string `json:"message"`
}

func InsertEmotionalDailyLogHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		claimsUserId, err := auth.ValidateToken(s, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var postRequest = UpsertPostRequest{}
		err = json.NewDecoder(r.Body).Decode(&postRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var emotionLog = models.EmotionalDailyLog{
			Id:             id.String(),
			Date:           time.Now(),                                             // Capture current date and time
			Status:         "Feliz",                                                // Emotional status of the user
			ActivitiesDone: []string{"Trabajo", "Meditación", "Tiempo con amigos"}, // Activities done during the day
			LevelEnergy:    8,                                                      // Energy level on a scale of 1 to 10
			Category:       "Ingreso",                                              // Whether this is an 'Ingreso' or 'Egreso'
			Photos:         "https://example.com/photo1.jpg",                       // URL of the photo linked to the entry
			Tags:           []string{"Bienestar", "Productividad"},                 // Tags added to the entry
			CreatedAt:      time.Now(),                                             // When the log was created
			UpdatedAt:      time.Now(),                                             // When the log was last updated
			UserId:         claimsUserId,
		}

		err = repository.InsertEmotionalDailyLog(r.Context(), &emotionLog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostResponse{
			Id:             emotionLog.Id,
			Status:         emotionLog.Status,
			ActivitiesDone: emotionLog.ActivitiesDone,
			LevelEnergy:    emotionLog.LevelEnergy,
			Category:       emotionLog.Category,
			Tags:           emotionLog.Tags,
			Photos:         emotionLog.Photos,
		})
	}
}

func GetEmotionalDailyLogByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		post, err := repository.GetEmotionalDailyLogById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}
