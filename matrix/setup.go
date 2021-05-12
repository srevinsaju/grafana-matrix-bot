package matrix

import (
	"github.com/withmandala/go-log"
	"maunium.net/go/mautrix"
	"os"
)


var logger = log.New(os.Stdout)


func Setup (c Config) *mautrix.Client {
	homeserver := c.Homeserver
	password := c.Password
	username := c.Username

	if username == "" && password == "" && homeserver == "" {
		// the user doesnt want to bridge to matrix network
		return nil
	}

	if username == "" || password == "" || homeserver == "" {
		// something information is not provided
		logger.Fatal("Invalid params for Matrix bot. Couldn't find either username, password or homeserver")
		return nil
	}

	client, err := mautrix.NewClient(homeserver, "", "")
	if err != nil {
		logger.Fatalf("Couldn't connect to matrix %s, %s", homeserver, err)
		return nil
	}

	_, err = client.Login(&mautrix.ReqLogin{
		Type:             "m.login.password",
		Identifier:       mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: username},
		Password:         password,
		StoreCredentials: true,
	})
	if err != nil {
		logger.Fatalf("Couldn't login to Matrix %s with %s, %s", homeserver, username, err)
		return nil
	}

	logger.Infof("Authorization successful as %s on %s", username, homeserver)



	return client

}