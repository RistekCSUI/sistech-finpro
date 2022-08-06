package depedencies

import (
	"context"
	"github.com/RistekCSUI/sistech-finpro/shared/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(env *config.EnvConfig) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(env.MongoUrl))
	if err != nil {
		panic(err)
	}

	return client.Database(env.MongoDatabase)
}
