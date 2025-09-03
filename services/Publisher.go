package services
import(
	"log"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/YogeshDharya/fiberBackend/models"
)

type RedisPublisher struct{
	client *redis.Client
	topic string
}

func NewRedisPublisher(client *redis.Client, topic string) *RedisPublisher{
	return &RedisPublisher{
		client: client,
		topic: topic,
	}
}

func (publisher *RedisPublisher) Publish(ctx context.Context, audit *models.Audit) error {
	payload, err := json.Marshal(audit)
	if err != nil {
		return err
	}
	err = publisher.client.Publish(ctx, publisher.topic, payload).Err()
	if err != nil {
		log.Printf("Publish Failed: ")
		return err
	}
	log.Printf("Published audit log : %s", payload)
	return nil
}