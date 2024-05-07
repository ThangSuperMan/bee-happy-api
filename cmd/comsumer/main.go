package comsumer

import (
	"fmt"

	AppSetting "github.com/thangsuperman/bee-happy/config"
	"github.com/thangsuperman/bee-happy/internal"
)

func main() {
	conn, err := internal.ConnectRabbitMQ(
		AppSetting.Envs.RabbitMQUsername,
		AppSetting.Envs.RabbitMQPassword,
		fmt.Sprintf("%s:%s", AppSetting.Envs.RabbitMQHost, AppSetting.Envs.RabbitMQPort),
		"customers")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}

	defer client.Close()
}
