package indicator

import (
	"crypto-market/model"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/graphql-go/graphql"
)

//
type Indicator struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

var (
	//
	Indicators []Indicator = []Indicator{
		Indicator{
			Name: "MA日线",
			Description: `

			MA指标:英文名称是Moving average，叫做移动平均线指标。MA具有趋势的特性，
			其取值是把某段时间的币价加以平均，再依据这个平均值作出的趋势线，通常比较
			平稳且不像日K线那样有剧烈波动，是一种趋势追踪工具，便于识别趋势已经终结
			或者反转，新的趋势是否正在形成。 MA均线还可以组成组合指标来看，比如说，
			当5，10，30，60日均线从上到下依次出现。我们常称为多头排列，暗示出现上
			涨行情的可能性更大。 反之这些均线从下到上依次出现，则成为空头排列。空
			头排列方式出现下跌行情的可能性更大。
	
	
			使用技巧
			重点参考著名的葛兰碧法则	
			1.平均线从下降逐渐转为走平，而价格从下方突破平均线，为买进信号。
			2.价格虽然跌破平均线，但是又立刻回升到平均线上，此时平均线仍然持续上升，仍为买进信号。
			3.价格趋势走在平均线上，价格下跌并未跌破平均线且立刻反转上升，也是买进信号。
			4.价格突然暴跌，跌破平均线，且远离平均线，则有可能反弹上升，也为买进时机。
			5.平均线从上升逐渐转为盘局或下跌，而价格向下跌破平均线，为卖出信号。
			6.价格虽然向上突破平均线，但是又立刻回跌至平均线下，此时平均线仍然持续的下降，仍为卖出信号。
			7.价格趋势走在平均线下，价格上升并未突破平均线且立刻反转下跌，也是卖出信号。
			8.价格突然暴涨，突破平均线，且远离平均线，则有可能反弹回跌，也为卖出时机。
	
			`,
		},
		Indicator{
			Name: "MACD",
			Description: `
				
			MACD指标:英文是Moving Average Convergence and Divergence，叫做异同移动平均线，
			是从双指数移动平均线发展而来的。由快的指数移动平均线（EMA12）减去慢的指数移动平均
			线（EMA26）得到快线DIF，再用2×（快线DIF-DIF的9日加权移动均线DEA）得到MACD柱。
			MACD的变化代表着市场趋势的变化，不同K线级别的MACD代表当前级别周期中的买卖趋势。
			
			使用技巧
			1.DIF、DEA均为正，DIF向上突破DEA，买入信号参考。
			2.DIF、DEA均为负，DIF向下跌破DEA，卖出信号参考。
			3.DIF线与K线发生背离，行情可能出现反转信号。
			4.DIF、DEA的值从正数变成负数，或者从负数变成正数并不是交易信号，因为它们落后于市场。
		
			`,
		},
		Indicator{
			Name: "RSI",
			Description: `
				
			RSI指标，英文是Relatives Strength Index，中文通常译作相对强弱指数，是据一定时期
			内上涨点数和涨跌点数之和的比率制作出的一种技术曲线，RSI可以近似代表买卖双方的力量对比。

			使用技巧
			1.RSI值将0到100之间分成了从"极弱"、"弱""强"到"极强"四个区域。"强"和"弱"以50作为分界线,
			但"极弱"和"弱"之间以及"强"和"极强"之间的界限则要随着RSI参数的变化而变化。
			2.不同的参数,其区域的划分就不同。一般而言,参数越大,分界线离中心线50就越近,离100和0就越远。
			不过一般都应落在15、30到70、85的区间内。
			3.RSI值如果超过50,表明市场进入强市,可以考虑买入,但是如果继续进入"极强"区,就要考虑物极必反,准备卖出了。
			4.同理RSI值在50以下也是如此,如果进入了"极弱"区,则表示超卖,应该伺机买入。

			`,
		},
		Indicator{
			Name: "W&R威谦指标",
			Description: `
			
			W&R(Williams Overbought/Oversold Index)，主要依据最低价和收盘价之间的关系，来判断市
			场的超买超卖现象，预测币价中短期的走势。它主要是利用振荡点来反映市场的超买超卖行为，分析
			多空双方力量的对比，从而提出有效的信号来研判市场中短期行为的走势。

			使用技巧
			1.当W&R高于80，即处于超卖状态，行情即将见底，应当考虑买进。
			2.当W&R低于20，即处于超买状态，行情即将见顶，应当考虑卖出。
			3.在W&amp;R进入高位后，一般要回头，如果币价继续上升就产生了背离，是卖出信号。
			4.在W&amp;R进入低位后，一般要反弹，如果币价继续下降就产生了背离。
			5.W&R连续几次撞顶（底），局部形成双重或多重顶（底），是卖出（买进）的信号。
			6.同时，使用过程中应该注意与其他技术指标相互配合。在盘整的过程中，W&R的准确性较
			高，而在上升或下降趋势当中，却不能只以W&R超买超卖信号作为行情判断的依据。

			`,
		},
		Indicator{
			Name: "DMI指标",
			Description: `
				
			DMI指标又叫动向指标或趋向指标，通过分析币价价格在涨跌过程中买卖双方力量均
			衡点的变化情况，即多空双方的力量的变化受价格波动的影响而发生由均衡到失衡的
			循环过程，从而提供对趋势判断依据的一种技术指标。

			使用技巧
			多空指标包括(+DI多方、-DI空方)
			1.+DI在-DI上方,币价行情以上涨为主;+DI在-DI下方，币价行情以下跌为主。
			2.在币价价格上涨行情中，当+DI向上交叉-DI，是买进信号，相反,当+DI向下交叉-DI，
			是卖出信号。
			3.-DI从20以下上升到50以上,币价价格很有可能会有一波中级下跌行情。
			4.+DI从20以下上升到50以上,币价价格很有可能会有一波中级上涨行情。
			5.+DI和-DI以20为基准线上下波动时，该币价多空双方拉锯战,币价价格以箱体整理为主。
			
			`,
		},
		Indicator{
			Name: "OBV指标",
			Description: `

			能量潮指标（On Balance Volume，OBV）将市场的人气——成交量与价格的关系数字化，
			直观化，以市场的成交量变化来衡量市场的推动力，从而研判价格的走势。利用OBV可以
			验证当前价格走势的可靠性，并可以得到趋势可能反转的信号。比起单独使用成交量来，
			OBV看得更清楚。
	
			使用技巧
			1、当币价升而OBV线下降，表示买盘无力，币价可能会回跌。
			2、币价下降时而OBV线上升，表示买盘旺盛，币价可能会止跌回升。
			3、OBV线缓慢上升，表示买气逐渐加强，为买进信号。
			4、OBV线急速上升时，表示力量将用尽为卖出信号。
			5、OBV线从正的累积数转为负数时，为下跌趋势，应该卖出。反之，OBV线从负的累积数转为正数时，应该买进。
			6、OBV线最大的用处，在于观察盘面整理后，何时会脱离盘局以及突破后的未来走势，OBV线变动方向是重要
			参考指数，其具体的数值并无实际意义。
			7、OBV线对双重顶第二个高峰的确定有较为标准的显示，当币价自双重顶第一个高峰下跌又再次回升时，如果OBV线能够随
			币价趋势同步上升且价量配合，则可持续多头市场并出现更高峰。相反，当币价再次回升时OBV线未能同步配合，却见下降，
			则可能形成第二个顶峰，完成双重顶的形态，导致币价反转下跌。
	
			`,
		},
		Indicator{
			Name: "CCI指标",
			Description: `
				
			CCI指标(Commodity Channel Index)，又叫顺势指标，专门衡量价格常态分布范围，强调价格平均绝对偏差在市场技
			术分析中的重要性，属于超卖超卖类指标的一种。

			使用技巧
			1.当CCI指标曲线在+100线～-100线的常态区间里运行时,CCI指标参考意义不大，可以用KDJ等其它技术指标进行研判。
			2.当CCI指标曲线从上向下突破+100线而重新进入常态区间时，表明市场价格的上涨阶段可能结束，将进入一个比较长
			时间的震荡整理阶段，应及时平多做空。
			3.当CCI指标曲线从上向下突破-100线而进入另一个非常态区间（超卖区）时，表明市场价格的弱势状态已经形成，将
			进入一个比较长的寻底过程，可以持有空单等待更高利润。如果CCI指标曲线在超卖区运行了相当长的一段时间后开始
			掉头向上，表明价格的短期底部初步探明，可以少量建仓。CCI指标曲线在超卖区运行的时间越长，确认短期的底部的
			准确度越高。
			4.CCI指标曲线从下向上突破-100线而重新进入常态区间时，表明市场价格的探底阶段可能结束，有可能进入一个盘
			整阶段，可以逢低少量做多。
			5.CCI指标曲线从下向上突破+100线而进入非常态区间(超买区)时，表明市场价格已经脱离常态而进入强势状态，如
			果伴随较大的市场交投，应及时介入成功率将很大。
			6.CCI指标曲线从下向上突破+100线而进入非常态区间(超买区)后，只要CCI指标曲线一直朝上运行，表明价格依然
			保持强势可以继续持有待涨。但是，如果在远离+100线的地方开始掉头向下时，则表明市场价格的强势状态将可能
			难以维持，涨势可能转弱，应考虑卖出。如果前期的短期涨幅过高同时价格回落时交投活跃，则应该果断逢高卖出或做空。
			
			`,
		},

		Indicator{
			Name: "ORC指标",
			Description: `
			
			变动率指标（ROC），是以当日的收盘价和N天前的收盘价比较，通过计算币价某一段时间内收盘价变动的比例，应
			用价格的移动比较来测量价位动量，达到事先探测币价买卖供需力量的强弱，进而分析币价的趋势及其是否有转势
			的意愿，属于反趋势指标之一。

			使用技巧
			1.在趋势明显的市场中，当ROC由上往下跌破0时，为卖出时机；当ROC由下往上穿破0时，为买进时机。
			2.在趋势不明显的平衡震荡行情中，当ROC由上往下跌破MAROC时，为卖出时机；而当ROC由下往上穿破MAROC时，为买进时机。
			3.当币价创新低点，而ROC未配合下降，意味下跌动力减弱，此背离现象，应逢低承接；当币价创新高点，而ROC未配合
			上升，意味上涨动力减弱，此背离现象，应慎防币价反转而下。
			4.若币价与ROC在低水平同步上升，显示短期趋向正常或短期会有币价反弹现象；若币价与ROC在高水平同步下降，
			显示短期趋向正常或短期会有币价回落现
			5.ROC波动于“常态范围”内，上升至第一条超买线时，应卖出币价；下降至第一条超卖线时，应买进币价。	
			6.ROC向上突破第一条超买线后，指标继续朝第二条超买线涨升的可能性很大，指标碰触第二条超买线时，涨势多半将结束。		
			7.ROC向下跌破第一条超卖线后，指标继续朝第二条超卖线下跌的可能性很大，指标碰触第二条超卖线时，跌势多半将停止。
			8.ROC向上穿越第三条超买线时，属于疯狂性多头行情，应尽量不轻易卖出持股。
			9.ROC向下穿越第三条超卖线时，属于崩溃性空头行情，应克制不轻易买进币价。

			`,
		},

		Indicator{
			Name: "SAR指标",
			Description: `
			
			抛物线指标（SAR）也称为停损点转向指标，这种指标与移动平均线的原理颇为相似，属于价格与时间并重的分析工具。

			使用技巧
			1.当币价上涨时，SAR的红色圆圈位于币价的下方，当该股的收盘价向下跌破SAR时，则应立即停损卖出。
			2.当币价下跌时，SAR的绿色圆圈位于币价的上方，当收盘价向上突破SAR时，可以重新买回。
			3.当币价在SAR的红色圆圈之上时，如果预见CR出现四条线集于一点的信号时，是比较难得的短线主升加速信号，
			应该加大注意力度。
			4.当币价在SAR的绿色曲线之下时，表明当前是空头市场，应离场观望。特别是在下跌过程中，虽然有时也出现
			绿圆圈翻红圆圈的现象，但如果只出现3个以下的红圆圈时，则又不能构成单一的买进信号，必须配合其它的技术
			指标来判断。

			`,
		},

		Indicator{
			Name: "BOLL指标",
			Description: `

			布林线指标，即BOLL指标，其英文全称是“Bollinger Bands”，通过计算价格的“标准差”，再求价格的“信赖区间”，
			是一个路径型指标，该指标在图形上画出三条线，上下两条线可以分别看成是价格的压力线和支撑线，中间为价格平均线。
			价格波动在上限和下限的区间之内，价格涨跌幅度加大时，袋装区会变宽，涨跌幅度狭小盘整时，带状区会变窄。
			价格超越上限时，代表超买，价格超越下限时，代表超卖。

			使用技巧
			1.当价格运行在布林通道的中轨和上轨之间的区域时，只要不破中轨，说明市场处于多头行情中，只考虑逢低买进，不考虑做空。
			2.在中轨和下轨之间时，只要不破中轨，说明是空头市场，交易策略是逢高卖出，不考虑买进。
			3.当市场价格沿着布林通道上轨运行时，说明市场是单边上涨行情，持有的多单要守住，只要价格不脱离上轨区域就耐心持有。
			4.沿着下轨运行时，说明市场目前为单边下跌行情，一般为一波快速下跌行情，持有的空单，只要价格不脱离下轨区域就耐心持有。
			5.当价格运行在中轨区域时，说明市场目前为盘整震荡行情，对趋势交易者来说，这是最容易赔钱的一种行情，应回避，空仓观望为上。
			6.布林通道的缩口状态。价格在中轨附近震荡，上下轨逐渐缩口，此是大行情来临的预兆，应空仓观望，等待时机。
			7.通道缩口后的突然扩张状态。意味着一波爆发性行情来临，此后，行情很可能走单边，可以积极调整建仓，顺势而为。
			8.当布林通道缩口后，在一波大行情来临之前，往往会出现假突破行情，这是主力的陷阱，应提高警惕，可以通过调整仓位化解。
			9.布林通道的时间周期应以周线为主，在单边行情时，所持仓单已有高额利润，为防止大的回调，可以参考日线布林通道的原则出局。


			`,
		},

		Indicator{
			Name: "KDJ指标",
			Description: `
			
			随机指标KDJ是以最高价、最低价及收盘价为基本数据进行计算，得出的K值、D值和J值分别在指标的坐标上形成的一个点，
			连接无数个这样的点位，就形成一个完整的、能反映价格波动趋势的KDJ指标。它主要是利用价格波动的真实波幅来反映
			价格走势的强弱和超买超卖现象，是在价格尚未上升或下降之前发出买卖信号的一种技术工具。

			使用技巧
			1.K与D值永远介于0到100之间。D大于80时，行情呈现超买现象。D小于20时，行情呈现超卖现象。
			2.上涨趋势中，K值小于D值，K线向上突破D线时，为买进信号。下跌趋势中，K值大于D值，K线向下跌破D线时，为卖出信号。
			3.KD指标不仅能反映出市场的超买超卖程度，还能通过交叉突破发出买卖信号。
			4.KD指标不适于发行量小、交易不活跃的币种。
			5.当随机指标与币价出现背离时，一般为转势的信号。
			6.K值和D值上升或者下跌的速度减弱，倾斜度趋于平缓是短期转势的预警信号。

			`,
		},
	}
)

//
func GetIndicators(req *model.IndicatorReq) []Indicator {
	if req != nil {
		var indicators []Indicator
		if req.Name != "" {
			for _, indicator := range Indicators {
				if strings.Contains(indicator.Name, req.Name) {
					switch req.Field {
					case "name":
						indicator.Description = ""
					case "description":
						indicator.Name = ""
					}
					indicators = append(indicators, indicator)
				}
			}
		}
		return indicators
	}
	return Indicators
}

//
func GraphQL() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
