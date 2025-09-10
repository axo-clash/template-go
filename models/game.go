package models

type GameWinner string

const (
	GameWinnerFirst  GameWinner = "FIRST_BOT"
	GameWinnerSecond GameWinner = "SECOND_BOT"
	GameWinnerDraw   GameWinner = "DRAW"
)

type GameEndReason string

const (
	EndReasonFirstBotLife             GameEndReason = "FIRST_BOT_LIFE"
	EndReasonSecondBotLife            GameEndReason = "SECOND_BOT_LIFE"
	EndReasonFirstBotForbiddenAction  GameEndReason = "FIRST_BOT_FORBIDDEN_ACTION"
	EndReasonSecondBotForbiddenAction GameEndReason = "SECOND_BOT_FORBIDDEN_ACTION"
	EndReasonGameTurn                 GameEndReason = "GAME_TURN"
)

type GameFailureReason string

const (
	FailureBotConnectionError      GameFailureReason = "BOT_CONNECTION_ERROR"
	FailureBotInfoRetrievalError   GameFailureReason = "BOT_INFO_RETRIEVAL_ERROR"
	FailureGameExecutionError      GameFailureReason = "GAME_EXECUTION_ERROR"
	FailureInvalidBoardState       GameFailureReason = "INVALID_BOARD_STATE"
	FailureBotCommunicationTimeout GameFailureReason = "BOT_COMMUNICATION_TIMEOUT"
	FailureJSONProcessingError     GameFailureReason = "JSON_PROCESSING_ERROR"
	FailureUnknownError            GameFailureReason = "UNKNOWN_ERROR"
)

type GameStatus string

const (
	GameStatusStarting GameStatus = "STARTING"
	GameStatusRunning  GameStatus = "RUNNING"
	GameStatusPaused   GameStatus = "PAUSED"
	GameStatusFinished GameStatus = "FINISHED"
	GameStatusFailed   GameStatus = "FAILED"
)

type GameDTO struct {
	ID            int64             `json:"id"`
	FirstBot      *BotDTO           `json:"firstBot"`
	SecondBot     *BotDTO           `json:"secondBot"`
	Board         [][]CaseDTO       `json:"board"`
	CurrentTurn   int32             `json:"currentTurn"`
	Actions       []ActionDTO       `json:"actions"`
	Winner        GameWinner        `json:"winner"`
	EndReason     GameEndReason     `json:"endReason"`
	FailureReason GameFailureReason `json:"failureReason"`
	Status        GameStatus        `json:"status"`
}
