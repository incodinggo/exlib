package searcher

import (
	"encoding/gob"
	"fmt"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"strings"
)

var (
	searcher = riot.Engine{}
)

type Opt struct {
	Using         int
	GseDict       string
	StopTokenFile string
}

/*
Init
******************************************************************************
	初始化引擎函数
*******************************************************************************/
func Init(opt Opt) {

	// 初始化
	gob.Register(ScoringFields{})
	fmt.Println("引擎开始初始化")
	searcher.Init(types.EngineOpts{
		Using:         opt.Using,
		GseDict:       opt.GseDict,
		StopTokenFile: opt.StopTokenFile,
		IndexerOpts: &types.IndexerOpts{
			IndexType: types.LocsIndex,
		},
		// 如果你希望使用持久存储，启用下面的选项
		// 默认使用leveldb持久化，如果你希望修改数据库类型
		// 请用 StoreEngine: " " 或者修改 Riot_Store_Engine 环境变量
		// UseStore: true,
		// StoreFolder: "weibo_search",
		// StoreEngine: "bg",
	})
	fmt.Println("引擎初始化完毕")
	//wbs = make(map[string]Weibo)
	defer searcher.Close()

}

func AddDocs(data []Attrib) {

	for _, v := range data {
		article.Set(v.Id, v)
		go func(v Attrib) {
			searcher.Index(v.Id, types.DocData{
				Content: v.Text,
				Labels:  v.Tags,
				Fields: ScoringFields{
					Timestamp: v.Timestamp,
					View:      v.View,
				},
			})
		}(v)

	}

	searcher.Flush()
}

func SearchByLabels(offset, rows int, labels ...string) (docs []Attrib, num int) {

	output := searcher.SearchDoc(types.SearchReq{
		Labels: labels,
		Logic: types.Logic{
			Must: true,
			Expr: types.Expr{Must: labels},
		},
		RankOpts: &types.RankOpts{
			ScoringCriteria: &ScoringCriteria{},
			OutputOffset:    offset,
			MaxOutputs:      rows,
		},
	})

	docs, num = formatOutput(output)

	return
}

func SearchByTags(offset, rows int, tags string) (docs []Attrib, num int) {
	labels := strings.Split(tags, ",")
	output := searcher.SearchDoc(types.SearchReq{
		Labels: labels,
		Logic: types.Logic{
			Must: true,
			Expr: types.Expr{Must: labels},
		},
		RankOpts: &types.RankOpts{
			ScoringCriteria: &ScoringCriteria{},
			OutputOffset:    offset,
			MaxOutputs:      rows,
		},
	})

	docs, num = formatOutput(output)
	return

}

func SearchByDoc(offset, rows int, keyword string) (docs []Attrib, num int) {

	output := searcher.SearchDoc(types.SearchReq{
		Text: keyword,
		RankOpts: &types.RankOpts{
			ScoringCriteria: &ScoringCriteria{},
			OutputOffset:    offset,
			MaxOutputs:      rows,
		},
	})
	docs, num = formatOutput(output)

	return
}

func formatOutput(output types.SearchDoc) (docs []Attrib, num int) {
	docs = []Attrib{}
	// 整理为输出格式
	for _, doc := range output.Docs {
		docs = append(docs, article.Get(doc.DocId))
	}
	num = output.NumDocs
	return
}
