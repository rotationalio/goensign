package ensign

import (
	"context"

	api "github.com/rotationalio/go-ensign/api/v1beta1"
	"github.com/rotationalio/go-ensign/stream"
	"google.golang.org/grpc"
)

func (c *Client) Publish(topic string, events ...*Event) (err error) {
	// Ensure the publisher is open before publishing
	c.openPub.Do(func() {
		c.pub, err = stream.NewPublisher(c, c.copts...)
	})

	// If the publisher could not be opened, return an error
	if err != nil {
		return err
	}

	// Attempt to send all events to the server, stopping on the first error.
	for _, event := range events {
		if event.pub, err = c.pub.Publish(topic, event.toPB()); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) PublishStream(ctx context.Context, opts ...grpc.CallOption) (api.Ensign_PublishClient, error) {
	return c.api.Publish(ctx, opts...)
}
