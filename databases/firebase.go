package databases

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/mariasilva795/go-api-rest/models"
	"google.golang.org/api/option"
)

type FirestoreRepository struct {
	client *firestore.Client
}

func NewFirestoreRepository(urlDocument string) (*FirestoreRepository, error) {
	opt := option.WithCredentialsFile(urlDocument)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}
	return &FirestoreRepository{client}, nil // Return the FirestoreRepository
}

func (repo *FirestoreRepository) InsertUser(ctx context.Context, user *models.User) error {
	docRef := repo.client.Collection("emotional_banck").Doc("44SR9J4VS8aowdbcVmUO6").Collection("users").Doc(user.Email)
	_, err := docRef.Set(ctx, user)
	return err
}

func (repo *FirestoreRepository) Close() error {
	return nil
}

func (repo *FirestoreRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	query := repo.client.Collection("emotional_banck").Doc("44SR9J4VS8aowdbcVmUO6").Collection("users")

	docs, err := query.Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var user models.User
		if err := doc.DataTo(&user); err != nil {
			return nil, err
		}
		fmt.Println(user)

		if user.Id == id {
			return &user, nil
		}
	}

	return nil, nil
}

func (repo *FirestoreRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	docRef := repo.client.Collection("emotional_banck").Doc("44SR9J4VS8aowdbcVmUO6").Collection("users").Doc(email)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}
	if !doc.Exists() {
		return nil, nil
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

//EmotionalDailyLog

func (repo *FirestoreRepository) InsertEmotionalDailyLog(ctx context.Context, emotion *models.EmotionalDailyLog) error {
	docRef := repo.client.Collection("emotional_banck").Doc("44SR9J4VS8aowdbcVmUO6").Collection("emotionalDailyLog").NewDoc()
	_, err := docRef.Set(ctx, emotion)
	return err
}

func (repo *FirestoreRepository) GetEmotionalDailyLogById(ctx context.Context, id string) (*models.EmotionalDailyLog, error) {

	query := repo.client.Collection("emotional_banck").Doc("44SR9J4VS8aowdbcVmUO6").Collection("emotionalDailyLog")

	docs, err := query.Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var log models.EmotionalDailyLog
		if err := doc.DataTo(&log); err != nil {
			return nil, err
		}
		if log.Id == id {
			return &log, nil
		}
	}

	return nil, nil
}

func (repo *FirestoreRepository) UpdateEmotionalDailyLog(ctx context.Context, post *models.EmotionalDailyLog, userId string) error {

	query := repo.client.Collection("emotional_banck").
		Doc("44SR9J4VS8aowdbcVmUO6").
		Collection("emotionalDailyLog").
		Where("UserId", "==", userId).
		Where("Id", "==", post.Id)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return err
	}

	if len(docs) == 0 {
		return nil
	}

	doc := docs[0]

	_, err = doc.Ref.Set(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FirestoreRepository) DeleteEmotionalDailyLog(ctx context.Context, id string, userId string) error {
	query := repo.client.Collection("emotional_banck").
		Doc("44SR9J4VS8aowdbcVmUO6").
		Collection("emotionalDailyLog").
		Where("UserId", "==", userId).
		Where("Id", "==", id)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return err
	}

	if len(docs) == 0 {
		return nil
	}

	doc := docs[0]

	_, err = doc.Ref.Delete(ctx)

	return err
}

func (repo *FirestoreRepository) ListEmotionalDailyLogs(ctx context.Context) ([]*models.EmotionalDailyLog, error) {
	query := repo.client.Collection("emotional_banck").
		Doc("44SR9J4VS8aowdbcVmUO6").
		Collection("emotionalDailyLog")

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var logs []*models.EmotionalDailyLog

	for _, doc := range docs {
		var log models.EmotionalDailyLog

		if err := doc.DataTo(&log); err != nil {
			return nil, err
		}

		logs = append(logs, &log)
	}

	return logs, nil
}
