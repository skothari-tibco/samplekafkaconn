package kafka_shared

import (
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/connection"
)

type Settings struct {
	Topic string `md:"topic"`
}
type Input struct {
	Message    string             `md:"message,required"` // The message to send
	Connection connection.Manager `md:"connection,required"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message":    i.Message,
		"connection": i.Connection,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Message, err = coerce.ToString(values["message"])
	i.Connection, err = coerce.ToConnection(values["connection"])
	return err
}

type Output struct {
	Partition int32 `md:"partition"` // Documents the partition that the message was placed on
	OffSet    int64 `md:"offset"`    // Documents the offset for the message
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"partition": o.Partition,
		"offset":    o.OffSet,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Partition, err = coerce.ToInt32(values["partition"])
	if err != nil {
		return err
	}

	o.OffSet, err = coerce.ToInt64(values["offset"])
	if err != nil {
		return err
	}

	return nil
}
