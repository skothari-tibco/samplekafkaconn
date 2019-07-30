package main

import (
	_ "github.com/project-flogo/flow"
	_ "github.com/project-flogo/contrib/trigger/rest"
	_ "github.com/project-flogo/contrib/function/string"
	_ "github.com/project-flogo/kafkaconnection"
	_ "github.com/project-flogo/contrib/trigger/timer"
	_ "github.com/project-flogo/contrib/activity/kafka_shared"
)
