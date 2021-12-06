package response

import model "ycoem/lib/models"

type ServiceDep struct {
	Area    string `json:"area"`
	Manager string `json:"manager"`
	Support string `json:"support"`
}

type Solution struct {
	Solut model.SolutModel  `json:"solut"`
	Fatd  []model.FatdModel `json:"fatd"`
}

type ShowContacts struct {
	DeptInfos []string `json:"dept_infos"`
	Address   string   `json:"address"`
}

type Lunbo struct {
	Img string `json:"img"`
	Url string `json:"url"`
}

type LunboTable struct {
	Img  string `json:"img"`
	Name string `json:"name"`
}

type FirmModel struct {
	ID         int      `json:"id" db:"id"`
	Title      string   `json:"title" db:"title"`
	Img        []string `json:"img" db:"img"`
	WxLink     string   `json:"wx_link" db:"wx_link"`
	Content    string   `json:"content" db:"content"`
	Name       string   `json:"name" db:"name"`
	CreateTime string   `json:"create_time" db:"create_time"`
}
