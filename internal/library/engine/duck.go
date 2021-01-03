package engine

import (
	"net/http"

	"Caesar/internal/relation"
)

type mvcRuler interface {
	AlphaFuzz() // 第一次扫描
	Aftermath() // 善后工作
}

type StandRuler interface {
	mvcRuler
	BetaFuzz() // 根据第一次扫描后的path再次进行扫描
}

// 鸭子类型
func StandFuzz(att StandRuler) {

	att.AlphaFuzz()
	att.BetaFuzz()
	att.Aftermath()

}

// mvc版的鸭子类型，区别在于mvc网站不需要二段扫描
func MVCFuzz(att mvcRuler) {

	att.AlphaFuzz()
	att.Aftermath()

}

// 简单工厂模式
func CreateFactory(status int, req RequestInfo, resp ResponseInfo, opts ServerOpt) StandRuler {
	switch status {

	case http.StatusOK:
		return &target200{
			request:  req,
			response: resp,
			opts:     opts,
			application: application{
				Store:   []relation.StorePath{},
				Results: []relation.ResultPtah{},
			},
		}

	case http.StatusFound, http.StatusMovedPermanently, http.StatusTemporaryRedirect:
		return &target30x{
			request:  req,
			response: resp,
			opts:     opts,
			application: application{
				Store:   []relation.StorePath{},
				Results: []relation.ResultPtah{},
			},
		}

	case http.StatusNotFound:
		return &target404{
			request:  req,
			response: resp,
			opts:     opts,
			application: application{
				Store:   []relation.StorePath{},
				Results: []relation.ResultPtah{},
			},
		}

	default:
		return &target404{
			request:  req,
			response: resp,
			opts:     opts,
			application: application{
				Store:   []relation.StorePath{},
				Results: []relation.ResultPtah{},
			},
		}

	}

}
