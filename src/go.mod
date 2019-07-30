module main

go 1.12

require (
	github.com/project-flogo/contrib/activity/kafka_shared v0.0.0
	github.com/project-flogo/contrib/function/string v0.9.0
	github.com/project-flogo/contrib/trigger/rest v0.9.0
	github.com/project-flogo/contrib/trigger/timer v0.9.0
	github.com/project-flogo/core v0.9.3-0.20190726142805-ef75331bd75a
	github.com/project-flogo/flow v0.9.2
	github.com/project-flogo/kafkaconnection v0.0.0
)

replace github.com/project-flogo/kafkaconnection => ../kafkaconnection

replace github.com/project-flogo/core v0.9.3-0.20190726142805-ef75331bd75a => /Users/skothari-tibco/Documents/core

replace github.com/project-flogo/contrib/activity/kafka_shared => ../kafka_shared
