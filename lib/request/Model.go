package request

import model "ycoem/lib/models"

type Seo struct {
	Keywords    string `json:"keywords"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PhoneQQ struct {
	Phone string `json:"phone"`
	QQ    string `json:"qq"`
}

type Webinfo struct {
	Copyright  string `json:"copyright"`
	Icp        string `json:"icp"`
	RecordNo   string `json:"record_no"`
	RecordInfo string `json:"record_info"`
	WxImg      string `json:"wx_img"`
	AppImg     string `json:"app_img"`
}

type Contacts struct {
	DeptInfos string `json:"dept_infos"`
	Address   string `json:"address"`
}

type Logistics struct {
	Concat string `json:"concat"`
	Area   string `json:"area"`
}

type BranchOffice struct {
	Name          string `json:"name"`
	ServiceExtent string `json:"service_extent"`
	Area          string `json:"area"`
	Concat        string `json:"concat"`
}

type Contact struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Area    string `json:"area"`
	Manager string `json:"manager"`
	Name    string `json:"name"`
	Support string `json:"support"`
	Id      int    `json:"id"`
}

type Cases struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Title string `json:"title"`
	// Logo 图片
	Logo string `json:"logo"`
	// Desc 描述
	Desc string `json:"desc"`
	// Type 类型
	Type string `json:"type"`
	// Addr 位置
	Addr string `json:"addr"`
	Id   int    `json:"id"`
}

type Product struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name"  db:"name"`
	Img   string `json:"img" db:"img"`
}

type Service struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name"  db:"name"`
	Type  string `json:"type" db:"type"`
}

type Firm struct {
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Img     []string `json:"img"`
	WxLink  string   `json:"wx_link"`
	Content string   `json:"content"`
	Name    string   `json:"name"`
}

type ServiceAnswer struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Answer   string `json:"answer"`
	ParentID string `json:"parent_id"`
	Img      string `json:"img"`
}

type ServiceDown struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	ID          int    `json:"id"`
	DownloadURL string `json:"download_url"`
	Title       string `json:"title"`
	Version     string `json:"version"`
	ParentID    string `json:"parent_id"`
}

type ProductNode struct {
	ID        int      `json:"id" db:"id"`
	ProductID int      `json:"product_id" db:"product_id"`
	Imgs      []string `json:"imgs" db:"imgs"`
	Name      string   `json:"name" db:"name"`
}

type Solut struct {
	ID    int               `json:"id"`
	Page  int               `json:"page"`
	Limit int               `json:"limit"`
	Name  string            `json:"name"`
	Solut model.SolutModel  `json:"solut"`
	Fatd  []model.FatdModel `json:"fatd"`
}

type LunboTable struct {
	Img  string `json:"img"`
	Name string `json:"name"`
}
