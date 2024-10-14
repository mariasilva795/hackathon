package models

import (
	"time"
)

type EmotionalDailyLog struct {
	Id             string    `json:"id"`
	Date           time.Time `json:"date"`
	Status         string    `json:"status"`         //Estado emocional del usuario (ej. "Feliz", "Ansioso", "Triste").
	ActivitiesDone []string  `json:"activitiesDone"` //Lista de actividades realizadas (ej. ["Trabajo", "Ejercicio", "Tiempo con amigos"]).
	LevelEnergy    int       `json:"levelEnergy"`    //Nivel de energía reportado por el usuario (valor entre 1 y 10).
	Category       string    `json:"category"`       //Categoría asignada al registro, puede ser "Ingreso" o "Egreso" (o definida por la AI).

	Tags      []string  `json:"tags"` //Etiquetas personalizadas que el usuario agrega al registro (ej. ["Familia", "Trabajo"]).
	Photos    string    `json:"photos"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId    string    `json:"userId"`
}
