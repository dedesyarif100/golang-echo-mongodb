package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ScoresResult struct {
	ID			primitive.ObjectID	`json:"id" bson:"_id"`
	Student		string				`json:"student" bson:"student"`
	HomeWork	[]int				`json:"-" bson:"homework"`
	Quiz		[]int				`json:"-" bson:"quiz"`
	ExtraCredit	int					`json:"extracredit" bson:"extracredit"`
	TotalHomework	int				`json:"totalHomework" bson:"totalHomework"`
	TotalQuiz		int				`json:"totalQuiz" bson:"totalQuiz"`
	TotalScore		int				`json:"totalScore" bson:"totalScore"`
	PassingScores	int				`json:"passing_scores" bson:"passing_scores"`
	Score		ScoresEmbedded		`json:"embedded_score" bson:"embedded_score"`
}

type ScoresEmbedded struct {
	Uts			int		`json:"-" bson:"uts"`
	Uas			int		`json:"-" bson:"uas"`
	Scores		int		`json:"scores" bson:"scores"`
}
