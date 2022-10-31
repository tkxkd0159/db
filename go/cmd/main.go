package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/tkxkd0159/db/go/pg"
)

func main() {
	db := pg.GetPgConn("ljs", "127.0.0.1", "writer", "secret", 6666, "disable")
	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel

	res, err := pg.GetAllPilots(db, ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*res[0])

	res2, err := pg.GetJetWithCond(db, ctx, 1, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*res2)
	}

	res3, err := pg.GetJetWithRawQ(db, ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, r := range res3 {
			fmt.Println("Raw Query : ", *r)
		}
	}

	res4, err := pg.FindJet(db, ctx, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*res4)
	}

	res5, err := pg.EagerLoadWithRels(db, ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*res5[0].R.GetPilot())
	}

	res6, err := pg.BindJetInfo(db, ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*res6)
	}

	pg.DeleteJet(db, ctx)
	pg.InsertJet(db, ctx)
	pg.UpdateJet(db, ctx)
}
