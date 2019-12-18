package constant

const COMMAND_READY string = "COMMAND_READY"
const EVENT_VALIDATION_TOPIC string = "EVENT_VALIDATION_TOPIC"
const COMMAND_VALIDATION_FUNCTIONS string = "COMMAND_VALIDATION_FUNCTIONS"
const EVENT_VALIDATION_FUNCTIONS string = "EVENT_VALIDATION_FUNCTIONS"
const COMMAND_VALIDATION_FUNCTIONS_REPLY string = "COMMAND_VALIDATION_FUNCTIONS_REPLY"
const WORKER_SERVICE_CLASS_CAPTURE string = "WORKER_SERVICE_CLASS_CAPTURE"

type State int
const (
    ONGOING State = iota
    SUCCES
    FAIL 
)