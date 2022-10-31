package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	pgm "github.com/tkxkd0159/db/go/models"
)

func GetPgConn(dbname, host, user, pw string, port int, sslmode string) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s sslmode=%s", dbname, host, port, user, pw, sslmode))
	if err != nil {
		panic(err)
	}
	return db
}

func GetAllPilots(db boil.ContextExecutor, ctx context.Context) (pgm.PilotSlice, error) {
	pilots, err := pgm.Pilots().All(ctx, db)
	if err != nil {
		return nil, err
	}
	return pilots, nil
}

type BaseJet struct {
	Name, Color string
}

func GetJetWithCond(db boil.ContextExecutor, ctx context.Context, id, pilotID int) (*BaseJet, error) {
	j, err := pgm.Jets(
		qm.Select(pgm.JetColumns.Name, pgm.JetColumns.Color),
		qm.Where("id=? AND pilot_id=?", id, pilotID)).One(ctx, db)
	if err != nil {
		return nil, err
	}
	return &BaseJet{Name: j.Name, Color: j.Color}, nil
}

func GetJetWithRawQ(db boil.ContextExecutor, ctx context.Context) (pgm.JetSlice, error) {
	jets, err := pgm.Jets(qm.SQL("select * from jets LIMIT $1 OFFSET $2", 10, 0)).All(ctx, db)
	if err != nil {
		return nil, err
	}
	return jets, nil
}

func FindJet(db boil.ContextExecutor, ctx context.Context, id int) (*pgm.Jet, error) {
	jet, err := pgm.FindJet(ctx, db, id)
	if err != nil {
		return nil, err
	}
	return jet, nil
}

func EagerLoadWithRels(db boil.ContextExecutor, ctx context.Context) (pgm.JetSlice, error) {
	jets, err := pgm.Jets(qm.Load(qm.Rels(pgm.JetRels.Pilot))).All(ctx, db)
	if err != nil {
		return nil, err
	}
	return jets, nil
}

type JetInfo struct {
	AgeSum int `boil:"age_sum"`
	Count  int `boil:"jet_count"`
}

func BindJetInfo(db boil.ContextExecutor, ctx context.Context) (*JetInfo, error) {
	info := new(JetInfo)
	err := queries.Raw(`select sum(age) as "age_sum", count(*) as "jet_count" from jets`).Bind(ctx, db, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func InsertJet(db boil.ContextExecutor, ctx context.Context) {
	j := pgm.Jet{Name: "SuperSonic", PilotID: 1, Age: 30, Color: "Black"}
	err := j.Insert(ctx, db, boil.Infer())
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateJet(db boil.ContextExecutor, ctx context.Context) {
	rowsAff, err := pgm.Jets(qm.Where("id=?", 0)).UpdateAll(ctx, db, pgm.M{"id": 5})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rowsAff)
	}
}

func DeleteJet(db boil.ContextExecutor, ctx context.Context) {
	numrows, err := pgm.Jets(qm.Where("pilot_id=?", 10)).DeleteAll(ctx, db)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numrows)
	}
}
