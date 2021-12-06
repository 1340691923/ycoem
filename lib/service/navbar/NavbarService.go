package seo

import (
	"encoding/json"
)

type NavbarService struct {

}

func(this *NavbarService)Get()(navbar map[string]interface{},err error){
	jsonStr := `{
"产品":{"url":"123",
"node":[{"url":"123","text":"KTV泛娱乐"},{"url":"123","text":"家用娱乐"},{"url":"123","text":"H3酒店产品"},{"url":"123","text":"歌咏亭产品"}]
},
"解决方案":{"url":"123",
"node":[{"url":"123","text":"KTV娱乐解决方案"},{"url":"123","text":"家用解决方案"},{"url":"123","text":"酒店解决方案"},{"url":"123","text":"歌咏亭解决方案"}]
},"案例":{"url":"123",
"node":[{"url":"123","text":"连锁KTV"},{"url":"123","text":"家用娱乐"},{"url":"123","text":"酒店水疗"}]
},"案例":{"url":"123",
"node":[{"url":"123","text":"连锁KTV"},{"url":"123","text":"家用娱乐"},{"url":"123","text":"酒店水疗"}]
},"服务支持":{"url":"123",
"node":[{"url":"123","text":"我们承诺"},{"url":"123","text":"解决手册"},{"url":"123","text":"产品资料"},{"url":"123","text":"专业解答"}]
},"关于音创":{"url":"123",
"node":[{"url":"123","text":"音创介绍"},{"url":"123","text":"音创动态"},{"url":"123","text":"展会活动"},{"url":"123","text":"诚邀合作"}]
}
}`

	err = json.Unmarshal([]byte(jsonStr),&navbar)
	return
}
