<template>
  <div class="app-container">
<el-card class="box-card">

    <el-button type="primary" @click="handleAdd" icon="el-icon-plus">新建产品</el-button>
    <el-table
:header-cell-style="{background:'#eef1f6',color:'#606266'}"
:data="list" style="width: 100%;margin-top:30px;" @row-dblclick='handleEdit' border>
      <el-table-column
        label="序号"
        align="center"
        fixed
        width="50">
        <template slot-scope="scope">
          {{scope.$index+1}}
        </template>
      </el-table-column>
      <el-table-column align="center" label="产品id" width="220">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="产品名" width="220">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="产品介绍图" >
        <template slot-scope="scope">
          <img src="/image/5aa0b332af2af.png">
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" fixed="right">
        <template slot-scope="scope">
          <el-button type="primary" size="small" @click.stop="handleEdit(scope.row)" icon="el-icon-edit">编辑</el-button>
          <el-button type="danger" size="small" @click.stop="handleDelete(scope)" icon="el-icon-delete">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :visible.sync="dialogVisible" :title="dialogType==='edit'?'修改产品':'新建产品'">
      <el-form :model="form" label-width="120px" label-position="left">
        <el-form-item label="产品名">
          <el-input v-model="form.name" placeholder="产品名" />
        </el-form-item>
        <el-form-item label="产品介绍图">
          <Upload :limit="1" :data="form.introduce"></Upload>
        </el-form-item>
        <el-form-item label="点播系统">
          <Upload :limit="9" :data="form.system"></Upload>
        </el-form-item>
        <el-form-item label="特色功能">
          <Upload :limit="9" :data="form.item"></Upload>
        </el-form-item>
        <el-form-item label="硬件产品">
          <Upload :limit="9" :data="form.hardware"></Upload>
        </el-form-item>
        <el-form-item label="外设配套">
          <Upload :limit="9" :data="form.device"></Upload>
        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="dialogVisible=false" icon="el-icon-close">取消</el-button>
        <el-button type="primary" @click="submit" icon="el-icon-check">确定</el-button>
      </div>
    </el-dialog>

</el-card></div>
</template>

<script>
  import {getList,test} from "../../api/product"

  export default {
    data() {
      return {
        list: [],
        dialogVisible: false,
        dialogType: 'new',
        form:{
          name:'',
          introduce:[{ url: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',name:"111" }],
          system:[{ url: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',name:"111" }],
          item:[{ url: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',name:"111" }],
          hardware:[{ url: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',name:"111" }],
          device:[{ url: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',name:"111" }]
        },
        param:''
      }
    },
    created() {
      this.Search()
    },
    components:{
      Upload :()=>import("@/components/UploadImgs"),
    },
    methods: {
      submit(){
        this.param = new FormData()
        this.param.append("name",this.form.name)
        this.imgClassify(this.form.introduce,"introduce")
        this.imgClassify(this.form.introduce,"system")
        this.imgClassify(this.form.introduce,"item")
        this.imgClassify(this.form.introduce,"hardware")
        this.imgClassify(this.form.introduce,"device")

        test(this.param).then(res=>{

        }).catch(err=>{

        })
      },
      imgClassify(fileList,type){
        let prex = "_yinchuang_"
        let images = [...fileList] // 把数组存储为一个参数，便于后期操作
        // 2.2，遍历数组
        images.forEach((img, index) => {
          console.log(img,"img")
          this.param.append(`${type}${prex}${index}`, img) // 把单个图片重命名，存储起来（给后台）
        })
      },
      Search(){
        getList().then(res=>{
          this.list = res.data
        }).catch(err=>{
          this.$message({
            type: 'error',
            message: "网络异常"
          })
          console.log(err)
        })
      },
      // Reshape the routes structure so that it looks the same as the sidebar
      handleAdd() {
        this.dialogType = 'new'
        this.dialogVisible = true
      },
      handleEdit(row) {
        this.dialogType = 'edit'
        this.dialogVisible = true
      },
      handleDelete({ $index, row }) {
        this.$confirm('确定删除这个商品吗?', '警告', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
          .then(async() => {

          })
          .catch(err => { console.error(err) })
      },
      // reference: src/view/layout/components/Sidebar/SidebarItem.vue
    }
  }
</script>

<style lang="scss" scoped>
  .app-container {
    .roles-table {
      margin-top: 30px;
    }
    .permission-tree {
      margin-bottom: 30px;
    }
  }
</style>
