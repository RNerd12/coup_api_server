package engine

func resolveAction(game *Game, action Action) string {
	switch action.Type {
	case "income":
	case "foreign_aid":
	case "coup":
	case "tax":
	case "assassinate":
	case "exchange":
	case "steal":
	default:
		return "unknown action type"
	}
}
