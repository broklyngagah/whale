package httpapi

const (
	/**
     * 获取积分任务列表
     */
	integral_integralList = "Integral/integralList"
	/**
	 * 消息中心
	 */
	index_noticesIndex = "index/noticesIndex"

	/**
	 * 公告列表
	 */
	index_systemMessageList = "index/systemMessageList"

	/**
	 * 公告详情
	 */
	index_systemMessageDetail = "index/systemMessageDetail"

	/**
	 * 获取banner接口
	 */
	index_slider = "index/slider"

	/**
	 * 获取文章列表
	 */
	article_articleList = "article/articleList"

	/**
	 * 文章分类下的子类（发现首页那边的文章列表 包括分类）
	 */
	category_category = "category/category"

	/**
	 * 评论点赞
	 */
	comment_commentDiggUp = "comment/commentDiggUp"

	/**
	 * 文章详情
	 */

	article_articledetail = "article/articledetail"

	/**
	 * 文章评论列表
	 */
	comment_commentList = "comment/commentList"

	/**
	 * 文章点赞
	 */
	article_articleDiggUp = "article/articleDiggUp"

	/**
	 * 文章分享
	 */
	article_articleShare = "article/articleShare"

	/**
	 * 发表帖子
	 */
	article_articlePost = "article/articlePost"

	/**
	 * 发表评论
	 */
	comment_commentPost = "comment/commentPost"

	/**
	 * 评论详情
	 */
	comment_commentdetail = "comment/commentdetail"

	/**
	 * 评论删除
	 */
	index_commentDel = "user/commentDel"

	/**
	 * 简要描述
	 */
	index_relativeindex = "index/relativeindex"

	/**
	 * 顶部分类
	 */
	category_categoryTop = "category/categoryTop"

	/**
	 * 发现首页右上角消息红点显示
	 */
	index_noticesDot = "index/noticesDot"

	/**
	 * 话题
	 */
	category_categoryDetail = "category/categoryDetail"

	/**
	 * 搜索文章/股票/用户
	 */
	index_indexsearch = "index/indexsearch"

	/**
	 * 发送短信
	 */
	reg_sendsms = "reg/sendsms"

	/**
	 * 登陆
	 */
	reg_apiStartCaptcha = "reg/apiStartCaptcha"

	/**
	 * 登陆
	 */
	reg_login = "reg/login"

	/**
	 * openid登录
	 */
	reg_loginByOpenId = "reg/loginByOpenId"

	/**
	 * 退出登陆
	 */
	reg_logOut = "reg/logOut"

	/**
	 * 消息-评论
	 */
	index_msgcomment = "index/msgcomment"

	/**
	 * 股票获取实时数据
	 */
	stock_real = "Stock/real"

	/**
	 * 股票获取K线数据
	 */
	stock_Kline = "Stock/Kline"

	/**
	 * 检查股票是否已经关注
	 */
	Mystock_checkUserSelect = "Mystock/checkUserSelect"

	/**
	 * 交易参数初始化
	 */
	Order_index = "Order/index"
	/**
	 * 个人首页
	 */
	user_index = "user/index"
	/**
	 * 我的关注
	 */
	user_myfollow = "user/myfollow"
	/**
	 * 用户设置
	 */
	user_userset = "user/userset"
	/**
	 * 获取七牛token
	 */
	Upload_qiniuToken = "Upload/qiniuToken"
	/**
	 * 加关注
	 */
	user_followAdd = "user/followAdd"

	/**
	 * 我的文章
	 */
	user_myArticle = "user/myArticle"

	/**
	 * 获取文章评论数目
	 */
	comment_commentNumByAid = "comment/commentNumByAid"

	/**
	 * 创建订单
	 */
	Order_createOrder = "Order/createOrder"
	/**
	 * 支付成功页
	 */
	Order_paySuccess = "Order/paySuccess"
	/**
	 * 钱包消费记录
	 */
	Order_orderDetail = "Order/orderDetail"
	/**
	 * 我的粉丝
	 */
	user_myfans = "user/myfans"
	/**
	 * 积分消费记录
	 */
	Integral_integralDetail = "Integral/integralDetail"
	/**
	 * 收藏-列表
	 */
	user_myCollect = "user/myCollect"
	/**
	 * 收藏-编辑
	 */
	user_myCollectDel = "user/myCollectDel"
	/**
	 * 用户反馈
	 */
	user_userFeedback = "user/userFeedback"
	/**
	 * 分享接口
	 */
	Share_getShareInfo = "Share/getShareInfo"
	/**
	 * 分享回调
	 */
	Share_shareReturn = "Share/shareReturn"
	/**
	 * 我的订阅
	 */
	user_mySubscribe = "user/mySubscribe"
	/**
	 * 自媒体认证、用户协议
	 */
	index_commParams = "index/commParams"
	/**
	 * 联系我们
	 */
	Other_contactUs = "Other/contactUs"
	/**
	 * app版本更新
	 */
	Other_appUpdate = "Other/appUpdate"
	/**
	 * 我的动态 对应大V那边的评论
	 */
	user_myDynamic = "user/myDynamic"

	/**
	 * 获取自选股模块首页三个指数
	 */
	Mystock_stockIndex = "Mystock/stockIndex"

	/**
	 * 我的自选股列表
	 */
	Mystock_myStock = "Mystock/myStock"

	/**
	 * 添加自选股
	 */
	Mystock_addStock = "Mystock/addStock"

	/**
	 * 自选股排序
	 */
	Mystock_stockSort = "Mystock/stockSort"

	/**
	 * 删除自选股
	 */
	Mystock_unStock = "Mystock/unStock"

	/**
	 * 批量/全部删除自选股
	 */
	Mystock_batchUnStock = "Mystock/batchUnStock"

	/**
	 * 批量添加自选股（未登录状态添加的自选股，登录后同步）
	 */
	Mystock_batchSyncStock = "Mystock/batchSyncStock"

	/**
	 * 股票分时数据
	 */
	Stock_trend = "Stock/trend"

	/**
	 * 5日分时图
	 */
	Stock_trend5day = "Stock/trend5day"

	/**
	 * 行情排序，可获取板块集
	 */
	Stock_sort = "Stock/sort"

	Stock_blockSort = "Stock/blockSort"

	/**
	 * 推荐订阅
	 */
	index_recSubscribe = "index/recSubscribe"

	/**
	 * 我的订阅
	 */
	index_myRelative = "index/myRelative"

	/**
	 * 收藏文章
	 */
	user_collectAdd = "user/collectAdd"

	/**
	 * 动态（最新/最热）
	 */
	article_articleListByHot = "article/articleListByHot"

	/**
	 * 支付
	 */
	Order_goPay = "Order/goPay"

	/**
	 * 积分购买
	 */
	Order_integralBuy = "Order/integralBuy"


	/**
	 * 余额购买
	 */
	Order_balanceBuy = "Order/balanceBuy"

	/**
	 * 搜索热词返回
	 */
	Index_hotWord = "Index/hotWord"

	/**
	 * 文章删除
	 */
	user_articleDel = "user/articleDel"

	/**
	 * 文章评论开启/关闭设置
	 */
	user_myArticleSet = "user/myArticleSet"

	/**
	 * 获取与我相关未读消息条数
	 */
	article_getRelateToMeCount = "article/getRelateToMeCount"

	/**
	 * 动态--与我相关
	 */
	article_relatedToMe = "article/relatedToMe"

	/**
	 * 3潜伏4中长线分类列表
	 */
	category_categoryList = "category/categoryList"

	/**
	 * 直播界面24小时人气排行
	 */
	chat_twentyHourlist = "chat/twentyHourlist"

	/**
	 * 播主榜单界面（播主排行列表）
	 */
	chat_rankList = "chat/rankList"

	/**
	 * 正在直播
	 */
	chat_nowList = "chat/nowList"

	/**
	 * 直播我的关注
	 */
	chat_myfollow = "chat/myfollow"
)
